package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client"
	control_api "github.com/skupperproject/skupper/internal/control-api"
	"github.com/skupperproject/skupper/internal/crypto"
	"github.com/skupperproject/skupper/pkg/version"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1informer "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
	"log"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

type ControlPlaneController struct {
	insecureSkipTlsVerify bool
	url                   *url.URL
	address               string
	siteId                string
	serviceNetworkId      string
	privateKey            crypto.Key
	publicKey             crypto.Key
	bearerToken           []byte

	ctx        context.Context
	cancel     context.CancelFunc
	shutdownWg sync.WaitGroup

	client                    *control_api.APIClient
	sitesInformer             *control_api.Informer[control_api.ModelsSite]
	vanClient                 *client.VanClient
	linkSecretsInformer       cache.SharedIndexInformer
	siteServerSecretsInformer cache.SharedIndexInformer
	cli                       *client.VanClient
}

func NewControlPlaneController(cli *client.VanClient, spec types.ControlPlaneSpec, c *Controller) (*ControlPlaneController, error) {

	privateKey, err := crypto.ParseKey(spec.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid control plane private key: %w", err)
	}

	sealed, err := crypto.ParseSealed(spec.BearerTokenEncrypted)
	if err != nil {
		return nil, fmt.Errorf("control plane bearer token invalid: %w", err)
	}

	bearerToken, err := sealed.Open(privateKey[:])
	if err != nil {
		return nil, fmt.Errorf("control plane bearer token decryption failed: %w", err)
	}

	publicKey, err := privateKey.PublicKey()
	if err != nil {
		return nil, fmt.Errorf("invalid control plane public key: %w", err)
	}

	siteServerSecretsInformer := corev1informer.NewFilteredSecretInformer(
		cli.KubeClient,
		cli.Namespace,
		time.Second*30,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		func(options *metav1.ListOptions) {
			options.FieldSelector = "metadata.name=skupper-site-server"
			options.LabelSelector = "!" + types.SiteControllerIgnore
		})

	u, err := url.Parse(spec.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid control plane url: %w", err)
	}

	return &ControlPlaneController{
		cli:                       cli,
		insecureSkipTlsVerify:     spec.InsecureSkipTlsVerify,
		url:                       u,
		address:                   spec.Address,
		siteId:                    spec.SiteId,
		serviceNetworkId:          spec.ServiceNetworkId,
		privateKey:                privateKey,
		publicKey:                 publicKey,
		bearerToken:               bearerToken,
		vanClient:                 c.vanClient,
		linkSecretsInformer:       c.tokenHandler.informer,
		siteServerSecretsInformer: siteServerSecretsInformer,
	}, nil
}

type ChanEventHandler chan struct{}

func (c ChanEventHandler) OnAdd(obj interface{}) {
	c <- struct{}{}
}

func (c ChanEventHandler) OnUpdate(oldObj, newObj interface{}) {
	c <- struct{}{}
}

func (c ChanEventHandler) OnDelete(obj interface{}) {
	c <- struct{}{}
}

