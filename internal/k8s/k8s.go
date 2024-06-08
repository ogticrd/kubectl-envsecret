package k8s

import (
	"context"
	"fmt"

	"github.com/ogticrd/kubectl-envsecret/internal/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sClient struct {
	client    kubernetes.Interface
	namespace string
}

type K8sConfig struct {
	rest      *rest.Config
	namespace string
}

func NewK8sClient(client kubernetes.Interface, namespace string) *K8sClient {
	return &K8sClient{
		client:    client,
		namespace: namespace,
	}
}

func NewK8sConfig(rest *rest.Config, namespace string) *K8sConfig {
	return &K8sConfig{
		rest:      rest,
		namespace: namespace,
	}
}

func NewK8sClientFromConfig(config *K8sConfig) (*K8sClient, error) {
	clientset, err := kubernetes.NewForConfig(config.rest)
	if err != nil {
		return nil, err
	}
	return NewK8sClient(clientset, config.namespace), nil
}

func (c *K8sClient) CreateSecret(secretName string, secrets map[string]string) error {
	if len(secrets) == 0 {
		return fmt.Errorf(fmt.Sprintf("No secrets provided."))
	}

	const secretType v1.SecretType = "Opaque"

	_, err := c.client.CoreV1().Secrets(c.namespace).Create(
		context.TODO(),
		&v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name: secretName,
			},
			Type: secretType,
			Data: utils.MapStringToBytes(secrets),
		},
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	fmt.Printf("\nSecret %s created successfuly.", secretName)
	return nil
}
