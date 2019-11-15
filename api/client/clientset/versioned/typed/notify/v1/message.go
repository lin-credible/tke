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

// MessagesGetter has a method to return a MessageInterface.
// A group's client should implement this interface.
type MessagesGetter interface {
	Messages() MessageInterface
}

// MessageInterface has methods to work with Message resources.
type MessageInterface interface {
	Create(*v1.Message) (*v1.Message, error)
	Update(*v1.Message) (*v1.Message, error)
	UpdateStatus(*v1.Message) (*v1.Message, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Message, error)
	List(opts metav1.ListOptions) (*v1.MessageList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Message, err error)
	MessageExpansion
}

// messages implements MessageInterface
type messages struct {
	client rest.Interface
}

// newMessages returns a Messages
func newMessages(c *NotifyV1Client) *messages {
	return &messages{
		client: c.RESTClient(),
	}
}

// Get takes name of the message, and returns the corresponding message object, and an error if there is any.
func (c *messages) Get(name string, options metav1.GetOptions) (result *v1.Message, err error) {
	result = &v1.Message{}
	err = c.client.Get().
		Resource("messages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Messages that match those selectors.
func (c *messages) List(opts metav1.ListOptions) (result *v1.MessageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MessageList{}
	err = c.client.Get().
		Resource("messages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messages.
func (c *messages) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("messages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a message and creates it.  Returns the server's representation of the message, and an error, if there is any.
func (c *messages) Create(message *v1.Message) (result *v1.Message, err error) {
	result = &v1.Message{}
	err = c.client.Post().
		Resource("messages").
		Body(message).
		Do().
		Into(result)
	return
}

// Update takes the representation of a message and updates it. Returns the server's representation of the message, and an error, if there is any.
func (c *messages) Update(message *v1.Message) (result *v1.Message, err error) {
	result = &v1.Message{}
	err = c.client.Put().
		Resource("messages").
		Name(message.Name).
		Body(message).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *messages) UpdateStatus(message *v1.Message) (result *v1.Message, err error) {
	result = &v1.Message{}
	err = c.client.Put().
		Resource("messages").
		Name(message.Name).
		SubResource("status").
		Body(message).
		Do().
		Into(result)
	return
}

// Delete takes name of the message and deletes it. Returns an error if one occurs.
func (c *messages) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("messages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched message.
func (c *messages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Message, err error) {
	result = &v1.Message{}
	err = c.client.Patch(pt).
		Resource("messages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