func (c *ControlPlaneController) start(stopCh <-chan struct{}) error {
	log.Println("starting control plane controller")
	var err error

	options := []control_api.Option{
		control_api.WithUserAgent(fmt.Sprintf("skupper-service-controller/%s (%s; %s)", version.Version, runtime.GOOS, runtime.GOARCH)),
		control_api.WithBearerToken(string(c.bearerToken)),
	}
	if c.address != "" {
		options = append(options, control_api.WithHostAddress(c.address))
	}
	if c.insecureSkipTlsVerify { // #nosec G402
		options = append(options, control_api.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}

	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.client, err = control_api.NewAPIClientWithOptions(c.ctx, c.url.String(), nil, options...)
	if err != nil {
		return fmt.Errorf("failed to connect to control plane: %w", err)
	}

	// event stream sharing occurs due to the informers sharing the context created in following line:
	informerCtx := c.client.ServiceNetworkApi.WatchEventsInServiceNetwork(c.ctx, c.serviceNetworkId).NewSharedInformerContext()
	c.sitesInformer = c.client.ServiceNetworkApi.ListSitesInServiceNetwork(informerCtx, c.serviceNetworkId).Informer()

	linkSecretsChanged := make(ChanEventHandler, 1)
	c.linkSecretsInformer.AddEventHandler(linkSecretsChanged)

	siteServerSecretsChanged := make(ChanEventHandler, 1)
	c.siteServerSecretsInformer.AddEventHandler(siteServerSecretsChanged)
	go c.siteServerSecretsInformer.Run(c.ctx.Done())

	c.shutdownWg.Add(1)
	go func() {
		defer c.shutdownWg.Done()

		log.Println("running control plane controller event loop")
		err := c.reconcileControlPlaneSites()
		if err != nil {
			log.Println("failed to reconcile sites:", err)
		}

		for {
			select {
			case <-stopCh:
			case <-c.ctx.Done():
				return // occurs when controller is stopped.

			case <-linkSecretsChanged:
				log.Printf("reconciling due to: link secrets change\n")
				err := c.reconcileControlPlaneSites()
				if err != nil {
					log.Println("failed to reconcile sites:", err)
				}

			case <-siteServerSecretsChanged:
				log.Printf("reconciling due to: site secrets change\n")
				err := c.reconcileControlPlaneSites()
				if err != nil {
					log.Println("failed to reconcile sites:", err)
				}

			case <-c.sitesInformer.Changed():
				log.Printf("reconciling due to: control plance sites change\n")
				err := c.reconcileControlPlaneSites()
				if err != nil {
					log.Println("failed to reconcile sites:", err)
				}

			}
		}
	}()
	return nil
}
func (c *ControlPlaneController) stop() {
	fmt.Println("stopping control plane controller")
	c.cancel()
	c.shutdownWg.Wait()
}

func (c *ControlPlaneController) reconcileControlPlaneSites() error {
	var sites map[string]control_api.ModelsSite = nil

	// to avoid spinning when the control plane is not available, we retry with exponential backoff.
	err := retryWithExponentialBackOff(c.ctx, func() error {
		var err error
		sites, err = control_api.Simplify(c.sitesInformer.Execute())
		if err != nil {
			log.Println("control plane request error:", err)
		}
		return err
	})
	if err != nil {
		return err
	}

	ourSite, found := sites[c.siteId]
	if !found {
		return fmt.Errorf("our site not found in the network defined by the control plane")
	}

	updateOurSite := false
	update := control_api.ModelsUpdateSite{}

	hostname, err := os.Hostname()
	if err == nil && ourSite.GetHostname() != hostname {
		log.Println("updating our site hostname: ", hostname)
		update.SetHostname(hostname)
		updateOurSite = true
	}

	os := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	if err == nil && ourSite.GetOs() != os {
		log.Println("updating our site os: ", os)
		update.SetOs(os)
		updateOurSite = true
	}

	log.Println("getting skupper-site-server secret")
	skupperSiteServerSecret, err := c.getSkupperSiteServerSecret()
	if err != nil {
		return err
	}

	log.Println("parsing the skupper-site-server secrets")
	cert, err := ParseCertificate(skupperSiteServerSecret.Data["tls.crt"])
	if err != nil {
		return fmt.Errorf("invalid skupper-site-server secret: %w", err)
	}
	key, err := ParseCertificateKey(skupperSiteServerSecret.Data["tls.key"])
	if err != nil {
		return fmt.Errorf("invalid skupper-site-server secret: %w", err)
	}

	controlPlaneSigned := strings.HasPrefix(cert.Issuer.CommonName, c.url.String())

	err = VerifyCertificate(skupperSiteServerSecret.Data["tls.crt"], skupperSiteServerSecret.Data["ca.crt"])
	if err != nil {
		log.Println("control plane certificates did not pass validation: ", err)
		controlPlaneSigned = false // lets get them signed...
	}

	if !controlPlaneSigned {

		log.Println("the skupper-site-server secrets are not signed by the control plane")
		// get our cert signed by the control plane...

		template := x509.CertificateRequest{
			Version:            cert.Version,
			Signature:          cert.Signature,
			SignatureAlgorithm: cert.SignatureAlgorithm,
			PublicKeyAlgorithm: cert.PublicKeyAlgorithm,
			PublicKey:          cert.PublicKey,
			Subject:            cert.Subject,
			Extensions:         cert.Extensions,
			ExtraExtensions:    cert.ExtraExtensions,
			DNSNames:           cert.DNSNames,
			EmailAddresses:     cert.EmailAddresses,
			IPAddresses:        cert.IPAddresses,
			URIs:               cert.URIs,
		}
		csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, key)

		response, err := control_api.Simplify(c.client.CAApi.
			SignCSR(c.ctx).
			CertificateSigningRequest(control_api.ModelsCertificateSigningRequest{
				Request: control_api.PtrString(string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes}))),
				Usages: []control_api.ModelsKeyUsage{
					control_api.UsageSigning,
					control_api.UsageKeyEncipherment,
					control_api.UsageServerAuth,
					control_api.UsageClientAuth,
				},
			}).
			Execute(),
		)
		if err != nil {
			return err
		}

		log.Println("updating skupper-site-server secrets with the versions signed by the control plane")

		cert, err = ParseCertificate([]byte(response.GetCertificate()))
		if err != nil {
			return fmt.Errorf("control plane returned an invalid certificate: %w", err)
		}

		_, err = ParseCertificate([]byte(response.GetCa()))
		if err != nil {
			return fmt.Errorf("control plane returned an invalid ca: %w", err)
		}

		// replace the cert/ca with the signed cert/ca from the control plane
		if skupperSiteServerSecret.Data == nil {
			skupperSiteServerSecret.Data = map[string][]byte{}
		}
		skupperSiteServerSecret.Data["tls.crt"] = []byte(response.GetCertificate())
		//skupperSiteServerSecret.Data["ca.crt"] = []byte(response.GetCa())
		skupperSiteServerSecret.Data["ca.crt"] = append(skupperSiteServerSecret.Data["ca.crt"], []byte(response.GetCa())...)

		err = VerifyCertificate(skupperSiteServerSecret.Data["tls.crt"], skupperSiteServerSecret.Data["ca.crt"])
		if err != nil {
			return fmt.Errorf("control plane certificates did not pass validation: %w", err)
		}

		_, err = c.vanClient.KubeClient.CoreV1().Secrets(skupperSiteServerSecret.Namespace).Update(c.ctx, skupperSiteServerSecret, metav1.UpdateOptions{})
		if err != nil {
			return err
		}

		// restart the router to pick up the new certs
		c.cli.RestartRouter(c.cli.Namespace)
	}

	// consider using c.vanClient.annotateConnectorToken() instead of the following code...
	token, _, err := c.vanClient.ConnectorTokenCreate(c.ctx, "skupper-control-plane-token", c.vanClient.Namespace)
	if err != nil {
		return fmt.Errorf("could not create connector token for the control plane: %w", err)
	}
	token.Data = nil // we don't need the data, just the metadata (contains the connection address)
	linkSecretBytes, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return fmt.Errorf("could not json encode connector token for the control plane: %w", err)
	}
	linkSecret := string(linkSecretBytes)

	if ourSite.GetLinkSecret() != linkSecret {
		log.Println("updating link secret: ", linkSecret)
		update.SetLinkSecret(linkSecret)
		updateOurSite = true
	}

	if updateOurSite {
		_, err = control_api.Simplify(c.client.SitesApi.
			UpdateSite(c.ctx, c.siteId).
			Update(update).
			Execute(),
		)
		if err != nil {
			return err
		}
	}

	secretsApi := c.vanClient.KubeClient.CoreV1().Secrets(c.vanClient.Namespace)

	// Let's get all the secrets that we have created for our peers...
	secretsByName := map[string]corev1.Secret{}
	for _, o := range c.linkSecretsInformer.GetStore().List() {
		s := o.(*corev1.Secret)
		// we are only interested in the link secrets that we created
		if s.Labels[types.InternalQualifier+"/source"] == "control-plane" {
			secretsByName[s.Name] = *s
		}
	}

	for _, site := range sites {
		// skip our site since we don't want a link to ourselves
		if site.GetId() == c.siteId {
			continue
		}
		secretTxt := site.GetLinkSecret()
		if secretTxt == "" {
			log.Printf("peer site has no link secret: %s\n", site.GetId())
		} else {
			secret := corev1.Secret{}
			err := json.Unmarshal([]byte(secretTxt), &secret)
			if err != nil {
				log.Printf("failed to json unmarshal link secret for site %s: %s, data: %s\n", site.GetId(), err, secretTxt)
				continue
			}

			secret.Namespace = c.vanClient.Namespace
			secret.Name = fmt.Sprintf("cp-%s", site.GetId())
			if secret.Labels == nil {
				secret.Labels = map[string]string{}
			}
			secret.Labels[types.InternalQualifier+"/source"] = "control-plane"

			// Connect to the peer using this site's identity
			secret.Data = map[string][]byte{}
			for key, value := range skupperSiteServerSecret.Data {
				secret.Data[key] = value
			}

			// TODO: we could add more attributes to the site record to determine if we should
			// create a link to the peer site.  For example there could be a setting that says our site should only
			// connect to other sites that are relay sites (we could also filter peer sites server side to aid with
			// scale out performance).
			// For now, just link to all peers.

			if existing, found := secretsByName[secret.Name]; found {
				// TODO: are there any other fields we should be checking?
				if !reflect.DeepEqual(existing.Data, secret.Data) {
					existing.Data = secret.Data
					_, err = secretsApi.Update(c.ctx, &existing, metav1.UpdateOptions{})
					if err != nil {
						log.Printf("failed to update link secret for site %s: %s\n", site.GetId(), err)
						continue
					}
					log.Printf("updated link secret for site %s\n", site.GetId())
				} else {
					log.Printf("link secret is up to date for site %s\n", site.GetId())
				}
				delete(secretsByName, secret.Name) // remove from map so we know we've processed it
			} else {
				_, err = secretsApi.Create(c.ctx, &secret, metav1.CreateOptions{})
				if err != nil {
					log.Printf("failed to create link secret for site %s: %s\n", site.GetId(), err)
					continue
				}
				log.Printf("created link secret for site %s\n", site.GetId())
			}
		}
	}

	// delete any secrets that we created that are no longer needed
	for _, secret := range secretsByName {
		err := secretsApi.Delete(c.ctx, secret.Name, metav1.DeleteOptions{})
		if err != nil {
			log.Printf("failed to delete link secret %s: %s\n", secret.Name, err)
		}
		log.Printf("deleted link secret for site %s\n", secret.Name)
	}

	return nil
}

