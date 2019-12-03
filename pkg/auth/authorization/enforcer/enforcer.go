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

package enforcer

import (
	"time"

	"github.com/casbin/casbin/v2"
	"tkestack.io/tke/pkg/auth/registry"
)

// PolicyEnforcer contains a casbin Enforcer to handle policy crud with the casbin
type PolicyEnforcer struct {
	Enforcer *casbin.SyncedEnforcer
	registry *registry.Registry
}

// NewPolicyEnforcer creates a policy Enforcer object with casbin Enforcer.
func NewPolicyEnforcer(enforcer *casbin.SyncedEnforcer, registry *registry.Registry) *PolicyEnforcer {
	pe := &PolicyEnforcer{Enforcer: enforcer, registry: registry}
	return pe
}

func (pe *PolicyEnforcer) StartAutoLoadPolicy(duration time.Duration) {
	pe.Enforcer.StartAutoLoadPolicy(duration)
}
