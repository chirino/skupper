package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/internal/control-api"
	"github.com/skupperproject/skupper/internal/crypto"
	"github.com/skupperproject/skupper/pkg/version"
	"net/http/httputil"
	"net/url"
	"runtime"
)

func JsonUnmarshal(from map[string]interface{}, to interface{}) error {
	b, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, to)
}

type PlatformSettings struct {
	Ingress string `json:"ingress,omitempty"`
}
type RegKeySettings struct {
	RouterMode string                      `json:"router_mode,omitempty"`
	labels     []string                    `json:"labels,omitempty"`
	SiteName   string                      `json:"site_name,omitempty"`
	Platforms  map[string]PlatformSettings `json:"platforms,omitempty"`
}

// registerWithControlPlane registers the site with the control pane, get initial site configuration, and
// configures the site with enough information so that he can contact the control plane to get updates
func registerWithControlPlane(initFlags *InitFlags, routerCreateOpts *types.SiteConfigSpec, platform types.Platform) error {

	controlPaneURL, err := url.Parse(initFlags.regKey)
	if err != nil {
		return fmt.Errorf("invalid --reg-key flag: %s", err)
	}
	if controlPaneURL.Scheme != "https" {
		return fmt.Errorf("invalid --service-url flag: 'https://' URL scheme is required")
	}
	controlPaneURL.Host = "api." + controlPaneURL.Host
	controlPaneURL.Path = ""

	token := controlPaneURL.Fragment
	controlPaneURL.Fragment = ""

	options := []control_api.Option{
		control_api.WithUserAgent(fmt.Sprintf("skupper/%s (%s; %s)", version.Version, runtime.GOOS, runtime.GOARCH)),
		control_api.WithBearerToken(token), // the token is used as a Bearer Token in the initial registration API interactions.
	}
	if initFlags.controlPlaneAddress != "" {
		options = append(options, control_api.WithHostAddress(initFlags.controlPlaneAddress))
	}
	if initFlags.insecureSkipTlsVerify { // #nosec G402
		options = append(options, control_api.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}

	ctx := context.Background()
	client, err := control_api.NewAPIClientWithOptions(ctx, controlPaneURL.String(), nil, options...)

	regKeyModel, httpResp, err := client.RegKeyApi.GetRegKey(ctx, "me").Execute()
	if err != nil {
		x, _ := httputil.DumpResponse(httpResp, true)
		fmt.Println(string(x))
		return fmt.Errorf("could not fetch reg key settings: %w", err)
	}

	settings := RegKeySettings{}
	err = JsonUnmarshal(regKeyModel.Settings, &settings)
	if err != nil {
		return fmt.Errorf("could not parse reg key settings: %w", err)
	}

	priv, err := crypto.NewKey()
	if err != nil {
		return fmt.Errorf("could create key pair: %w", err)
	}
	pub, err := priv.PublicKey()
	if err != nil {
		return fmt.Errorf("could create key pair: %w", err)
	}

	site, _, err := client.SitesApi.CreateSite(ctx).Site(control_api.ModelsAddSite{
		Platform:         control_api.PtrString(string(platform)),
		Name:             control_api.PtrString(routerCreateOpts.SkupperName),
		PublicKey:        control_api.PtrString(pub.String()),
		ServiceNetworkId: regKeyModel.ServiceNetworkId,
	}).Execute()
	if err != nil {
		return fmt.Errorf("could register site with the control plan: %w", err)
	}

	// Configure init settings from the RegKey settings
	if settings.RouterMode != "" {
		initFlags.routerMode = settings.RouterMode
	}
	if settings.labels != nil {
		initFlags.labels = settings.labels
	}
	if settings.SiteName != "" {
		routerCreateOpts.SkupperName = settings.SiteName
	}
	if settings.SiteName != "" {
		routerCreateOpts.Ingress = settings.SiteName
	}
	if p, found := settings.Platforms[string(platform)]; found {
		if p.Ingress != "" {
			routerCreateOpts.Ingress = p.Ingress
		}
	}

	// Configure the router know it should connect to the control plane
	routerCreateOpts.ControlPlane.Enabled = true
	routerCreateOpts.ControlPlane.InsecureSkipTlsVerify = initFlags.insecureSkipTlsVerify
	routerCreateOpts.ControlPlane.URL = controlPaneURL.String()
	routerCreateOpts.ControlPlane.Address = initFlags.controlPlaneAddress
	routerCreateOpts.ControlPlane.SiteId = site.GetId()
	routerCreateOpts.ControlPlane.ServiceNetworkId = site.GetServiceNetworkId()

	// Now that the site is registered, subsequent API calls must use the per site bearer token
	// which can only be obtained by decrypting it with this private which will only be known by the site.
	routerCreateOpts.ControlPlane.PrivateKey = priv.String()
	routerCreateOpts.ControlPlane.BearerTokenEncrypted = site.GetBearerToken()

	return nil
}
