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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "tkestack.io/tke/api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Channels returns a ChannelInformer.
	Channels() ChannelInformer
	// ConfigMaps returns a ConfigMapInformer.
	ConfigMaps() ConfigMapInformer
	// Messages returns a MessageInformer.
	Messages() MessageInformer
	// MessageRequests returns a MessageRequestInformer.
	MessageRequests() MessageRequestInformer
	// Receivers returns a ReceiverInformer.
	Receivers() ReceiverInformer
	// ReceiverGroups returns a ReceiverGroupInformer.
	ReceiverGroups() ReceiverGroupInformer
	// Templates returns a TemplateInformer.
	Templates() TemplateInformer
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

// Channels returns a ChannelInformer.
func (v *version) Channels() ChannelInformer {
	return &channelInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ConfigMaps returns a ConfigMapInformer.
func (v *version) ConfigMaps() ConfigMapInformer {
	return &configMapInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Messages returns a MessageInformer.
func (v *version) Messages() MessageInformer {
	return &messageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MessageRequests returns a MessageRequestInformer.
func (v *version) MessageRequests() MessageRequestInformer {
	return &messageRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Receivers returns a ReceiverInformer.
func (v *version) Receivers() ReceiverInformer {
	return &receiverInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ReceiverGroups returns a ReceiverGroupInformer.
func (v *version) ReceiverGroups() ReceiverGroupInformer {
	return &receiverGroupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Templates returns a TemplateInformer.
func (v *version) Templates() TemplateInformer {
	return &templateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
