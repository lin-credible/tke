/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	auth "tkestack.io/tke/api/auth"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
)

// LocalIdentitiesGetter has a method to return a LocalIdentityInterface.
// A group's client should implement this interface.
type LocalIdentitiesGetter interface {
	LocalIdentities() LocalIdentityInterface
}

// LocalIdentityInterface has methods to work with LocalIdentity resources.
type LocalIdentityInterface interface {
	Create(*auth.LocalIdentity) (*auth.LocalIdentity, error)
	Update(*auth.LocalIdentity) (*auth.LocalIdentity, error)
	UpdateStatus(*auth.LocalIdentity) (*auth.LocalIdentity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*auth.LocalIdentity, error)
	List(opts v1.ListOptions) (*auth.LocalIdentityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.LocalIdentity, err error)
	LocalIdentityExpansion
}

// localIdentities implements LocalIdentityInterface
type localIdentities struct {
	client rest.Interface
}

// newLocalIdentities returns a LocalIdentities
func newLocalIdentities(c *AuthClient) *localIdentities {
	return &localIdentities{
		client: c.RESTClient(),
	}
}

// Get takes name of the localIdentity, and returns the corresponding localIdentity object, and an error if there is any.
func (c *localIdentities) Get(name string, options v1.GetOptions) (result *auth.LocalIdentity, err error) {
	result = &auth.LocalIdentity{}
	err = c.client.Get().
		Resource("localidentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalIdentities that match those selectors.
func (c *localIdentities) List(opts v1.ListOptions) (result *auth.LocalIdentityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &auth.LocalIdentityList{}
	err = c.client.Get().
		Resource("localidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localIdentities.
func (c *localIdentities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("localidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a localIdentity and creates it.  Returns the server's representation of the localIdentity, and an error, if there is any.
func (c *localIdentities) Create(localIdentity *auth.LocalIdentity) (result *auth.LocalIdentity, err error) {
	result = &auth.LocalIdentity{}
	err = c.client.Post().
		Resource("localidentities").
		Body(localIdentity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a localIdentity and updates it. Returns the server's representation of the localIdentity, and an error, if there is any.
func (c *localIdentities) Update(localIdentity *auth.LocalIdentity) (result *auth.LocalIdentity, err error) {
	result = &auth.LocalIdentity{}
	err = c.client.Put().
		Resource("localidentities").
		Name(localIdentity.Name).
		Body(localIdentity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *localIdentities) UpdateStatus(localIdentity *auth.LocalIdentity) (result *auth.LocalIdentity, err error) {
	result = &auth.LocalIdentity{}
	err = c.client.Put().
		Resource("localidentities").
		Name(localIdentity.Name).
		SubResource("status").
		Body(localIdentity).
		Do().
		Into(result)
	return
}

// Delete takes name of the localIdentity and deletes it. Returns an error if one occurs.
func (c *localIdentities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("localidentities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localIdentities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("localidentities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched localIdentity.
func (c *localIdentities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.LocalIdentity, err error) {
	result = &auth.LocalIdentity{}
	err = c.client.Patch(pt).
		Resource("localidentities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
