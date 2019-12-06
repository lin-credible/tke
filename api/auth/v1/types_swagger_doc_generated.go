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

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_APIKey = map[string]string{
	"":     "APIKey contains expiration time used to apply the api key.",
	"spec": "Spec defines the desired identities of APIkey in this set.",
}

func (APIKey) SwaggerDoc() map[string]string {
	return map_APIKey
}

var map_APIKeyList = map[string]string{
	"":      "LocalIdentityList is the whole list of all identities.",
	"items": "List of api keys.",
}

func (APIKeyList) SwaggerDoc() map[string]string {
	return map_APIKeyList
}

var map_APIKeyReq = map[string]string{
	"":            "APIKeyReq contains expiration time used to apply the api key.",
	"expire":      "Exipre is required, holds the duration of the api key become invalid. By default, 168h(= seven days)",
	"description": "Description describes api keys usage.",
}

func (APIKeyReq) SwaggerDoc() map[string]string {
	return map_APIKeyReq
}

var map_APIKeyReqPassword = map[string]string{
	"":            "APIKeyReqPassword contains userinfo and expiration time used to apply the api key.",
	"tenantID":    "TenantID for user",
	"username":    "Username",
	"password":    "Password (encoded by base64)",
	"description": "Description describes api keys usage.",
	"expire":      "Expire holds the duration of the api key become invalid. By default, 168h(= seven days)",
}

func (APIKeyReqPassword) SwaggerDoc() map[string]string {
	return map_APIKeyReqPassword
}

var map_APIKeySpec = map[string]string{
	"":            "APIKeySpec is a description of an apiKey.",
	"apiKey":      "APIkey is the jwt token used to authenticate user, and contains user info and sign.",
	"username":    "Username is creator",
	"description": "Description describes api keys usage.",
	"issue_at":    "IssueAt is the created time for api key",
	"expire_at":   "ExpireAt is the expire time for api key",
}

func (APIKeySpec) SwaggerDoc() map[string]string {
	return map_APIKeySpec
}

var map_APIKeyStatus = map[string]string{
	"":         "APIKeyStatus is a description of an api key status.",
	"disabled": "Disabled represents whether the apikey has been disabled.",
}

func (APIKeyStatus) SwaggerDoc() map[string]string {
	return map_APIKeyStatus
}

var map_APISigningKey = map[string]string{
	"": "APISigningKey hold encryption and signing key.",
}

func (APISigningKey) SwaggerDoc() map[string]string {
	return map_APISigningKey
}

var map_APISigningKeyList = map[string]string{
	"":      "APISigningKeyList is the whole list of all signing key.",
	"items": "List of keys.",
}

func (APISigningKeyList) SwaggerDoc() map[string]string {
	return map_APISigningKeyList
}

var map_Action = map[string]string{
	"":            "Action defines a action verb for authorization.",
	"name":        "Name represents user access review request verb.",
	"description": "Description describes the action.",
}

func (Action) SwaggerDoc() map[string]string {
	return map_Action
}

var map_Binding = map[string]string{
	"":         "Binding references the objects a policy applies to, but does not contain it.",
	"subjects": "Subjects holds references to the objects the policy applies to.",
}

func (Binding) SwaggerDoc() map[string]string {
	return map_Binding
}

var map_Category = map[string]string{
	"": "Category defines a category of actions for policy.",
}

func (Category) SwaggerDoc() map[string]string {
	return map_Category
}

var map_CategoryList = map[string]string{
	"":      "CategoryList is the whole list of policy Category.",
	"items": "List of category.",
}

func (CategoryList) SwaggerDoc() map[string]string {
	return map_CategoryList
}

var map_CategorySpec = map[string]string{
	"categoryName": "CategoryName identifies action category",
	"displayName":  "DisplayName used to display category name",
	"actions":      "Actions represents a series of actions work on the policy category",
}

func (CategorySpec) SwaggerDoc() map[string]string {
	return map_CategorySpec
}

var map_ConfigMap = map[string]string{
	"":           "ConfigMap holds configuration data for tke to consume.",
	"data":       "Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.",
	"binaryData": "BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process.",
}

func (ConfigMap) SwaggerDoc() map[string]string {
	return map_ConfigMap
}

var map_ConfigMapList = map[string]string{
	"":      "ConfigMapList is a resource containing a list of ConfigMap objects.",
	"items": "Items is the list of ConfigMaps.",
}

func (ConfigMapList) SwaggerDoc() map[string]string {
	return map_ConfigMapList
}

var map_LocalIdentity = map[string]string{
	"":     "LocalIdentity is an object that contains the metadata about identify used to login to TKE.",
	"spec": "Spec defines the desired identities of identity in this set.",
}

func (LocalIdentity) SwaggerDoc() map[string]string {
	return map_LocalIdentity
}

var map_LocalIdentityList = map[string]string{
	"":      "LocalIdentityList is the whole list of all identities.",
	"items": "List of identities.",
}

func (LocalIdentityList) SwaggerDoc() map[string]string {
	return map_LocalIdentityList
}

var map_LocalIdentitySpec = map[string]string{
	"": "LocalIdentitySpec is a description of an identity.",
}

func (LocalIdentitySpec) SwaggerDoc() map[string]string {
	return map_LocalIdentitySpec
}

var map_LocalIdentityStatus = map[string]string{
	"":               "LocalIdentityStatus is a description of an identity status.",
	"LastUpdateTime": "The last time the local identity was updated.",
}

func (LocalIdentityStatus) SwaggerDoc() map[string]string {
	return map_LocalIdentityStatus
}

var map_Policy = map[string]string{
	"":     "Policy represents a policy document for access control.",
	"spec": "Spec defines the desired identities of policy document in this set.",
}

func (Policy) SwaggerDoc() map[string]string {
	return map_Policy
}

var map_PolicyList = map[string]string{
	"":      "PolicyList is the whole list of all policies.",
	"items": "List of policies.",
}

func (PolicyList) SwaggerDoc() map[string]string {
	return map_PolicyList
}

var map_PolicySpec = map[string]string{
	"": "PolicySpec is a description of a policy.",
}

func (PolicySpec) SwaggerDoc() map[string]string {
	return map_PolicySpec
}

var map_PolicyStatus = map[string]string{
	"":         "PolicyStatus represents information about the status of a policy.",
	"subjects": "Subjects represents the objects the policy applies to.",
}

func (PolicyStatus) SwaggerDoc() map[string]string {
	return map_PolicyStatus
}

var map_Rule = map[string]string{
	"":     "Rule represents a rule document for access control.",
	"spec": "Spec defines the desired identities of policy document in this set.",
}

func (Rule) SwaggerDoc() map[string]string {
	return map_Rule
}

var map_RuleList = map[string]string{
	"":      "RuleList is the whole list of all policies.",
	"items": "List of rules.",
}

func (RuleList) SwaggerDoc() map[string]string {
	return map_RuleList
}

var map_RuleSpec = map[string]string{
	"": "RuleSpec is a description of a policy.",
}

func (RuleSpec) SwaggerDoc() map[string]string {
	return map_RuleSpec
}

var map_Statement = map[string]string{
	"":       "Statement defines a series of action on resource can be done or not.",
	"effect": "Effect indicates action on the resource is allowed or not, can be \"allow\" or \"deny\"",
}

func (Statement) SwaggerDoc() map[string]string {
	return map_Statement
}

var map_Subject = map[string]string{
	"": "Subject references a user can specify by id or name.",
}

func (Subject) SwaggerDoc() map[string]string {
	return map_Subject
}

// AUTO-GENERATED FUNCTIONS END HERE
