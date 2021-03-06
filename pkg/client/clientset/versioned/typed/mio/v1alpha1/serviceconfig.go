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
	scheme "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceConfigsGetter has a method to return a ServiceConfigInterface.
// A group's client should implement this interface.
type ServiceConfigsGetter interface {
	ServiceConfigs(namespace string) ServiceConfigInterface
}

// ServiceConfigInterface has methods to work with ServiceConfig resources.
type ServiceConfigInterface interface {
	Create(*v1alpha1.ServiceConfig) (*v1alpha1.ServiceConfig, error)
	Update(*v1alpha1.ServiceConfig) (*v1alpha1.ServiceConfig, error)
	UpdateStatus(*v1alpha1.ServiceConfig) (*v1alpha1.ServiceConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceConfig, err error)
	ServiceConfigExpansion
}

// serviceConfigs implements ServiceConfigInterface
type serviceConfigs struct {
	client rest.Interface
	ns     string
}

// newServiceConfigs returns a ServiceConfigs
func newServiceConfigs(c *MioV1alpha1Client, namespace string) *serviceConfigs {
	return &serviceConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceConfig, and returns the corresponding serviceConfig object, and an error if there is any.
func (c *serviceConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceConfig, err error) {
	result = &v1alpha1.ServiceConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceConfigs that match those selectors.
func (c *serviceConfigs) List(opts v1.ListOptions) (result *v1alpha1.ServiceConfigList, err error) {
	result = &v1alpha1.ServiceConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceConfigs.
func (c *serviceConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a serviceConfig and creates it.  Returns the server's representation of the serviceConfig, and an error, if there is any.
func (c *serviceConfigs) Create(serviceConfig *v1alpha1.ServiceConfig) (result *v1alpha1.ServiceConfig, err error) {
	result = &v1alpha1.ServiceConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceconfigs").
		Body(serviceConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceConfig and updates it. Returns the server's representation of the serviceConfig, and an error, if there is any.
func (c *serviceConfigs) Update(serviceConfig *v1alpha1.ServiceConfig) (result *v1alpha1.ServiceConfig, err error) {
	result = &v1alpha1.ServiceConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceconfigs").
		Name(serviceConfig.Name).
		Body(serviceConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *serviceConfigs) UpdateStatus(serviceConfig *v1alpha1.ServiceConfig) (result *v1alpha1.ServiceConfig, err error) {
	result = &v1alpha1.ServiceConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceconfigs").
		Name(serviceConfig.Name).
		SubResource("status").
		Body(serviceConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceConfig and deletes it. Returns an error if one occurs.
func (c *serviceConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceConfig.
func (c *serviceConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceConfig, err error) {
	result = &v1alpha1.ServiceConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
