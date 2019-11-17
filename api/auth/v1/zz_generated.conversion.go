// +build !ignore_autogenerated

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	auth "tkestack.io/tke/api/auth"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*APIKey)(nil), (*auth.APIKey)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIKey_To_auth_APIKey(a.(*APIKey), b.(*auth.APIKey), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APIKey)(nil), (*APIKey)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APIKey_To_v1_APIKey(a.(*auth.APIKey), b.(*APIKey), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIKeyList)(nil), (*auth.APIKeyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIKeyList_To_auth_APIKeyList(a.(*APIKeyList), b.(*auth.APIKeyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APIKeyList)(nil), (*APIKeyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APIKeyList_To_v1_APIKeyList(a.(*auth.APIKeyList), b.(*APIKeyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIKeyReq)(nil), (*auth.APIKeyReq)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIKeyReq_To_auth_APIKeyReq(a.(*APIKeyReq), b.(*auth.APIKeyReq), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APIKeyReq)(nil), (*APIKeyReq)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APIKeyReq_To_v1_APIKeyReq(a.(*auth.APIKeyReq), b.(*APIKeyReq), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIKeyReqPassword)(nil), (*auth.APIKeyReqPassword)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIKeyReqPassword_To_auth_APIKeyReqPassword(a.(*APIKeyReqPassword), b.(*auth.APIKeyReqPassword), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APIKeyReqPassword)(nil), (*APIKeyReqPassword)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APIKeyReqPassword_To_v1_APIKeyReqPassword(a.(*auth.APIKeyReqPassword), b.(*APIKeyReqPassword), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIKeySpec)(nil), (*auth.APIKeySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIKeySpec_To_auth_APIKeySpec(a.(*APIKeySpec), b.(*auth.APIKeySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APIKeySpec)(nil), (*APIKeySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APIKeySpec_To_v1_APIKeySpec(a.(*auth.APIKeySpec), b.(*APIKeySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIKeyStatus)(nil), (*auth.APIKeyStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIKeyStatus_To_auth_APIKeyStatus(a.(*APIKeyStatus), b.(*auth.APIKeyStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APIKeyStatus)(nil), (*APIKeyStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APIKeyStatus_To_v1_APIKeyStatus(a.(*auth.APIKeyStatus), b.(*APIKeyStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APISigningKey)(nil), (*auth.APISigningKey)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APISigningKey_To_auth_APISigningKey(a.(*APISigningKey), b.(*auth.APISigningKey), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APISigningKey)(nil), (*APISigningKey)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APISigningKey_To_v1_APISigningKey(a.(*auth.APISigningKey), b.(*APISigningKey), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APISigningKeyList)(nil), (*auth.APISigningKeyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APISigningKeyList_To_auth_APISigningKeyList(a.(*APISigningKeyList), b.(*auth.APISigningKeyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.APISigningKeyList)(nil), (*APISigningKeyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_APISigningKeyList_To_v1_APISigningKeyList(a.(*auth.APISigningKeyList), b.(*APISigningKeyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMap)(nil), (*auth.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMap_To_auth_ConfigMap(a.(*ConfigMap), b.(*auth.ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.ConfigMap)(nil), (*ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_ConfigMap_To_v1_ConfigMap(a.(*auth.ConfigMap), b.(*ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMapList)(nil), (*auth.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMapList_To_auth_ConfigMapList(a.(*ConfigMapList), b.(*auth.ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.ConfigMapList)(nil), (*ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_ConfigMapList_To_v1_ConfigMapList(a.(*auth.ConfigMapList), b.(*ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalIdentity)(nil), (*auth.LocalIdentity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalIdentity_To_auth_LocalIdentity(a.(*LocalIdentity), b.(*auth.LocalIdentity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.LocalIdentity)(nil), (*LocalIdentity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_LocalIdentity_To_v1_LocalIdentity(a.(*auth.LocalIdentity), b.(*LocalIdentity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalIdentityList)(nil), (*auth.LocalIdentityList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalIdentityList_To_auth_LocalIdentityList(a.(*LocalIdentityList), b.(*auth.LocalIdentityList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.LocalIdentityList)(nil), (*LocalIdentityList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_LocalIdentityList_To_v1_LocalIdentityList(a.(*auth.LocalIdentityList), b.(*LocalIdentityList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalIdentitySpec)(nil), (*auth.LocalIdentitySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalIdentitySpec_To_auth_LocalIdentitySpec(a.(*LocalIdentitySpec), b.(*auth.LocalIdentitySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.LocalIdentitySpec)(nil), (*LocalIdentitySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_LocalIdentitySpec_To_v1_LocalIdentitySpec(a.(*auth.LocalIdentitySpec), b.(*LocalIdentitySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalIdentityStatus)(nil), (*auth.LocalIdentityStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalIdentityStatus_To_auth_LocalIdentityStatus(a.(*LocalIdentityStatus), b.(*auth.LocalIdentityStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*auth.LocalIdentityStatus)(nil), (*LocalIdentityStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_auth_LocalIdentityStatus_To_v1_LocalIdentityStatus(a.(*auth.LocalIdentityStatus), b.(*LocalIdentityStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_APIKey_To_auth_APIKey(in *APIKey, out *auth.APIKey, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_APIKeySpec_To_auth_APIKeySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_APIKeyStatus_To_auth_APIKeyStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_APIKey_To_auth_APIKey is an autogenerated conversion function.
func Convert_v1_APIKey_To_auth_APIKey(in *APIKey, out *auth.APIKey, s conversion.Scope) error {
	return autoConvert_v1_APIKey_To_auth_APIKey(in, out, s)
}

func autoConvert_auth_APIKey_To_v1_APIKey(in *auth.APIKey, out *APIKey, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_auth_APIKeySpec_To_v1_APIKeySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_auth_APIKeyStatus_To_v1_APIKeyStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_auth_APIKey_To_v1_APIKey is an autogenerated conversion function.
func Convert_auth_APIKey_To_v1_APIKey(in *auth.APIKey, out *APIKey, s conversion.Scope) error {
	return autoConvert_auth_APIKey_To_v1_APIKey(in, out, s)
}

func autoConvert_v1_APIKeyList_To_auth_APIKeyList(in *APIKeyList, out *auth.APIKeyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]auth.APIKey)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_APIKeyList_To_auth_APIKeyList is an autogenerated conversion function.
func Convert_v1_APIKeyList_To_auth_APIKeyList(in *APIKeyList, out *auth.APIKeyList, s conversion.Scope) error {
	return autoConvert_v1_APIKeyList_To_auth_APIKeyList(in, out, s)
}

func autoConvert_auth_APIKeyList_To_v1_APIKeyList(in *auth.APIKeyList, out *APIKeyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]APIKey)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_auth_APIKeyList_To_v1_APIKeyList is an autogenerated conversion function.
func Convert_auth_APIKeyList_To_v1_APIKeyList(in *auth.APIKeyList, out *APIKeyList, s conversion.Scope) error {
	return autoConvert_auth_APIKeyList_To_v1_APIKeyList(in, out, s)
}

func autoConvert_v1_APIKeyReq_To_auth_APIKeyReq(in *APIKeyReq, out *auth.APIKeyReq, s conversion.Scope) error {
	out.Expire = in.Expire
	out.Description = in.Description
	return nil
}

// Convert_v1_APIKeyReq_To_auth_APIKeyReq is an autogenerated conversion function.
func Convert_v1_APIKeyReq_To_auth_APIKeyReq(in *APIKeyReq, out *auth.APIKeyReq, s conversion.Scope) error {
	return autoConvert_v1_APIKeyReq_To_auth_APIKeyReq(in, out, s)
}

func autoConvert_auth_APIKeyReq_To_v1_APIKeyReq(in *auth.APIKeyReq, out *APIKeyReq, s conversion.Scope) error {
	out.Expire = in.Expire
	out.Description = in.Description
	return nil
}

// Convert_auth_APIKeyReq_To_v1_APIKeyReq is an autogenerated conversion function.
func Convert_auth_APIKeyReq_To_v1_APIKeyReq(in *auth.APIKeyReq, out *APIKeyReq, s conversion.Scope) error {
	return autoConvert_auth_APIKeyReq_To_v1_APIKeyReq(in, out, s)
}

func autoConvert_v1_APIKeyReqPassword_To_auth_APIKeyReqPassword(in *APIKeyReqPassword, out *auth.APIKeyReqPassword, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.Username = in.Username
	out.Password = in.Password
	out.Description = in.Description
	out.Expire = in.Expire
	return nil
}

// Convert_v1_APIKeyReqPassword_To_auth_APIKeyReqPassword is an autogenerated conversion function.
func Convert_v1_APIKeyReqPassword_To_auth_APIKeyReqPassword(in *APIKeyReqPassword, out *auth.APIKeyReqPassword, s conversion.Scope) error {
	return autoConvert_v1_APIKeyReqPassword_To_auth_APIKeyReqPassword(in, out, s)
}

func autoConvert_auth_APIKeyReqPassword_To_v1_APIKeyReqPassword(in *auth.APIKeyReqPassword, out *APIKeyReqPassword, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.Username = in.Username
	out.Password = in.Password
	out.Description = in.Description
	out.Expire = in.Expire
	return nil
}

// Convert_auth_APIKeyReqPassword_To_v1_APIKeyReqPassword is an autogenerated conversion function.
func Convert_auth_APIKeyReqPassword_To_v1_APIKeyReqPassword(in *auth.APIKeyReqPassword, out *APIKeyReqPassword, s conversion.Scope) error {
	return autoConvert_auth_APIKeyReqPassword_To_v1_APIKeyReqPassword(in, out, s)
}

func autoConvert_v1_APIKeySpec_To_auth_APIKeySpec(in *APIKeySpec, out *auth.APIKeySpec, s conversion.Scope) error {
	out.APIkey = in.APIkey
	out.TenantID = in.TenantID
	out.Username = in.Username
	out.Description = in.Description
	out.IssueAt = in.IssueAt
	out.ExpireAt = in.ExpireAt
	return nil
}

// Convert_v1_APIKeySpec_To_auth_APIKeySpec is an autogenerated conversion function.
func Convert_v1_APIKeySpec_To_auth_APIKeySpec(in *APIKeySpec, out *auth.APIKeySpec, s conversion.Scope) error {
	return autoConvert_v1_APIKeySpec_To_auth_APIKeySpec(in, out, s)
}

func autoConvert_auth_APIKeySpec_To_v1_APIKeySpec(in *auth.APIKeySpec, out *APIKeySpec, s conversion.Scope) error {
	out.APIkey = in.APIkey
	out.TenantID = in.TenantID
	out.Username = in.Username
	out.Description = in.Description
	out.IssueAt = in.IssueAt
	out.ExpireAt = in.ExpireAt
	return nil
}

// Convert_auth_APIKeySpec_To_v1_APIKeySpec is an autogenerated conversion function.
func Convert_auth_APIKeySpec_To_v1_APIKeySpec(in *auth.APIKeySpec, out *APIKeySpec, s conversion.Scope) error {
	return autoConvert_auth_APIKeySpec_To_v1_APIKeySpec(in, out, s)
}

func autoConvert_v1_APIKeyStatus_To_auth_APIKeyStatus(in *APIKeyStatus, out *auth.APIKeyStatus, s conversion.Scope) error {
	out.Disabled = in.Disabled
	return nil
}

// Convert_v1_APIKeyStatus_To_auth_APIKeyStatus is an autogenerated conversion function.
func Convert_v1_APIKeyStatus_To_auth_APIKeyStatus(in *APIKeyStatus, out *auth.APIKeyStatus, s conversion.Scope) error {
	return autoConvert_v1_APIKeyStatus_To_auth_APIKeyStatus(in, out, s)
}

func autoConvert_auth_APIKeyStatus_To_v1_APIKeyStatus(in *auth.APIKeyStatus, out *APIKeyStatus, s conversion.Scope) error {
	out.Disabled = in.Disabled
	return nil
}

// Convert_auth_APIKeyStatus_To_v1_APIKeyStatus is an autogenerated conversion function.
func Convert_auth_APIKeyStatus_To_v1_APIKeyStatus(in *auth.APIKeyStatus, out *APIKeyStatus, s conversion.Scope) error {
	return autoConvert_auth_APIKeyStatus_To_v1_APIKeyStatus(in, out, s)
}

func autoConvert_v1_APISigningKey_To_auth_APISigningKey(in *APISigningKey, out *auth.APISigningKey, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.SigningKey = *(*[]byte)(unsafe.Pointer(&in.SigningKey))
	out.SigningKeyPub = *(*[]byte)(unsafe.Pointer(&in.SigningKeyPub))
	return nil
}

// Convert_v1_APISigningKey_To_auth_APISigningKey is an autogenerated conversion function.
func Convert_v1_APISigningKey_To_auth_APISigningKey(in *APISigningKey, out *auth.APISigningKey, s conversion.Scope) error {
	return autoConvert_v1_APISigningKey_To_auth_APISigningKey(in, out, s)
}

func autoConvert_auth_APISigningKey_To_v1_APISigningKey(in *auth.APISigningKey, out *APISigningKey, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.SigningKey = *(*[]byte)(unsafe.Pointer(&in.SigningKey))
	out.SigningKeyPub = *(*[]byte)(unsafe.Pointer(&in.SigningKeyPub))
	return nil
}

// Convert_auth_APISigningKey_To_v1_APISigningKey is an autogenerated conversion function.
func Convert_auth_APISigningKey_To_v1_APISigningKey(in *auth.APISigningKey, out *APISigningKey, s conversion.Scope) error {
	return autoConvert_auth_APISigningKey_To_v1_APISigningKey(in, out, s)
}

func autoConvert_v1_APISigningKeyList_To_auth_APISigningKeyList(in *APISigningKeyList, out *auth.APISigningKeyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]auth.APISigningKey)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_APISigningKeyList_To_auth_APISigningKeyList is an autogenerated conversion function.
func Convert_v1_APISigningKeyList_To_auth_APISigningKeyList(in *APISigningKeyList, out *auth.APISigningKeyList, s conversion.Scope) error {
	return autoConvert_v1_APISigningKeyList_To_auth_APISigningKeyList(in, out, s)
}

func autoConvert_auth_APISigningKeyList_To_v1_APISigningKeyList(in *auth.APISigningKeyList, out *APISigningKeyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]APISigningKey)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_auth_APISigningKeyList_To_v1_APISigningKeyList is an autogenerated conversion function.
func Convert_auth_APISigningKeyList_To_v1_APISigningKeyList(in *auth.APISigningKeyList, out *APISigningKeyList, s conversion.Scope) error {
	return autoConvert_auth_APISigningKeyList_To_v1_APISigningKeyList(in, out, s)
}

func autoConvert_v1_ConfigMap_To_auth_ConfigMap(in *ConfigMap, out *auth.ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_v1_ConfigMap_To_auth_ConfigMap is an autogenerated conversion function.
func Convert_v1_ConfigMap_To_auth_ConfigMap(in *ConfigMap, out *auth.ConfigMap, s conversion.Scope) error {
	return autoConvert_v1_ConfigMap_To_auth_ConfigMap(in, out, s)
}

func autoConvert_auth_ConfigMap_To_v1_ConfigMap(in *auth.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_auth_ConfigMap_To_v1_ConfigMap is an autogenerated conversion function.
func Convert_auth_ConfigMap_To_v1_ConfigMap(in *auth.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	return autoConvert_auth_ConfigMap_To_v1_ConfigMap(in, out, s)
}

func autoConvert_v1_ConfigMapList_To_auth_ConfigMapList(in *ConfigMapList, out *auth.ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]auth.ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ConfigMapList_To_auth_ConfigMapList is an autogenerated conversion function.
func Convert_v1_ConfigMapList_To_auth_ConfigMapList(in *ConfigMapList, out *auth.ConfigMapList, s conversion.Scope) error {
	return autoConvert_v1_ConfigMapList_To_auth_ConfigMapList(in, out, s)
}

func autoConvert_auth_ConfigMapList_To_v1_ConfigMapList(in *auth.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_auth_ConfigMapList_To_v1_ConfigMapList is an autogenerated conversion function.
func Convert_auth_ConfigMapList_To_v1_ConfigMapList(in *auth.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	return autoConvert_auth_ConfigMapList_To_v1_ConfigMapList(in, out, s)
}

func autoConvert_v1_LocalIdentity_To_auth_LocalIdentity(in *LocalIdentity, out *auth.LocalIdentity, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_LocalIdentitySpec_To_auth_LocalIdentitySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_LocalIdentityStatus_To_auth_LocalIdentityStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_LocalIdentity_To_auth_LocalIdentity is an autogenerated conversion function.
func Convert_v1_LocalIdentity_To_auth_LocalIdentity(in *LocalIdentity, out *auth.LocalIdentity, s conversion.Scope) error {
	return autoConvert_v1_LocalIdentity_To_auth_LocalIdentity(in, out, s)
}

func autoConvert_auth_LocalIdentity_To_v1_LocalIdentity(in *auth.LocalIdentity, out *LocalIdentity, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_auth_LocalIdentitySpec_To_v1_LocalIdentitySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_auth_LocalIdentityStatus_To_v1_LocalIdentityStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_auth_LocalIdentity_To_v1_LocalIdentity is an autogenerated conversion function.
func Convert_auth_LocalIdentity_To_v1_LocalIdentity(in *auth.LocalIdentity, out *LocalIdentity, s conversion.Scope) error {
	return autoConvert_auth_LocalIdentity_To_v1_LocalIdentity(in, out, s)
}

func autoConvert_v1_LocalIdentityList_To_auth_LocalIdentityList(in *LocalIdentityList, out *auth.LocalIdentityList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]auth.LocalIdentity)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_LocalIdentityList_To_auth_LocalIdentityList is an autogenerated conversion function.
func Convert_v1_LocalIdentityList_To_auth_LocalIdentityList(in *LocalIdentityList, out *auth.LocalIdentityList, s conversion.Scope) error {
	return autoConvert_v1_LocalIdentityList_To_auth_LocalIdentityList(in, out, s)
}

func autoConvert_auth_LocalIdentityList_To_v1_LocalIdentityList(in *auth.LocalIdentityList, out *LocalIdentityList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]LocalIdentity)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_auth_LocalIdentityList_To_v1_LocalIdentityList is an autogenerated conversion function.
func Convert_auth_LocalIdentityList_To_v1_LocalIdentityList(in *auth.LocalIdentityList, out *LocalIdentityList, s conversion.Scope) error {
	return autoConvert_auth_LocalIdentityList_To_v1_LocalIdentityList(in, out, s)
}

func autoConvert_v1_LocalIdentitySpec_To_auth_LocalIdentitySpec(in *LocalIdentitySpec, out *auth.LocalIdentitySpec, s conversion.Scope) error {
	out.Username = in.Username
	out.DisplayName = in.DisplayName
	out.Email = in.Email
	out.PhoneNumber = in.PhoneNumber
	out.HashedPassword = in.HashedPassword
	out.OriginalPassword = in.OriginalPassword
	out.TenantID = in.TenantID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]string)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1_LocalIdentitySpec_To_auth_LocalIdentitySpec is an autogenerated conversion function.
func Convert_v1_LocalIdentitySpec_To_auth_LocalIdentitySpec(in *LocalIdentitySpec, out *auth.LocalIdentitySpec, s conversion.Scope) error {
	return autoConvert_v1_LocalIdentitySpec_To_auth_LocalIdentitySpec(in, out, s)
}

func autoConvert_auth_LocalIdentitySpec_To_v1_LocalIdentitySpec(in *auth.LocalIdentitySpec, out *LocalIdentitySpec, s conversion.Scope) error {
	out.Username = in.Username
	out.DisplayName = in.DisplayName
	out.Email = in.Email
	out.PhoneNumber = in.PhoneNumber
	out.HashedPassword = in.HashedPassword
	out.OriginalPassword = in.OriginalPassword
	out.TenantID = in.TenantID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]string)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_auth_LocalIdentitySpec_To_v1_LocalIdentitySpec is an autogenerated conversion function.
func Convert_auth_LocalIdentitySpec_To_v1_LocalIdentitySpec(in *auth.LocalIdentitySpec, out *LocalIdentitySpec, s conversion.Scope) error {
	return autoConvert_auth_LocalIdentitySpec_To_v1_LocalIdentitySpec(in, out, s)
}

func autoConvert_v1_LocalIdentityStatus_To_auth_LocalIdentityStatus(in *LocalIdentityStatus, out *auth.LocalIdentityStatus, s conversion.Scope) error {
	out.Locked = in.Locked
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

// Convert_v1_LocalIdentityStatus_To_auth_LocalIdentityStatus is an autogenerated conversion function.
func Convert_v1_LocalIdentityStatus_To_auth_LocalIdentityStatus(in *LocalIdentityStatus, out *auth.LocalIdentityStatus, s conversion.Scope) error {
	return autoConvert_v1_LocalIdentityStatus_To_auth_LocalIdentityStatus(in, out, s)
}

func autoConvert_auth_LocalIdentityStatus_To_v1_LocalIdentityStatus(in *auth.LocalIdentityStatus, out *LocalIdentityStatus, s conversion.Scope) error {
	out.Locked = in.Locked
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

// Convert_auth_LocalIdentityStatus_To_v1_LocalIdentityStatus is an autogenerated conversion function.
func Convert_auth_LocalIdentityStatus_To_v1_LocalIdentityStatus(in *auth.LocalIdentityStatus, out *LocalIdentityStatus, s conversion.Scope) error {
	return autoConvert_auth_LocalIdentityStatus_To_v1_LocalIdentityStatus(in, out, s)
}
