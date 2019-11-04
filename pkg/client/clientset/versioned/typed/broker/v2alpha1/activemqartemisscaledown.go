/*
Copyright The Kubernetes Authors.

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

package v2alpha1

import (
	v2alpha1 "github.com/rh-messaging/activemq-artemis-operator/pkg/apis/broker/v2alpha1"
	scheme "github.com/rh-messaging/activemq-artemis-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ActiveMQArtemisScaledownsGetter has a method to return a ActiveMQArtemisScaledownInterface.
// A group's client should implement this interface.
type ActiveMQArtemisScaledownsGetter interface {
	ActiveMQArtemisScaledowns(namespace string) ActiveMQArtemisScaledownInterface
}

// ActiveMQArtemisScaledownInterface has methods to work with ActiveMQArtemisScaledown resources.
type ActiveMQArtemisScaledownInterface interface {
	Create(*v2alpha1.ActiveMQArtemisScaledown) (*v2alpha1.ActiveMQArtemisScaledown, error)
	Update(*v2alpha1.ActiveMQArtemisScaledown) (*v2alpha1.ActiveMQArtemisScaledown, error)
	UpdateStatus(*v2alpha1.ActiveMQArtemisScaledown) (*v2alpha1.ActiveMQArtemisScaledown, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2alpha1.ActiveMQArtemisScaledown, error)
	List(opts v1.ListOptions) (*v2alpha1.ActiveMQArtemisScaledownList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2alpha1.ActiveMQArtemisScaledown, err error)
	ActiveMQArtemisScaledownExpansion
}

// activeMQArtemisScaledowns implements ActiveMQArtemisScaledownInterface
type activeMQArtemisScaledowns struct {
	client rest.Interface
	ns     string
}

// newActiveMQArtemisScaledowns returns a ActiveMQArtemisScaledowns
func newActiveMQArtemisScaledowns(c *BrokerV2alpha1Client, namespace string) *activeMQArtemisScaledowns {
	return &activeMQArtemisScaledowns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the activeMQArtemisScaledown, and returns the corresponding activeMQArtemisScaledown object, and an error if there is any.
func (c *activeMQArtemisScaledowns) Get(name string, options v1.GetOptions) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	result = &v2alpha1.ActiveMQArtemisScaledown{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ActiveMQArtemisScaledowns that match those selectors.
func (c *activeMQArtemisScaledowns) List(opts v1.ListOptions) (result *v2alpha1.ActiveMQArtemisScaledownList, err error) {
	result = &v2alpha1.ActiveMQArtemisScaledownList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested activeMQArtemisScaledowns.
func (c *activeMQArtemisScaledowns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a activeMQArtemisScaledown and creates it.  Returns the server's representation of the activeMQArtemisScaledown, and an error, if there is any.
func (c *activeMQArtemisScaledowns) Create(activeMQArtemisScaledown *v2alpha1.ActiveMQArtemisScaledown) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	result = &v2alpha1.ActiveMQArtemisScaledown{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		Body(activeMQArtemisScaledown).
		Do().
		Into(result)
	return
}

// Update takes the representation of a activeMQArtemisScaledown and updates it. Returns the server's representation of the activeMQArtemisScaledown, and an error, if there is any.
func (c *activeMQArtemisScaledowns) Update(activeMQArtemisScaledown *v2alpha1.ActiveMQArtemisScaledown) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	result = &v2alpha1.ActiveMQArtemisScaledown{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		Name(activeMQArtemisScaledown.Name).
		Body(activeMQArtemisScaledown).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *activeMQArtemisScaledowns) UpdateStatus(activeMQArtemisScaledown *v2alpha1.ActiveMQArtemisScaledown) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	result = &v2alpha1.ActiveMQArtemisScaledown{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		Name(activeMQArtemisScaledown.Name).
		SubResource("status").
		Body(activeMQArtemisScaledown).
		Do().
		Into(result)
	return
}

// Delete takes name of the activeMQArtemisScaledown and deletes it. Returns an error if one occurs.
func (c *activeMQArtemisScaledowns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *activeMQArtemisScaledowns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched activeMQArtemisScaledown.
func (c *activeMQArtemisScaledowns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	result = &v2alpha1.ActiveMQArtemisScaledown{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("activemqartemisscaledowns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}