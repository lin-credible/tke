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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	business "tkestack.io/tke/api/business"
)

// PortalLister helps list Portals.
type PortalLister interface {
	// List lists all Portals in the indexer.
	List(selector labels.Selector) (ret []*business.Portal, err error)
	// Get retrieves the Portal from the index for a given name.
	Get(name string) (*business.Portal, error)
	PortalListerExpansion
}

// portalLister implements the PortalLister interface.
type portalLister struct {
	indexer cache.Indexer
}

// NewPortalLister returns a new PortalLister.
func NewPortalLister(indexer cache.Indexer) PortalLister {
	return &portalLister{indexer: indexer}
}

// List lists all Portals in the indexer.
func (s *portalLister) List(selector labels.Selector) (ret []*business.Portal, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*business.Portal))
	})
	return ret, err
}

// Get retrieves the Portal from the index for a given name.
func (s *portalLister) Get(name string) (*business.Portal, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(business.Resource("portal"), name)
	}
	return obj.(*business.Portal), nil
}
