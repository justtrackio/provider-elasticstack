// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AgentPolicyInitParameters struct {

	// (String) The identifier for the data output.
	// The identifier for the data output.
	DataOutputID *string `json:"dataOutputId,omitempty" tf:"data_output_id,omitempty"`

	// (String) The description of the agent policy.
	// The description of the agent policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The identifier for the Elastic Agent binary download server.
	// The identifier for the Elastic Agent binary download server.
	DownloadSourceID *string `json:"downloadSourceId,omitempty" tf:"download_source_id,omitempty"`

	// (String) The identifier for the Fleet server host.
	// The identifier for the Fleet server host.
	FleetServerHostID *string `json:"fleetServerHostId,omitempty" tf:"fleet_server_host_id,omitempty"`

	// (Boolean) Enable collection of agent logs.
	// Enable collection of agent logs.
	MonitorLogs *bool `json:"monitorLogs,omitempty" tf:"monitor_logs,omitempty"`

	// (Boolean) Enable collection of agent metrics.
	// Enable collection of agent metrics.
	MonitorMetrics *bool `json:"monitorMetrics,omitempty" tf:"monitor_metrics,omitempty"`

	// (String) The identifier for monitoring output.
	// The identifier for monitoring output.
	MonitoringOutputID *string `json:"monitoringOutputId,omitempty" tf:"monitoring_output_id,omitempty"`

	// (String) The name of the agent policy.
	// The name of the agent policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The namespace of the agent policy.
	// The namespace of the agent policy.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Unique identifier of the agent policy.
	// Unique identifier of the agent policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// (Boolean) Enable collection of system logs and metrics.
	// Enable collection of system logs and metrics.
	SysMonitoring *bool `json:"sysMonitoring,omitempty" tf:"sys_monitoring,omitempty"`
}

type AgentPolicyObservation struct {

	// (String) The identifier for the data output.
	// The identifier for the data output.
	DataOutputID *string `json:"dataOutputId,omitempty" tf:"data_output_id,omitempty"`

	// (String) The description of the agent policy.
	// The description of the agent policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The identifier for the Elastic Agent binary download server.
	// The identifier for the Elastic Agent binary download server.
	DownloadSourceID *string `json:"downloadSourceId,omitempty" tf:"download_source_id,omitempty"`

	// (String) The identifier for the Fleet server host.
	// The identifier for the Fleet server host.
	FleetServerHostID *string `json:"fleetServerHostId,omitempty" tf:"fleet_server_host_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Enable collection of agent logs.
	// Enable collection of agent logs.
	MonitorLogs *bool `json:"monitorLogs,omitempty" tf:"monitor_logs,omitempty"`

	// (Boolean) Enable collection of agent metrics.
	// Enable collection of agent metrics.
	MonitorMetrics *bool `json:"monitorMetrics,omitempty" tf:"monitor_metrics,omitempty"`

	// (String) The identifier for monitoring output.
	// The identifier for monitoring output.
	MonitoringOutputID *string `json:"monitoringOutputId,omitempty" tf:"monitoring_output_id,omitempty"`

	// (String) The name of the agent policy.
	// The name of the agent policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The namespace of the agent policy.
	// The namespace of the agent policy.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Unique identifier of the agent policy.
	// Unique identifier of the agent policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// (Boolean) Enable collection of system logs and metrics.
	// Enable collection of system logs and metrics.
	SysMonitoring *bool `json:"sysMonitoring,omitempty" tf:"sys_monitoring,omitempty"`
}

type AgentPolicyParameters struct {

	// (String) The identifier for the data output.
	// The identifier for the data output.
	// +kubebuilder:validation:Optional
	DataOutputID *string `json:"dataOutputId,omitempty" tf:"data_output_id,omitempty"`

	// (String) The description of the agent policy.
	// The description of the agent policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The identifier for the Elastic Agent binary download server.
	// The identifier for the Elastic Agent binary download server.
	// +kubebuilder:validation:Optional
	DownloadSourceID *string `json:"downloadSourceId,omitempty" tf:"download_source_id,omitempty"`

	// (String) The identifier for the Fleet server host.
	// The identifier for the Fleet server host.
	// +kubebuilder:validation:Optional
	FleetServerHostID *string `json:"fleetServerHostId,omitempty" tf:"fleet_server_host_id,omitempty"`

	// (Boolean) Enable collection of agent logs.
	// Enable collection of agent logs.
	// +kubebuilder:validation:Optional
	MonitorLogs *bool `json:"monitorLogs,omitempty" tf:"monitor_logs,omitempty"`

	// (Boolean) Enable collection of agent metrics.
	// Enable collection of agent metrics.
	// +kubebuilder:validation:Optional
	MonitorMetrics *bool `json:"monitorMetrics,omitempty" tf:"monitor_metrics,omitempty"`

	// (String) The identifier for monitoring output.
	// The identifier for monitoring output.
	// +kubebuilder:validation:Optional
	MonitoringOutputID *string `json:"monitoringOutputId,omitempty" tf:"monitoring_output_id,omitempty"`

	// (String) The name of the agent policy.
	// The name of the agent policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The namespace of the agent policy.
	// The namespace of the agent policy.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Unique identifier of the agent policy.
	// Unique identifier of the agent policy.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// (Boolean) Enable collection of system logs and metrics.
	// Enable collection of system logs and metrics.
	// +kubebuilder:validation:Optional
	SysMonitoring *bool `json:"sysMonitoring,omitempty" tf:"sys_monitoring,omitempty"`
}

// AgentPolicySpec defines the desired state of AgentPolicy
type AgentPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AgentPolicyInitParameters `json:"initProvider,omitempty"`
}

// AgentPolicyStatus defines the observed state of AgentPolicy.
type AgentPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AgentPolicy is the Schema for the AgentPolicys API. Creates or updates a Fleet Agent Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type AgentPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.__namespace__) || (has(self.initProvider) && has(self.initProvider.__namespace__))",message="spec.forProvider.namespace is a required parameter"
	Spec   AgentPolicySpec   `json:"spec"`
	Status AgentPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentPolicyList contains a list of AgentPolicys
type AgentPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentPolicy `json:"items"`
}

// Repository type metadata.
var (
	AgentPolicy_Kind             = "AgentPolicy"
	AgentPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AgentPolicy_Kind}.String()
	AgentPolicy_KindAPIVersion   = AgentPolicy_Kind + "." + CRDGroupVersion.String()
	AgentPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AgentPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AgentPolicy{}, &AgentPolicyList{})
}
