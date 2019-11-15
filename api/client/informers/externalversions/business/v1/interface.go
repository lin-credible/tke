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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "tkestack.io/tke/api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ChartGroups returns a ChartGroupInformer.
	ChartGroups() ChartGroupInformer
	// ConfigMaps returns a ConfigMapInformer.
	ConfigMaps() ConfigMapInformer
	// ImageNamespaces returns a ImageNamespaceInformer.
	ImageNamespaces() ImageNamespaceInformer
	// Namespaces returns a NamespaceInformer.
	Namespaces() NamespaceInformer
	// Platforms returns a PlatformInformer.
	Platforms() PlatformInformer
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ChartGroups returns a ChartGroupInformer.
func (v *version) ChartGroups() ChartGroupInformer {
	return &chartGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigMaps returns a ConfigMapInformer.
func (v *version) ConfigMaps() ConfigMapInformer {
	return &configMapInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ImageNamespaces returns a ImageNamespaceInformer.
func (v *version) ImageNamespaces() ImageNamespaceInformer {
	return &imageNamespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Namespaces returns a NamespaceInformer.
func (v *version) Namespaces() NamespaceInformer {
	return &namespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Platforms returns a PlatformInformer.
func (v *version) Platforms() PlatformInformer {
	return &platformInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
