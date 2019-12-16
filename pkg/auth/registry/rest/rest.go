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

package rest

import (
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	restclient "k8s.io/client-go/rest"
	"tkestack.io/tke/api/auth"
	v1 "tkestack.io/tke/api/auth/v1"
	authinternalclient "tkestack.io/tke/api/client/clientset/internalversion/typed/auth/internalversion"
	"tkestack.io/tke/pkg/apiserver/storage"
	"tkestack.io/tke/pkg/auth/authorization/enforcer"
	apikeystorage "tkestack.io/tke/pkg/auth/registry/apikey/storage"
	apisignstorage "tkestack.io/tke/pkg/auth/registry/apisigningkey/storage"
	categorystorage "tkestack.io/tke/pkg/auth/registry/category/storage"
	configmapstorage "tkestack.io/tke/pkg/auth/registry/configmap/storage"
	localidentitystorage "tkestack.io/tke/pkg/auth/registry/localidentity/storage"
	policystorage "tkestack.io/tke/pkg/auth/registry/policy/storage"
	rulestorage "tkestack.io/tke/pkg/auth/registry/rule/storage"

	"tkestack.io/tke/pkg/auth/util/sign"
)

// StorageProvider is a REST type for core resources storage that implement
// RestStorageProvider interface
type StorageProvider struct {
	LoopbackClientConfig *restclient.Config
	PolicyEnforcer       *enforcer.PolicyEnforcer
	PrivilegedUsername   string
}

// Implement RESTStorageProvider
var _ storage.RESTStorageProvider = &StorageProvider{}

// NewRESTStorage is a factory constructor to creates and returns the APIGroupInfo
func (s *StorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericserver.APIGroupInfo, bool) {
	apiGroupInfo := genericserver.NewDefaultAPIGroupInfo(auth.GroupName, auth.Scheme, auth.ParameterCodec, auth.Codecs)

	if apiResourceConfigSource.VersionEnabled(v1.SchemeGroupVersion) {
		apiGroupInfo.VersionedResourcesStorageMap[v1.SchemeGroupVersion.Version] = s.v1Storage(apiResourceConfigSource, restOptionsGetter, s.LoopbackClientConfig)
	}

	return apiGroupInfo, true
}

// GroupName return the api group name
func (*StorageProvider) GroupName() string {
	return auth.GroupName
}

func (s *StorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter, loopbackClientConfig *restclient.Config) map[string]rest.Storage {
	authClient := authinternalclient.NewForConfigOrDie(loopbackClientConfig)
	storageMap := make(map[string]rest.Storage)
	{

		configMapREST := configmapstorage.NewStorage(restOptionsGetter)
		storageMap["configmaps"] = configMapREST.ConfigMap

		localIdentityRest := localidentitystorage.NewStorage(restOptionsGetter, authClient, s.PolicyEnforcer, s.PrivilegedUsername)
		storageMap["localidentities"] = localIdentityRest.LocalIdentity
		storageMap["localidentities/status"] = localIdentityRest.Status

		keySigner := sign.NewGenericKeySigner(authClient)
		apiKeyRest := apikeystorage.NewStorage(restOptionsGetter, authClient, keySigner, s.PrivilegedUsername)
		storageMap["apikeys"] = apiKeyRest.APIKey
		storageMap["apikeys/password"] = apiKeyRest.Password
		storageMap["apikeys/token"] = apiKeyRest.Token
		storageMap["apikeys/status"] = apiKeyRest.Status

		apiSignRest := apisignstorage.NewStorage(restOptionsGetter)
		storageMap["apisigningkeys"] = apiSignRest

		categoryRest := categorystorage.NewStorage(restOptionsGetter)
		storageMap["categories"] = categoryRest

		policyRest := policystorage.NewStorage(restOptionsGetter, s.PolicyEnforcer, authClient, s.PrivilegedUsername)
		storageMap["policies"] = policyRest.Policy
		storageMap["policies/finalize"] = policyRest.Finalize
		storageMap["policies/status"] = policyRest.Status
		storageMap["policies/binding"] = policyRest.Binding
		storageMap["policies/unbinding"] = policyRest.Unbinding

		ruleRest := rulestorage.NewStorage(restOptionsGetter)
		storageMap["rules"] = ruleRest.Rule
	}

	return storageMap
}
