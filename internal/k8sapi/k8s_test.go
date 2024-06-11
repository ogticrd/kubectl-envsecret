package k8sapi_test

import (
	"testing"

	"github.com/ogticrd/kubectl-envsecret/internal/k8sapi"
	"github.com/stretchr/testify/assert"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes/fake"
)

func TestK8sCreateSecret(t *testing.T) {
	fakeClient := fake.NewSimpleClientset()

	k := k8sapi.NewK8sClient(fakeClient, "test")

	t.Run("test CreateSecret returns expected results", func(t *testing.T) {
		err := k.CreateSecret("test", mockSecretData())
		assert.Nil(t, err)
	})
	t.Run("test CreateSecret fails with alreadyExists", func(t *testing.T) {
		err := k.CreateSecret("test", mockSecretData())
		assert.NotNil(t, err)
		assert.True(t, kerr.IsAlreadyExists(err))
	})
}

func mockSecretData() map[string]string {
	secret := make(map[string]string)
	secret["foo"] = `line1
  line2
  line3
  line4`
	secret["bar"] = "line"
	return secret
}
