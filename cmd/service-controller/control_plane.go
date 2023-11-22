package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client"
	control_api "github.com/skupperproject/skupper/internal/control-api"
	"github.com/skupperproject/skupper/internal/crypto"
	"github.com/skupperproject/skupper/pkg/version"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"log"
	"os"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type ControlPlaneController struct {
	insecureSkipTlsVerify bool
	url                   string
	siteId                string
	vpcId                 string
	privateKey            crypto.Key
	publicKey             crypto.Key
	bearerToken           []byte

	ctx        context.Context
	cancel     context.CancelFunc
	shutdownWg sync.WaitGroup

	client              *control_api.APIClient
	sitesInformer       *control_api.Informer[control_api.ModelsSite]
	vanClient           *client.VanClient
	linkSecretsInformer cache.SharedIndexInformer
}

func NewControlPlaneController(spec types.ControlPlaneSpec, c *Controller) (*ControlPlaneController, error) {

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

	return &ControlPlaneController{
		insecureSkipTlsVerify: spec.InsecureSkipTlsVerify,
		url:                   spec.URL,
		siteId:                spec.SiteId,
		vpcId:                 spec.VpcId,
		privateKey:            privateKey,
		publicKey:             publicKey,
		bearerToken:           bearerToken,
		vanClient:             c.vanClient,
		linkSecretsInformer:   c.tokenHandler.informer,
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
	if c.insecureSkipTlsVerify { // #nosec G402
		options = append(options, control_api.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}

	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.client, err = control_api.NewAPIClientWithOptions(c.ctx, c.url, nil, options...)
	if err != nil {
		return fmt.Errorf("failed to connect to control plane: %w", err)
	}

	// event stream sharing occurs due to the informers sharing the context created in following line:
	informerCtx := c.client.VPCApi.WatchEvents(c.ctx, c.vpcId).PublicKey(string(c.publicKey)).NewSharedInformerContext()
	c.sitesInformer = c.client.VPCApi.ListSitesInVPC(informerCtx, c.vpcId).Informer()

	linkSecretsChanged := make(ChanEventHandler, 1)
	c.linkSecretsInformer.AddEventHandler(linkSecretsChanged)
	// TODO: is there is no RemoveEventHandler???

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
		update.SetHostname(hostname)
		updateOurSite = true
	}

	os := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	if err == nil && ourSite.GetOs() != os {
		update.SetOs(os)
		updateOurSite = true
	}

	// TODO: if ourSite.GetLinkSecret() is set, validate that it's still good..
	if ourSite.GetLinkSecret() == "" {
		s, _, err := c.vanClient.ConnectorTokenCreate(c.ctx, "skupper-control-plane-token", c.vanClient.Namespace)
		if err != nil {
			return fmt.Errorf("could not create connector token for the control plane: %w", err)
		}
		value, err := json.Marshal(s)
		if err != nil {
			return fmt.Errorf("could not json encode connector token for the control plane: %w", err)
		}
		update.SetLinkSecret(string(value))
		updateOurSite = true
	}

	if updateOurSite {
		log.Println("updating our site: ", update)
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

func retryWithExponentialBackOff(ctx context.Context, f func() error) error {
	ebo := backoff.NewExponentialBackOff()
	ebo.MaxInterval = 60 * time.Second
	ebo.MaxElapsedTime = 1 * time.Hour
	bo := backoff.WithContext(ebo, ctx)
	err := backoff.Retry(f, bo)
	return err
}
