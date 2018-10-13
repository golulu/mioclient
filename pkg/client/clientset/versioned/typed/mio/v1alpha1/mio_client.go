/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MioV1alpha1Interface interface {
	RESTClient() rest.Interface
	BuildsGetter
	BuildConfigsGetter
	PipelinesGetter
	PipelineConfigsGetter
	SourceConfigsGetter
}

// MioV1alpha1Client is used to interact with features provided by the mio.k8s.io group.
type MioV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MioV1alpha1Client) Builds(namespace string) BuildInterface {
	return newBuilds(c, namespace)
}

func (c *MioV1alpha1Client) BuildConfigs(namespace string) BuildConfigInterface {
	return newBuildConfigs(c, namespace)
}

func (c *MioV1alpha1Client) Pipelines(namespace string) PipelineInterface {
	return newPipelines(c, namespace)
}

func (c *MioV1alpha1Client) PipelineConfigs(namespace string) PipelineConfigInterface {
	return newPipelineConfigs(c, namespace)
}

func (c *MioV1alpha1Client) SourceConfigs(namespace string) SourceConfigInterface {
	return newSourceConfigs(c, namespace)
}

// NewForConfig creates a new MioV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MioV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MioV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MioV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MioV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MioV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MioV1alpha1Client {
	return &MioV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MioV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
