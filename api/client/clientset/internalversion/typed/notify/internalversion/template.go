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
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	notify "tkestack.io/tke/api/notify"
)

// TemplatesGetter has a method to return a TemplateInterface.
// A group's client should implement this interface.
type TemplatesGetter interface {
	Templates(namespace string) TemplateInterface
}

// TemplateInterface has methods to work with Template resources.
type TemplateInterface interface {
	Create(*notify.Template) (*notify.Template, error)
	Update(*notify.Template) (*notify.Template, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*notify.Template, error)
	List(opts v1.ListOptions) (*notify.TemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.Template, err error)
	TemplateExpansion
}

// templates implements TemplateInterface
type templates struct {
	client rest.Interface
	ns     string
}

// newTemplates returns a Templates
func newTemplates(c *NotifyClient, namespace string) *templates {
	return &templates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the template, and returns the corresponding template object, and an error if there is any.
func (c *templates) Get(name string, options v1.GetOptions) (result *notify.Template, err error) {
	result = &notify.Template{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("templates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Templates that match those selectors.
func (c *templates) List(opts v1.ListOptions) (result *notify.TemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &notify.TemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("templates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested templates.
func (c *templates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("templates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a template and creates it.  Returns the server's representation of the template, and an error, if there is any.
func (c *templates) Create(template *notify.Template) (result *notify.Template, err error) {
	result = &notify.Template{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("templates").
		Body(template).
		Do().
		Into(result)
	return
}

// Update takes the representation of a template and updates it. Returns the server's representation of the template, and an error, if there is any.
func (c *templates) Update(template *notify.Template) (result *notify.Template, err error) {
	result = &notify.Template{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("templates").
		Name(template.Name).
		Body(template).
		Do().
		Into(result)
	return
}

// Delete takes name of the template and deletes it. Returns an error if one occurs.
func (c *templates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("templates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *templates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("templates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched template.
func (c *templates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.Template, err error) {
	result = &notify.Template{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("templates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
