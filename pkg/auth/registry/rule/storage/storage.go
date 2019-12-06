/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the “License”); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an “AS IS” BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package rule

import (
	"context"

	"k8s.io/apiserver/pkg/registry/rest"
	"tkestack.io/tke/pkg/apiserver/authentication"

	"tkestack.io/tke/pkg/auth/registry/rule"
	"tkestack.io/tke/pkg/util/log"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	apiserverutil "tkestack.io/tke/pkg/apiserver/util"

	"tkestack.io/tke/api/auth"
)

// Storage includes storage for policies and all sub resources.
type Storage struct {
	Rule *REST
}

// NewStorage returns a Storage object that will work against policies.
func NewStorage(optsGetter generic.RESTOptionsGetter) *Storage {
	strategy := rule.NewStrategy()
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &auth.Rule{} },
		NewListFunc:              func() runtime.Object { return &auth.RuleList{} },
		DefaultQualifiedResource: auth.Resource("rules"),
		PredicateFunc:            rule.MatchRule,

		CreateStrategy: strategy,
		DeleteStrategy: strategy,
	}
	options := &generic.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    rule.GetAttrs,
	}

	if err := store.CompleteWithOptions(options); err != nil {
		log.Panic("Failed to create rule etcd rest storage", log.Err(err))
	}

	return &Storage{
		Rule: &REST{store},
	}
}

// REST implements a RESTStorage for clusters against etcd.
type REST struct {
	*registry.Store
}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"rul"}
}

// Create inserts a new item according to the unique key from the object.
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	if tenantID != "" {
		return nil, apierrors.NewMethodNotSupported(auth.Resource("rules"), "create")
	}
	return r.Store.Create(ctx, obj, createValidation, options)
}

// Delete enforces life-cycle rules for policy termination
func (r *REST) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	if tenantID != "" {
		return nil, false, apierrors.NewMethodNotSupported(auth.Resource("rules"), "delete")
	}
	return r.Store.Delete(ctx, name, deleteValidation, options)
}

// DeleteCollection selects all resources in the storage matching given 'listOptions'
// and deletes them.
func (r *REST) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions, listOptions *metainternal.ListOptions) (runtime.Object, error) {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	if tenantID != "" {
		return nil, apierrors.NewMethodNotSupported(auth.Resource("rules"), "delete collection")
	}
	return r.Store.DeleteCollection(ctx, deleteValidation, options, listOptions)
}

// Update alters the object subset of an object.
func (r *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	if tenantID != "" {
		return nil, false, apierrors.NewMethodNotSupported(auth.Resource("rules"), "update")
	}
	return r.Store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

// List selects resources in the storage which match to the selector. 'options' can be nil.
func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	wrappedOptions := apiserverutil.PredicateListOptions(ctx, options)
	return r.Store.List(ctx, wrappedOptions)
}
