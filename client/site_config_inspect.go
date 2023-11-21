package client

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/pkg/site"
)

func (cli *VanClient) SiteConfigInspect(ctx context.Context, input *corev1.ConfigMap, secretsInput *corev1.Secret) (*types.SiteConfig, error) {
	var namespace string
	if input == nil {
		namespace = cli.Namespace
	} else {
		namespace = input.ObjectMeta.Namespace
	}
	return cli.SiteConfigInspectInNamespace(ctx, input, secretsInput, namespace)
}

func (cli *VanClient) SiteConfigInspectInNamespace(ctx context.Context, input *corev1.ConfigMap, secretsInput *corev1.Secret, namespace string) (*types.SiteConfig, error) {
	var siteConfig *corev1.ConfigMap
	if input == nil {
		cm, err := cli.KubeClient.CoreV1().ConfigMaps(namespace).Get(ctx, types.SiteConfigMapName, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		siteConfig = cm
	} else {
		siteConfig = input
	}

	var siteSecrets *corev1.Secret = nil
	if secretsInput == nil {
		value, err := cli.KubeClient.CoreV1().Secrets(namespace).Get(ctx, types.SiteSecretName, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			value = &corev1.Secret{
				Data: map[string][]byte{},
			}
		} else if err != nil {
			return nil, err
		}
		siteSecrets = value
	} else {
		siteSecrets = secretsInput
	}
	return site.ReadSiteConfig(siteConfig, siteSecrets, namespace, cli.GetIngressDefault())
}
