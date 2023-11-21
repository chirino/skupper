package client

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/api/types"
)

func (cli *VanClient) SiteConfigRemove(ctx context.Context) error {
	err := cli.KubeClient.CoreV1().Secrets(cli.Namespace).Delete(ctx, types.SiteSecretName, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	return cli.KubeClient.CoreV1().ConfigMaps(cli.Namespace).Delete(ctx, types.SiteConfigMapName, metav1.DeleteOptions{})
}
