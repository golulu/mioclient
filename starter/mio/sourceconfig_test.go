package mio

import (
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestSourceConfigCurd(t *testing.T) {
	name := "test"
	name1 := "test1"
	namespace := "demo-dev"
	clientSet := fake.NewSimpleClientset().MioV1alpha1()
	config := newSourceConfig(clientSet)
	config1 := &v1alpha1.SourceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace:namespace,
		},
	}
	result, err := config.Create(config1)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	config2 := &v1alpha1.SourceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: name1,
			Namespace:namespace,
		},
	}
	result, err = config.Create(config2)
	assert.Equal(t, nil, err)
	assert.Equal(t, name1, result.Name)

	list, err := config.List(namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(list.Items))

	result, err = config.Get(name, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	con := &v1alpha1.SourceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	result, err = config.Update(name, namespace, con)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	err = config.Delete(name, namespace)
	assert.Equal(t, nil, err)

}