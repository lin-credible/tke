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
var map_AddonSpec = map[string]string{
	"": "AddonSpec describes the attributes on a Addon.",
}

func (AddonSpec) SwaggerDoc() map[string]string {
	return map_AddonSpec
}

var map_CSIOperator = map[string]string{
	"":     "CSIOperator is a operator to manages CSI external components.",
	"spec": "Spec defines the desired identities of storage operator.",
}

func (CSIOperator) SwaggerDoc() map[string]string {
	return map_CSIOperator
}

var map_CSIOperatorList = map[string]string{
	"":      "CSIOperatorList is the whole list of all storage operators which owned by a tenant.",
	"items": "List of storage operators.",
}

func (CSIOperatorList) SwaggerDoc() map[string]string {
	return map_CSIOperatorList
}

var map_CSIOperatorSpec = map[string]string{
	"":        "CSIOperatorSpec describes the attributes of a storage operator.",
	"version": "Version of the CSI operator.",
}

func (CSIOperatorSpec) SwaggerDoc() map[string]string {
	return map_CSIOperatorSpec
}

var map_CSIOperatorStatus = map[string]string{
	"":                            "CSIOperatorStatus is information about the current status of a storage operator.",
	"storageVendorVersion":        "StorageVendorVersion will be set to the config version of the storage vendor.",
	"phase":                       "Phase is the current lifecycle phase of the tapp controller of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (CSIOperatorStatus) SwaggerDoc() map[string]string {
	return map_CSIOperatorStatus
}

var map_CSIProxyOptions = map[string]string{
	"": "CSIProxyOptions is the query options to a kube-apiserver proxy call for CSI crd object.",
}

func (CSIProxyOptions) SwaggerDoc() map[string]string {
	return map_CSIProxyOptions
}

