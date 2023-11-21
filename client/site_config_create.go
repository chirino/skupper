package client

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/pkg/site"
)

func (cli *VanClient) SiteConfigCreate(ctx context.Context, spec types.SiteConfigSpec) (*types.SiteConfig, error) {
	siteConfig, siteSecrets, err := site.WriteSiteConfig(spec, cli.Namespace)
	if err != nil {
		return nil, err
	}
	if spec.IsIngressRoute() && cli.RouteClient == nil {
		return nil, fmt.Errorf("OpenShift cluster not detected for --ingress type route")
	}

	actual, err := cli.KubeClient.CoreV1().ConfigMaps(cli.Namespace).Create(ctx, siteConfig, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	if actual.TypeMeta.Kind == "" || actual.TypeMeta.APIVersion == "" { // why??
		actual.TypeMeta = siteConfig.TypeMeta
	}

	if len(siteSecrets.Data) > 0 {
		siteSecrets, err = cli.KubeClient.CoreV1().Secrets(cli.Namespace).Create(ctx, siteSecrets, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		if siteSecrets.TypeMeta.Kind == "" || siteSecrets.TypeMeta.APIVersion == "" { // why??
			siteSecrets.TypeMeta = siteConfig.TypeMeta
		}
	}

	return cli.SiteConfigInspect(ctx, actual, siteSecrets)
}