func (c *ControlPlaneController) getSkupperSiteServerSecret() (*corev1.Secret, error) {
	list := c.siteServerSecretsInformer.GetStore().List()
	if len(list) == 0 {
		return nil, fmt.Errorf("secret skupper-site-server not found")
	}
	return list[0].(*corev1.Secret), nil
}

func retryWithExponentialBackOff(ctx context.Context, f func() error) error {
	ebo := backoff.NewExponentialBackOff()
	ebo.MaxInterval = 60 * time.Second
	ebo.MaxElapsedTime = 1 * time.Hour
	bo := backoff.WithContext(ebo, ctx)
	err := backoff.Retry(f, bo)
	return err
}

func VerifyCertificate(certPEMBlock []byte, caCertPEMBlock []byte) error {

	cert, err := ParseCertificate(certPEMBlock)
	if err != nil {
		return err
	}

	// Create a pool of trusted CA certificates
	roots, err := NewCertPool(caCertPEMBlock)
	if err != nil {
		return fmt.Errorf("error parsing ca certificates: %w", err)
	}

	// Verify that the certificate was signed by the CA
	opts := x509.VerifyOptions{
		Roots:     roots,
		KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}

	if _, err := cert.Verify(opts); err != nil {
		return err
	}
	return nil
}

func ParseCertificate(certPEMBlock []byte) (*x509.Certificate, error) {
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil || certDERBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("cert pem block type is not CERTIFICATE")
	}
	cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the certificate: %w", err)
	}
	return cert, nil
}
func ParseCertificateKey(keyPEMBlock []byte) (any, error) {
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil || !strings.HasSuffix(keyDERBlock.Type, "PRIVATE KEY") {
		return nil, fmt.Errorf("key pem block type is not PRIVATE KEY: %s", keyDERBlock.Type)
	}
	key, err := x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the PKCS1 key: %w", err)
	}

	return key, nil
}

func NewCertPool(pemBytes []byte) (*x509.CertPool, error) {
	roots := x509.NewCertPool()
	for {
		var block *pem.Block
		block, pemBytes = pem.Decode(pemBytes)
		if block == nil {
			break
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
		roots.AddCert(cert)
	}
	return roots, nil
}
