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

type InputInitParameters struct {

	// (Boolean) Enable the integration policy.
	// Enable the input.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The identifier of the input.
	// The identifier of the input.
	InputID *string `json:"inputId,omitempty" tf:"input_id,omitempty"`
}

type InputObservation struct {

	// (Boolean) Enable the integration policy.
	// Enable the input.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The identifier of the input.
	// The identifier of the input.
	InputID *string `json:"inputId,omitempty" tf:"input_id,omitempty"`
}

type InputParameters struct {

	// (Boolean) Enable the integration policy.
	// Enable the input.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The identifier of the input.
	// The identifier of the input.
	// +kubebuilder:validation:Optional
	InputID *string `json:"inputId" tf:"input_id,omitempty"`

	// (String, Sensitive) Input streams as JSON.
	// Input streams as JSON.
	// +kubebuilder:validation:Optional
	StreamsJSONSecretRef *v1.SecretKeySelector `json:"streamsJsonSecretRef,omitempty" tf:"-"`

	// level variables as JSON.
	// Input variables as JSON.
	// +kubebuilder:validation:Optional
	VarsJSONSecretRef *v1.SecretKeySelector `json:"varsJsonSecretRef,omitempty" tf:"-"`
}

type IntegrationPolicyInitParameters struct {

	// (String) ID of the agent policy.
	// ID of the agent policy.
	AgentPolicyID *string `json:"agentPolicyId,omitempty" tf:"agent_policy_id,omitempty"`

	// (String) The description of the integration policy.
	// The description of the integration policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Enable the integration policy.
	// Enable the integration policy.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Boolean) Force operations, such as creation and deletion, to occur.
	// Force operations, such as creation and deletion, to occur.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// (Block List) (see below for nested schema)
	Input []InputInitParameters `json:"input,omitempty" tf:"input,omitempty"`

	// (String) The name of the integration package.
	// The name of the integration package.
	IntegrationName *string `json:"integrationName,omitempty" tf:"integration_name,omitempty"`

	// (String) The version of the integration package.
	// The version of the integration package.
	IntegrationVersion *string `json:"integrationVersion,omitempty" tf:"integration_version,omitempty"`

	// (String) The name of the integration policy.
	// The name of the integration policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The namespace of the integration policy.
	// The namespace of the integration policy.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Unique identifier of the integration policy.
	// Unique identifier of the integration policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

type IntegrationPolicyObservation struct {

	// (String) ID of the agent policy.
	// ID of the agent policy.
	AgentPolicyID *string `json:"agentPolicyId,omitempty" tf:"agent_policy_id,omitempty"`

	// (String) The description of the integration policy.
	// The description of the integration policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Enable the integration policy.
	// Enable the integration policy.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Boolean) Force operations, such as creation and deletion, to occur.
	// Force operations, such as creation and deletion, to occur.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List) (see below for nested schema)
	Input []InputObservation `json:"input,omitempty" tf:"input,omitempty"`

	// (String) The name of the integration package.
	// The name of the integration package.
	IntegrationName *string `json:"integrationName,omitempty" tf:"integration_name,omitempty"`

	// (String) The version of the integration package.
	// The version of the integration package.
	IntegrationVersion *string `json:"integrationVersion,omitempty" tf:"integration_version,omitempty"`

	// (String) The name of the integration policy.
	// The name of the integration policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The namespace of the integration policy.
	// The namespace of the integration policy.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Unique identifier of the integration policy.
	// Unique identifier of the integration policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

type IntegrationPolicyParameters struct {

	// (String) ID of the agent policy.
	// ID of the agent policy.
	// +kubebuilder:validation:Optional
	AgentPolicyID *string `json:"agentPolicyId,omitempty" tf:"agent_policy_id,omitempty"`

	// (String) The description of the integration policy.
	// The description of the integration policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Enable the integration policy.
	// Enable the integration policy.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Boolean) Force operations, such as creation and deletion, to occur.
	// Force operations, such as creation and deletion, to occur.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Input []InputParameters `json:"input,omitempty" tf:"input,omitempty"`

	// (String) The name of the integration package.
	// The name of the integration package.
	// +kubebuilder:validation:Optional
	IntegrationName *string `json:"integrationName,omitempty" tf:"integration_name,omitempty"`

	// (String) The version of the integration package.
	// The version of the integration package.
	// +kubebuilder:validation:Optional
	IntegrationVersion *string `json:"integrationVersion,omitempty" tf:"integration_version,omitempty"`

	// (String) The name of the integration policy.
	// The name of the integration policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The namespace of the integration policy.
	// The namespace of the integration policy.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Unique identifier of the integration policy.
	// Unique identifier of the integration policy.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// level variables as JSON.
	// Integration-level variables as JSON.
	// +kubebuilder:validation:Optional
	VarsJSONSecretRef *v1.SecretKeySelector `json:"varsJsonSecretRef,omitempty" tf:"-"`
}

// IntegrationPolicySpec defines the desired state of IntegrationPolicy
type IntegrationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationPolicyParameters `json:"forProvider"`
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
	InitProvider IntegrationPolicyInitParameters `json:"initProvider,omitempty"`
}

// IntegrationPolicyStatus defines the observed state of IntegrationPolicy.
type IntegrationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IntegrationPolicy is the Schema for the IntegrationPolicys API. Creates or updates a Fleet Integration Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type IntegrationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.agentPolicyId) || (has(self.initProvider) && has(self.initProvider.agentPolicyId))",message="spec.forProvider.agentPolicyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationName) || (has(self.initProvider) && has(self.initProvider.integrationName))",message="spec.forProvider.integrationName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationVersion) || (has(self.initProvider) && has(self.initProvider.integrationVersion))",message="spec.forProvider.integrationVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.__namespace__) || (has(self.initProvider) && has(self.initProvider.__namespace__))",message="spec.forProvider.namespace is a required parameter"
	Spec   IntegrationPolicySpec   `json:"spec"`
	Status IntegrationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationPolicyList contains a list of IntegrationPolicys
type IntegrationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationPolicy `json:"items"`
}

// Repository type metadata.
var (
	IntegrationPolicy_Kind             = "IntegrationPolicy"
	IntegrationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationPolicy_Kind}.String()
	IntegrationPolicy_KindAPIVersion   = IntegrationPolicy_Kind + "." + CRDGroupVersion.String()
	IntegrationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationPolicy{}, &IntegrationPolicyList{})
}
