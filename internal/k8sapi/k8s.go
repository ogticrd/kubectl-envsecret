// Package k8sapi provides functions and types for interacting with the Kubernetes API.
//
// This package includes utilities for creating Kubernetes clients and managing
// secrets within a specified namespace.
package k8sapi

import (
	"context"
	"fmt"

	"github.com/ogticrd/kubectl-envsecret/internal/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// K8sClient encapsulates a Kubernetes client and the namespace it operates within.
type K8sClient struct {
	client    kubernetes.Interface // Kubernetes client interface.
	namespace string               // Namespace for the Kubernetes operations.
}

// K8sConfig holds the configuration needed to create a Kubernetes client.
type K8sConfig struct {
	rest      *rest.Config // Kubernetes REST configuration.
	namespace string       // Namespace for the Kubernetes operations.
}

// NewK8sClient creates a new K8sClient instance.
//
// Parameters:
// - client: Kubernetes client interface.
// - namespace: Namespace for the Kubernetes operations.
//
// Returns:
// - A new K8sClient instance.
//
// Example usage:
// client := kubernetes.NewForConfig(config)
// k8sClient := NewK8sClient(client, "default")
func NewK8sClient(client kubernetes.Interface, namespace string) *K8sClient {
	return &K8sClient{
		client:    client,
		namespace: namespace,
	}
}

// NewK8sConfig creates a new K8sConfig instance.
//
// Parameters:
// - rest: Kubernetes REST configuration.
// - namespace: Namespace for the Kubernetes operations.
//
// Returns:
// - A new K8sConfig instance.
//
// Example usage:
// config := &rest.Config{Host: "https://my-k8s-cluster"}
// k8sConfig := NewK8sConfig(config, "default")
func NewK8sConfig(rest *rest.Config, namespace string) *K8sConfig {
	return &K8sConfig{
		rest:      rest,
		namespace: namespace,
	}
}

// NewK8sClientFromConfig creates a new K8sClient instance from the provided K8sConfig.
//
// Parameters:
// - config: K8sConfig instance containing the REST configuration and namespace.
//
// Returns:
// - A new K8sClient instance or an error if the client creation fails.
//
// Example usage:
// k8sConfig := NewK8sConfig(config, "default")
// k8sClient, err := NewK8sClientFromConfig(k8sConfig)
func NewK8sClientFromConfig(config *K8sConfig) (*K8sClient, error) {
	clientset, err := kubernetes.NewForConfig(config.rest)
	if err != nil {
		return nil, err
	}
	return NewK8sClient(clientset, config.namespace), nil
}

// CreateSecret creates a new Kubernetes secret with the provided name and data.
//
// Parameters:
// - secretName: Name of the Kubernetes secret.
// - secrets: Map containing the secret data as key-value pairs.
//
// Returns:
// - An error if the secret creation fails.
//
// Example usage:
// secrets := map[string]string{"username": "admin", "password": "secret"}
// err := k8sClient.CreateSecret("my-secret", secrets)
func (c *K8sClient) CreateSecret(secretName string, secrets map[string]string) error {
	if len(secrets) == 0 {
		return fmt.Errorf("no secrets provided")
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

	fmt.Printf("\nSecret %s created successfully.", secretName)
	return nil
}
