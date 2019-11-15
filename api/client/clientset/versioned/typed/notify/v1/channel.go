/*
 * Copyright 2019 THL A29 Limited, a Tencent company.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/versioned/scheme"
	v1 "tkestack.io/tke/api/notify/v1"
)

// ChannelsGetter has a method to return a ChannelInterface.
// A group's client should implement this interface.
type ChannelsGetter interface {
	Channels() ChannelInterface
}

// ChannelInterface has methods to work with Channel resources.
type ChannelInterface interface {
	Create(*v1.Channel) (*v1.Channel, error)
	Update(*v1.Channel) (*v1.Channel, error)
	UpdateStatus(*v1.Channel) (*v1.Channel, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Channel, error)
	List(opts metav1.ListOptions) (*v1.ChannelList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Channel, err error)
	ChannelExpansion
}

// channels implements ChannelInterface
type channels struct {
	client rest.Interface
}

// newChannels returns a Channels
func newChannels(c *NotifyV1Client) *channels {
	return &channels{
		client: c.RESTClient(),
	}
}

// Get takes name of the channel, and returns the corresponding channel object, and an error if there is any.
func (c *channels) Get(name string, options metav1.GetOptions) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Get().
		Resource("channels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Channels that match those selectors.
func (c *channels) List(opts metav1.ListOptions) (result *v1.ChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ChannelList{}
	err = c.client.Get().
		Resource("channels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested channels.
func (c *channels) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("channels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a channel and creates it.  Returns the server's representation of the channel, and an error, if there is any.
func (c *channels) Create(channel *v1.Channel) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Post().
		Resource("channels").
		Body(channel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a channel and updates it. Returns the server's representation of the channel, and an error, if there is any.
func (c *channels) Update(channel *v1.Channel) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Put().
		Resource("channels").
		Name(channel.Name).
		Body(channel).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *channels) UpdateStatus(channel *v1.Channel) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Put().
		Resource("channels").
		Name(channel.Name).
		SubResource("status").
		Body(channel).
		Do().
		Into(result)
	return
}

// Delete takes name of the channel and deletes it. Returns an error if one occurs.
func (c *channels) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("channels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched channel.
func (c *channels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Patch(pt).
		Resource("channels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
