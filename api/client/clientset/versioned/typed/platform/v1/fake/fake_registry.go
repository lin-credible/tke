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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeRegistries implements RegistryInterface
type FakeRegistries struct {
	Fake *FakePlatformV1
}

var registriesResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "registries"}

var registriesKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "Registry"}

// Get takes name of the registry, and returns the corresponding registry object, and an error if there is any.
func (c *FakeRegistries) Get(name string, options v1.GetOptions) (result *platformv1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(registriesResource, name), &platformv1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Registry), err
}

// List takes label and field selectors, and returns the list of Registries that match those selectors.
func (c *FakeRegistries) List(opts v1.ListOptions) (result *platformv1.RegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(registriesResource, registriesKind, opts), &platformv1.RegistryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.RegistryList{ListMeta: obj.(*platformv1.RegistryList).ListMeta}
	for _, item := range obj.(*platformv1.RegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested registries.
func (c *FakeRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(registriesResource, opts))
}

// Create takes the representation of a registry and creates it.  Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Create(registry *platformv1.Registry) (result *platformv1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(registriesResource, registry), &platformv1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Registry), err
}

// Update takes the representation of a registry and updates it. Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Update(registry *platformv1.Registry) (result *platformv1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(registriesResource, registry), &platformv1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Registry), err
}

// Delete takes name of the registry and deletes it. Returns an error if one occurs.
func (c *FakeRegistries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(registriesResource, name), &platformv1.Registry{})
	return err
}

// Patch applies the patch and returns the patched registry.
func (c *FakeRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(registriesResource, name, pt, data, subresources...), &platformv1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Registry), err
}