var map_Cluster = map[string]string{
	"":     "Cluster is a Kubernetes cluster in",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (Cluster) SwaggerDoc() map[string]string {
	return map_Cluster
}

var map_ClusterAddon = map[string]string{
	"":     "ClusterAddon contains the Addon component for the current kubernetes cluster",
	"spec": "Spec defines the desired identities of addons in this set.",
}

func (ClusterAddon) SwaggerDoc() map[string]string {
	return map_ClusterAddon
}

var map_ClusterAddonList = map[string]string{
	"":      "ClusterAddonList is the whole list of all ClusterAddon.",
	"items": "List of ClusterAddon",
}

func (ClusterAddonList) SwaggerDoc() map[string]string {
	return map_ClusterAddonList
}

var map_ClusterAddonSpec = map[string]string{
	"":        "ClusterAddonSpec indicates the specifications of the ClusterAddon.",
	"type":    "Addon type, one of Helm, PersistentEvent or LogCollector etc.",
	"level":   "AddonLevel is level of cluster addon.",
	"version": "Version",
}

func (ClusterAddonSpec) SwaggerDoc() map[string]string {
	return map_ClusterAddonSpec
}

var map_ClusterAddonStatus = map[string]string{
	"":       "ClusterAddonStatus is information about the current status of a ClusterAddon.",
	"phase":  "Phase is the current lifecycle phase of the addon of cluster.",
	"reason": "Reason is a brief CamelCase string that describes any failure.",
}

func (ClusterAddonStatus) SwaggerDoc() map[string]string {
	return map_ClusterAddonStatus
}

var map_ClusterAddonType = map[string]string{
	"":              "ClusterAddonType records the all addons of cluster available.",
	"type":          "Addon type, one of Helm, PersistentEvent or LogCollector etc.",
	"level":         "AddonLevel is level of cluster addon.",
	"latestVersion": "LatestVersion is latest version of the addon.",
	"description":   "Description is desc of the addon.",
}

func (ClusterAddonType) SwaggerDoc() map[string]string {
	return map_ClusterAddonType
}

var map_ClusterAddonTypeList = map[string]string{
	"": "ClusterAddonTypeList is a resource containing a list of ClusterAddonType objects.",
}

func (ClusterAddonTypeList) SwaggerDoc() map[string]string {
	return map_ClusterAddonTypeList
}

var map_ClusterAddress = map[string]string{
	"":     "ClusterAddress contains information for the cluster's address.",
	"type": "Cluster address type, one of Public, ExternalIP or InternalIP.",
	"host": "The cluster address.",
}

func (ClusterAddress) SwaggerDoc() map[string]string {
	return map_ClusterAddress
}

var map_ClusterApplyOptions = map[string]string{
	"": "ClusterApplyOptions is the query options to a kube-apiserver proxy call for cluster object.",
}

func (ClusterApplyOptions) SwaggerDoc() map[string]string {
	return map_ClusterApplyOptions
}

var map_ClusterComponent = map[string]string{
	"": "ClusterComponent records the number of copies of each component of the cluster master.",
}

func (ClusterComponent) SwaggerDoc() map[string]string {
	return map_ClusterComponent
}

var map_ClusterComponentReplicas = map[string]string{
	"": "ClusterComponentReplicas records the number of copies of each state of each component of the cluster master.",
}

func (ClusterComponentReplicas) SwaggerDoc() map[string]string {
	return map_ClusterComponentReplicas
}

var map_ClusterCondition = map[string]string{
	"":                   "ClusterCondition contains details for the current condition of this cluster.",
	"type":               "Type is the type of the condition.",
	"status":             "Status is the status of the condition. Can be True, False, Unknown.",
	"lastProbeTime":      "Last time we probed the condition.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "Unique, one-word, CamelCase reason for the condition's last transition.",
	"message":            "Human-readable message indicating details about last transition.",
}

func (ClusterCondition) SwaggerDoc() map[string]string {
	return map_ClusterCondition
}

var map_ClusterCredential = map[string]string{
	"":                  "ClusterCredential records the credential information needed to access the cluster.",
	"etcdCACert":        "For TKE in global reuse",
	"etcdAPIClientCert": "For TKE in global reuse",
	"etcdAPIClientKey":  "For TKE in global reuse",
	"caCert":            "For connect the cluster",
	"clientCert":        "For kube-apiserver X509 auth",
	"clientKey":         "For kube-apiserver X509 auth",
	"token":             "For kube-apiserver token auth",
	"bootstrapToken":    "For kubeadm init or join",
	"certificateKey":    "For kubeadm init or join",
}

func (ClusterCredential) SwaggerDoc() map[string]string {
	return map_ClusterCredential
}

var map_ClusterCredentialList = map[string]string{
	"":      "ClusterCredentialList is the whole list of all ClusterCredential which owned by a tenant.",
	"items": "List of clusters",
}

func (ClusterCredentialList) SwaggerDoc() map[string]string {
	return map_ClusterCredentialList
}

var map_ClusterFeature = map[string]string{
	"": "ClusterFeature records the features that are enabled by the cluster.",
}

func (ClusterFeature) SwaggerDoc() map[string]string {
	return map_ClusterFeature
}

var map_ClusterList = map[string]string{
	"":      "ClusterList is the whole list of all clusters which owned by a tenant.",
	"items": "List of clusters",
}

func (ClusterList) SwaggerDoc() map[string]string {
	return map_ClusterList
}

var map_ClusterMachine = map[string]string{
	"": "ClusterMachine is the master machine definition of cluster.",
}

func (ClusterMachine) SwaggerDoc() map[string]string {
	return map_ClusterMachine
}

var map_ClusterProperty = map[string]string{
	"": "ClusterProperty records the attribute information of the cluster.",
}

func (ClusterProperty) SwaggerDoc() map[string]string {
	return map_ClusterProperty
}

var map_ClusterResource = map[string]string{
	"":            "ClusterResource records the current available and maximum resource quota information for the cluster.",
	"capacity":    "Capacity represents the total resources of a cluster.",
	"allocatable": "Allocatable represents the resources of a cluster that are available for scheduling. Defaults to Capacity.",
}

func (ClusterResource) SwaggerDoc() map[string]string {
	return map_ClusterResource
}

var map_ClusterSpec = map[string]string{
	"":           "ClusterSpec is a description of a cluster.",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
	"dnsDomain":  "DNSDomain is the dns domain used by k8s services. Defaults to \"cluster.local\".",
}

func (ClusterSpec) SwaggerDoc() map[string]string {
	return map_ClusterSpec
}

var map_ClusterStatus = map[string]string{
	"":          "ClusterStatus represents information about the status of a cluster.",
	"message":   "A human readable message indicating details about why the cluster is in this condition.",
	"reason":    "A brief CamelCase message indicating details about why the cluster is in this state.",
	"addresses": "List of addresses reachable to the cluster.",
}

func (ClusterStatus) SwaggerDoc() map[string]string {
	return map_ClusterStatus
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

var map_CronHPA = map[string]string{
	"":     "CronHPA is a new kubernetes workload.",
	"spec": "Spec defines the desired identities of CronHPA.",
}

func (CronHPA) SwaggerDoc() map[string]string {
	return map_CronHPA
}

var map_CronHPAList = map[string]string{
	"":      "CronHPAList is the whole list of all CronHPAs which owned by a tenant.",
	"items": "List of CronHPAs",
}

func (CronHPAList) SwaggerDoc() map[string]string {
	return map_CronHPAList
}

var map_CronHPAProxyOptions = map[string]string{
	"": "CronHPAProxyOptions is the query options to a kube-apiserver proxy call.",
}

func (CronHPAProxyOptions) SwaggerDoc() map[string]string {
	return map_CronHPAProxyOptions
}

var map_CronHPASpec = map[string]string{
	"": "CronHPASpec describes the attributes on a CronHPA.",
}

func (CronHPASpec) SwaggerDoc() map[string]string {
	return map_CronHPASpec
}

var map_CronHPAStatus = map[string]string{
	"":                            "CronHPAStatus is information about the current status of a CronHPA.",
	"phase":                       "Phase is the current lifecycle phase of the CronHPA of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (CronHPAStatus) SwaggerDoc() map[string]string {
	return map_CronHPAStatus
}

var map_GPUManager = map[string]string{
	"":     "GPUManager is a kind of device plugin for kubelet to help manage GPUs.",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (GPUManager) SwaggerDoc() map[string]string {
	return map_GPUManager
}

var map_GPUManagerList = map[string]string{
	"":      "GPUManagerList is the whole list of all GPUManager which owned by a tenant.",
	"items": "List of GPUManagers",
}

func (GPUManagerList) SwaggerDoc() map[string]string {
	return map_GPUManagerList
}

var map_GPUManagerSpec = map[string]string{
	"": "GPUManagerSpec describes the attributes of a GPUManager.",
}

func (GPUManagerSpec) SwaggerDoc() map[string]string {
	return map_GPUManagerSpec
}

var map_GPUManagerStatus = map[string]string{
	"":           "GPUManagerStatus is information about the current status of a GPUManager.",
	"phase":      "Phase is the current lifecycle phase of the GPUManager of cluster.",
	"reason":     "Reason is a brief CamelCase string that describes any failure.",
	"retryCount": "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
}

func (GPUManagerStatus) SwaggerDoc() map[string]string {
	return map_GPUManagerStatus
}

var map_Helm = map[string]string{
	"":     "Helm is a kubernetes package manager.",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (Helm) SwaggerDoc() map[string]string {
	return map_Helm
}

var map_HelmList = map[string]string{
	"":      "HelmList is the whole list of all helms which owned by a tenant.",
	"items": "List of Helms",
}

func (HelmList) SwaggerDoc() map[string]string {
	return map_HelmList
}

var map_HelmProxyOptions = map[string]string{
	"":     "HelmProxyOptions is the query options to a Helm-api proxy call.",
	"path": "Path is the URL path to use for the current proxy request to helm-api.",
}

func (HelmProxyOptions) SwaggerDoc() map[string]string {
	return map_HelmProxyOptions
}

var map_HelmSpec = map[string]string{
	"": "HelmSpec describes the attributes on a Helm.",
}

func (HelmSpec) SwaggerDoc() map[string]string {
	return map_HelmSpec
}

var map_HelmStatus = map[string]string{
	"":                            "HelmStatus is information about the current status of a Helm.",
	"phase":                       "Phase is the current lifecycle phase of the helm of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (HelmStatus) SwaggerDoc() map[string]string {
	return map_HelmStatus
}

var map_IPAM = map[string]string{
	"":     "IPAM is a scheduler plugin for assigning IP.",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (IPAM) SwaggerDoc() map[string]string {
	return map_IPAM
}

var map_IPAMList = map[string]string{
	"":      "IPAMList is the whole list of all IPAMs which owned by a tenant.",
	"items": "List of IPAMs",
}

func (IPAMList) SwaggerDoc() map[string]string {
	return map_IPAMList
}

var map_IPAMProxyOptions = map[string]string{
	"":     "IPAMProxyOptions is the query options to a ipam-api proxy call.",
	"path": "Path is the URL path to use for the current proxy request to ipam-api.",
}

func (IPAMProxyOptions) SwaggerDoc() map[string]string {
	return map_IPAMProxyOptions
}

var map_IPAMSpec = map[string]string{
	"": "IPAMSpec describes the attributes on a IPAM.",
}

func (IPAMSpec) SwaggerDoc() map[string]string {
	return map_IPAMSpec
}

var map_IPAMStatus = map[string]string{
	"":                            "IPAMStatus is information about the current status of a IPAM.",
	"phase":                       "Phase is the current lifecycle phase of the addon of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (IPAMStatus) SwaggerDoc() map[string]string {
	return map_IPAMStatus
}

var map_LBCF = map[string]string{
	"":     "LBCF is a kubernetes load balancer manager.",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (LBCF) SwaggerDoc() map[string]string {
	return map_LBCF
}

var map_LBCFList = map[string]string{
	"":      "LBCFList is the whole list of all helms which owned by a tenant.",
	"items": "List of LBCFs",
}

func (LBCFList) SwaggerDoc() map[string]string {
	return map_LBCFList
}

var map_LBCFProxyOptions = map[string]string{
	"": "LBCFProxyOptions is the query options to a kube-apiserver proxy call.",
}

func (LBCFProxyOptions) SwaggerDoc() map[string]string {
	return map_LBCFProxyOptions
}

var map_LBCFSpec = map[string]string{
	"": "LBCFSpec describes the attributes on a Helm.",
}

func (LBCFSpec) SwaggerDoc() map[string]string {
	return map_LBCFSpec
}

var map_LBCFStatus = map[string]string{
	"":                            "LBCFStatus is information about the current status of a Helm.",
	"phase":                       "Phase is the current lifecycle phase of the helm of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (LBCFStatus) SwaggerDoc() map[string]string {
	return map_LBCFStatus
}

var map_LogCollector = map[string]string{
	"":     "LogCollector is a manager to collect logs of workload.",
	"spec": "Spec defines the desired identities of LogCollector.",
}

func (LogCollector) SwaggerDoc() map[string]string {
	return map_LogCollector
}

var map_LogCollectorList = map[string]string{
	"":      "LogCollectorList is the whole list of all LogCollector which owned by a tenant.",
	"items": "List of volume decorators.",
}

func (LogCollectorList) SwaggerDoc() map[string]string {
	return map_LogCollectorList
}

var map_LogCollectorProxyOptions = map[string]string{
	"": "LogCollectorProxyOptions is the query options to a kube-apiserver proxy call for LogCollector crd object.",
}

func (LogCollectorProxyOptions) SwaggerDoc() map[string]string {
	return map_LogCollectorProxyOptions
}

var map_LogCollectorSpec = map[string]string{
	"": "LogCollectorSpec describes the attributes of a LogCollector.",
}

func (LogCollectorSpec) SwaggerDoc() map[string]string {
	return map_LogCollectorSpec
}

var map_LogCollectorStatus = map[string]string{
	"":                            "LogCollectorStatus is information about the current status of a LogCollector.",
	"phase":                       "Phase is the current lifecycle phase of the LogCollector of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (LogCollectorStatus) SwaggerDoc() map[string]string {
	return map_LogCollectorStatus
}

var map_Machine = map[string]string{
	"":     "Machine instance in Kubernetes cluster",
	"spec": "Spec defines the desired identities of the Machine.",
}

func (Machine) SwaggerDoc() map[string]string {
	return map_Machine
}

var map_MachineAddress = map[string]string{
	"":        "MachineAddress contains information for the machine's address.",
	"type":    "Machine address type, one of Public, ExternalIP or InternalIP.",
	"address": "The machine address.",
}

func (MachineAddress) SwaggerDoc() map[string]string {
	return map_MachineAddress
}

var map_MachineCondition = map[string]string{
	"":                   "MachineCondition contains details for the current condition of this Machine.",
	"type":               "Type is the type of the condition.",
	"status":             "Status is the status of the condition. Can be True, False, Unknown.",
	"lastProbeTime":      "Last time we probed the condition.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "Unique, one-word, CamelCase reason for the condition's last transition.",
	"message":            "Human-readable message indicating details about last transition.",
}

func (MachineCondition) SwaggerDoc() map[string]string {
	return map_MachineCondition
}

var map_MachineList = map[string]string{
	"":      "MachineList is the whole list of all machine in an cluster.",
	"items": "List of clusters",
}

func (MachineList) SwaggerDoc() map[string]string {
	return map_MachineList
}

var map_MachineSpec = map[string]string{
	"":           "MachineSpec is a description of machine.",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
}

func (MachineSpec) SwaggerDoc() map[string]string {
	return map_MachineSpec
}

var map_MachineStatus = map[string]string{
	"":            "MachineStatus represents information about the status of an machine.",
	"message":     "A human readable message indicating details about why the machine is in this condition.",
	"reason":      "A brief CamelCase message indicating details about why the machine is in this state.",
	"addresses":   "List of addresses reachable to the machine.",
	"machineInfo": "Set of ids/uuids to uniquely identify the node.",
}

func (MachineStatus) SwaggerDoc() map[string]string {
	return map_MachineStatus
}

var map_MachineSystemInfo = map[string]string{
	"":                        "MachineSystemInfo is a set of ids/uuids to uniquely identify the node.",
	"machineID":               "MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html",
	"systemUUID":              "SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html",
	"bootID":                  "Boot ID reported by the node.",
	"kernelVersion":           "Kernel Version reported by the node.",
	"osImage":                 "OS Image reported by the node.",
	"containerRuntimeVersion": "ContainerRuntime Version reported by the node.",
	"kubeletVersion":          "Kubelet Version reported by the node.",
	"kubeProxyVersion":        "KubeProxy Version reported by the node.",
	"operatingSystem":         "The Operating System reported by the node",
	"architecture":            "The Architecture reported by the node",
}

func (MachineSystemInfo) SwaggerDoc() map[string]string {
	return map_MachineSystemInfo
}

var map_PVCRProxyOptions = map[string]string{
	"": "PVCRProxyOptions is the query options to a kube-apiserver proxy call for PVCR crd object.",
}

func (PVCRProxyOptions) SwaggerDoc() map[string]string {
	return map_PVCRProxyOptions
}

var map_PersistentBackEnd = map[string]string{
	"": "PersistentBackEnd indicates the backend type and attributes of the persistent log store.",
}

func (PersistentBackEnd) SwaggerDoc() map[string]string {
	return map_PersistentBackEnd
}

var map_PersistentEvent = map[string]string{
	"":     "PersistentEvent is a recorder of kubernetes event.",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (PersistentEvent) SwaggerDoc() map[string]string {
	return map_PersistentEvent
}

var map_PersistentEventList = map[string]string{
	"":      "PersistentEventList is the whole list of all clusters which owned by a tenant.",
	"items": "List of PersistentEvents",
}

func (PersistentEventList) SwaggerDoc() map[string]string {
	return map_PersistentEventList
}

var map_PersistentEventSpec = map[string]string{
	"": "PersistentEventSpec describes the attributes on a PersistentEvent.",
}

func (PersistentEventSpec) SwaggerDoc() map[string]string {
	return map_PersistentEventSpec
}

var map_PersistentEventStatus = map[string]string{
	"":                            "PersistentEventStatus is information about the current status of a PersistentEvent.",
	"phase":                       "Phase is the current lifecycle phase of the persistent event of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (PersistentEventStatus) SwaggerDoc() map[string]string {
	return map_PersistentEventStatus
}

var map_Prometheus = map[string]string{
	"":     "Prometheus is a kubernetes package manager.",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (Prometheus) SwaggerDoc() map[string]string {
	return map_Prometheus
}

var map_PrometheusList = map[string]string{
	"":      "PrometheusList is the whole list of all prometheus which owned by a tenant.",
	"items": "List of Prometheuss",
}

func (PrometheusList) SwaggerDoc() map[string]string {
	return map_PrometheusList
}

var map_PrometheusRemoteAddr = map[string]string{
	"": "PrometheusRemoteAddr is the remote write/read address for prometheus",
}

func (PrometheusRemoteAddr) SwaggerDoc() map[string]string {
	return map_PrometheusRemoteAddr
}

var map_PrometheusSpec = map[string]string{
	"":              "PrometheusSpec describes the attributes on a Prometheus.",
	"subVersion":    "SubVersion is the components version such as node-exporter.",
	"remoteAddress": "RemoteAddress is the remote address for prometheus when writing/reading outside of cluster.",
}

func (PrometheusSpec) SwaggerDoc() map[string]string {
	return map_PrometheusSpec
}

var map_PrometheusStatus = map[string]string{
	"":                            "PrometheusStatus is information about the current status of a Prometheus.",
	"phase":                       "Phase is the current lifecycle phase of the helm of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
	"subVersion":                  "SubVersion is the components version such as node-exporter.",
}

func (PrometheusStatus) SwaggerDoc() map[string]string {
	return map_PrometheusStatus
}

var map_Registry = map[string]string{
	"": "Registry records the third-party image repository information stored by the user.",
}

func (Registry) SwaggerDoc() map[string]string {
	return map_Registry
}

var map_RegistryList = map[string]string{
	"": "RegistryList is a resource containing a list of Registry objects.",
}

func (RegistryList) SwaggerDoc() map[string]string {
	return map_RegistryList
}

var map_RegistrySpec = map[string]string{
	"": "RegistrySpec indicates the specifications of the third-party image repository.",
}

func (RegistrySpec) SwaggerDoc() map[string]string {
	return map_RegistrySpec
}

var map_StorageBackEndCLS = map[string]string{
	"": "StorageBackEndCLS records the attributes required when the backend storage type is CLS.",
}

func (StorageBackEndCLS) SwaggerDoc() map[string]string {
	return map_StorageBackEndCLS
}

var map_StorageBackEndES = map[string]string{
	"": "StorageBackEndES records the attributes required when the backend storage type is ElasticSearch.",
}

func (StorageBackEndES) SwaggerDoc() map[string]string {
	return map_StorageBackEndES
}

var map_TappController = map[string]string{
	"":     "TappController is a new kubernetes workload.",
	"spec": "Spec defines the desired identities of tapp controller.",
}

func (TappController) SwaggerDoc() map[string]string {
	return map_TappController
}

var map_TappControllerList = map[string]string{
	"":      "TappControllerList is the whole list of all tapp controllers which owned by a tenant.",
	"items": "List of tapp controllers",
}

func (TappControllerList) SwaggerDoc() map[string]string {
	return map_TappControllerList
}

var map_TappControllerProxyOptions = map[string]string{
	"": "TappControllerProxyOptions is the query options to a kube-apiserver proxy call.",
}

func (TappControllerProxyOptions) SwaggerDoc() map[string]string {
	return map_TappControllerProxyOptions
}

var map_TappControllerSpec = map[string]string{
	"": "TappControllerSpec describes the attributes on a tapp controller.",
}

func (TappControllerSpec) SwaggerDoc() map[string]string {
	return map_TappControllerSpec
}

var map_TappControllerStatus = map[string]string{
	"":                            "TappControllerStatus is information about the current status of a tapp controller.",
	"phase":                       "Phase is the current lifecycle phase of the tapp controller of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (TappControllerStatus) SwaggerDoc() map[string]string {
	return map_TappControllerStatus
}

var map_VolumeDecorator = map[string]string{
	"":     "VolumeDecorator is a controller to manage PVC information.",
	"spec": "Spec defines the desired identities of volume decorator.",
}

func (VolumeDecorator) SwaggerDoc() map[string]string {
	return map_VolumeDecorator
}

var map_VolumeDecoratorList = map[string]string{
	"":      "VolumeDecoratorList is the whole list of all VolumeDecorator which owned by a tenant.",
	"items": "List of volume decorators.",
}

func (VolumeDecoratorList) SwaggerDoc() map[string]string {
	return map_VolumeDecoratorList
}

var map_VolumeDecoratorSpec = map[string]string{
	"": "VolumeDecoratorSpec describes the attributes of a VolumeDecorator.",
}

func (VolumeDecoratorSpec) SwaggerDoc() map[string]string {
	return map_VolumeDecoratorSpec
}

var map_VolumeDecoratorStatus = map[string]string{
	"":                            "VolumeDecoratorStatus is information about the current status of a VolumeDecorator.",
	"volumeTypes":                 "VolumeTypes is the supported volume types in this cluster.",
	"workloadAdmission":           "WorkloadAdmission will be true to enable the workload admission webhook.",
	"storageVendorVersion":        "StorageVendorVersion will be set to the config version of the storage vendor.",
	"phase":                       "Phase is the current lifecycle phase of the volume decorator of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (VolumeDecoratorStatus) SwaggerDoc() map[string]string {
	return map_VolumeDecoratorStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
