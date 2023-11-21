package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/skupperproject/skupper/api/types"
	control_api "github.com/skupperproject/skupper/internal/control-api"
	"github.com/skupperproject/skupper/internal/crypto"
	"github.com/skupperproject/skupper/pkg/version"
	"log"
	"os"
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

	client        *control_api.APIClient
	sitesInformer *control_api.Informer[control_api.ModelsSite]
}

func NewControlPlaneController(spec types.ControlPlaneSpec) (*ControlPlaneController, error) {

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
	}, nil
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
			case <-c.sitesInformer.Changed():
				log.Printf("reconciling sites\n")
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
		return fmt.Errorf("our site not found in sites")
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

	peerLinkTokens := make(map[string]string)
	for _, site := range sites {
		if site.GetId() == c.siteId {
			continue
		}
		if site.HasLinkSecret() {
			peerLinkTokens[site.GetId()] = site.GetLinkSecret()
		}
	}

	// todo: reconcile link tokens with kube..
	log.Printf("peers: %v\n", peerLinkTokens)
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
