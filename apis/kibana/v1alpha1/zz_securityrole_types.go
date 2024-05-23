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

type ElasticsearchInitParameters struct {

	// (Set of String) List of the cluster privileges.
	// List of the cluster privileges.
	// +listType=set
	Cluster []*string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (Block Set) A list of indices permissions entries. (see below for nested schema)
	// A list of indices permissions entries.
	Indices []IndicesInitParameters `json:"indices,omitempty" tf:"indices,omitempty"`

	// (Set of String) A list of usernames the owners of this role can impersonate.
	// A list of usernames the owners of this role can impersonate.
	// +listType=set
	RunAs []*string `json:"runAs,omitempty" tf:"run_as,omitempty"`
}

type ElasticsearchObservation struct {

	// (Set of String) List of the cluster privileges.
	// List of the cluster privileges.
	// +listType=set
	Cluster []*string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (Block Set) A list of indices permissions entries. (see below for nested schema)
	// A list of indices permissions entries.
	Indices []IndicesObservation `json:"indices,omitempty" tf:"indices,omitempty"`

	// (Set of String) A list of usernames the owners of this role can impersonate.
	// A list of usernames the owners of this role can impersonate.
	// +listType=set
	RunAs []*string `json:"runAs,omitempty" tf:"run_as,omitempty"`
}

type ElasticsearchParameters struct {

	// (Set of String) List of the cluster privileges.
	// List of the cluster privileges.
	// +kubebuilder:validation:Optional
	// +listType=set
	Cluster []*string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (Block Set) A list of indices permissions entries. (see below for nested schema)
	// A list of indices permissions entries.
	// +kubebuilder:validation:Optional
	Indices []IndicesParameters `json:"indices,omitempty" tf:"indices,omitempty"`

	// (Set of String) A list of usernames the owners of this role can impersonate.
	// A list of usernames the owners of this role can impersonate.
	// +kubebuilder:validation:Optional
	// +listType=set
	RunAs []*string `json:"runAs,omitempty" tf:"run_as,omitempty"`
}

type FeatureInitParameters struct {

	// (String) The name for the role.
	// Feature name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) The index level privileges that the owners of the role have on the specified indices.
	// Feature privileges.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`
}

type FeatureObservation struct {

	// (String) The name for the role.
	// Feature name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) The index level privileges that the owners of the role have on the specified indices.
	// Feature privileges.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`
}

type FeatureParameters struct {

	// (String) The name for the role.
	// Feature name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (Set of String) The index level privileges that the owners of the role have on the specified indices.
	// Feature privileges.
	// +kubebuilder:validation:Optional
	// +listType=set
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`
}

type FieldSecurityInitParameters struct {

	// (Set of String) List of the fields to which the grants will not be applied.
	// List of the fields to which the grants will not be applied.
	// +listType=set
	Except []*string `json:"except,omitempty" tf:"except,omitempty"`

	// (Set of String) List of the fields to grant the access to.
	// List of the fields to grant the access to.
	// +listType=set
	Grant []*string `json:"grant,omitempty" tf:"grant,omitempty"`
}

type FieldSecurityObservation struct {

	// (Set of String) List of the fields to which the grants will not be applied.
	// List of the fields to which the grants will not be applied.
	// +listType=set
	Except []*string `json:"except,omitempty" tf:"except,omitempty"`

	// (Set of String) List of the fields to grant the access to.
	// List of the fields to grant the access to.
	// +listType=set
	Grant []*string `json:"grant,omitempty" tf:"grant,omitempty"`
}

type FieldSecurityParameters struct {

	// (Set of String) List of the fields to which the grants will not be applied.
	// List of the fields to which the grants will not be applied.
	// +kubebuilder:validation:Optional
	// +listType=set
	Except []*string `json:"except,omitempty" tf:"except,omitempty"`

	// (Set of String) List of the fields to grant the access to.
	// List of the fields to grant the access to.
	// +kubebuilder:validation:Optional
	// +listType=set
	Grant []*string `json:"grant,omitempty" tf:"grant,omitempty"`
}

type IndicesInitParameters struct {

	// (Block List, Max: 1) The document fields that the owners of the role have read access to. (see below for nested schema)
	// The document fields that the owners of the role have read access to.
	FieldSecurity []FieldSecurityInitParameters `json:"fieldSecurity,omitempty" tf:"field_security,omitempty"`

	// (Set of String) A list of indices (or index name patterns) to which the permissions in this entry apply.
	// A list of indices (or index name patterns) to which the permissions in this entry apply.
	// +listType=set
	Names []*string `json:"names,omitempty" tf:"names,omitempty"`

	// (Set of String) The index level privileges that the owners of the role have on the specified indices.
	// The index level privileges that the owners of the role have on the specified indices.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (String) A search query that defines the documents the owners of the role have read access to.
	// A search query that defines the documents the owners of the role have read access to.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type IndicesObservation struct {

	// (Block List, Max: 1) The document fields that the owners of the role have read access to. (see below for nested schema)
	// The document fields that the owners of the role have read access to.
	FieldSecurity []FieldSecurityObservation `json:"fieldSecurity,omitempty" tf:"field_security,omitempty"`

	// (Set of String) A list of indices (or index name patterns) to which the permissions in this entry apply.
	// A list of indices (or index name patterns) to which the permissions in this entry apply.
	// +listType=set
	Names []*string `json:"names,omitempty" tf:"names,omitempty"`

	// (Set of String) The index level privileges that the owners of the role have on the specified indices.
	// The index level privileges that the owners of the role have on the specified indices.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (String) A search query that defines the documents the owners of the role have read access to.
	// A search query that defines the documents the owners of the role have read access to.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type IndicesParameters struct {

	// (Block List, Max: 1) The document fields that the owners of the role have read access to. (see below for nested schema)
	// The document fields that the owners of the role have read access to.
	// +kubebuilder:validation:Optional
	FieldSecurity []FieldSecurityParameters `json:"fieldSecurity,omitempty" tf:"field_security,omitempty"`

	// (Set of String) A list of indices (or index name patterns) to which the permissions in this entry apply.
	// A list of indices (or index name patterns) to which the permissions in this entry apply.
	// +kubebuilder:validation:Optional
	// +listType=set
	Names []*string `json:"names" tf:"names,omitempty"`

	// (Set of String) The index level privileges that the owners of the role have on the specified indices.
	// The index level privileges that the owners of the role have on the specified indices.
	// +kubebuilder:validation:Optional
	// +listType=set
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`

	// (String) A search query that defines the documents the owners of the role have read access to.
	// A search query that defines the documents the owners of the role have read access to.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type KibanaInitParameters struct {

	// (Set of String) A base privilege. When specified, the base must be ["all"] or ["read"]. When the base privileges are specified, you are unable to use the "feature" section.
	// A base privilege. When specified, the base must be ["all"] or ["read"]. When the base privileges are specified, you are unable to use the "feature" section.
	// +listType=set
	Base []*string `json:"base,omitempty" tf:"base,omitempty"`

	// (Block Set) List of privileges for specific features. When the feature privileges are specified, you are unable to use the "base" section. (see below for nested schema)
	// List of privileges for specific features. When the feature privileges are specified, you are unable to use the "base" section.
	Feature []FeatureInitParameters `json:"feature,omitempty" tf:"feature,omitempty"`

	// (Set of String) The spaces to apply the privileges to. To grant access to all spaces, set to ["*"], or omit the value.
	// The spaces to apply the privileges to. To grant access to all spaces, set to ["*"], or omit the value.
	// +listType=set
	Spaces []*string `json:"spaces,omitempty" tf:"spaces,omitempty"`
}

type KibanaObservation struct {

	// (Set of String) A base privilege. When specified, the base must be ["all"] or ["read"]. When the base privileges are specified, you are unable to use the "feature" section.
	// A base privilege. When specified, the base must be ["all"] or ["read"]. When the base privileges are specified, you are unable to use the "feature" section.
	// +listType=set
	Base []*string `json:"base,omitempty" tf:"base,omitempty"`

	// (Block Set) List of privileges for specific features. When the feature privileges are specified, you are unable to use the "base" section. (see below for nested schema)
	// List of privileges for specific features. When the feature privileges are specified, you are unable to use the "base" section.
	Feature []FeatureObservation `json:"feature,omitempty" tf:"feature,omitempty"`

	// (Set of String) The spaces to apply the privileges to. To grant access to all spaces, set to ["*"], or omit the value.
	// The spaces to apply the privileges to. To grant access to all spaces, set to ["*"], or omit the value.
	// +listType=set
	Spaces []*string `json:"spaces,omitempty" tf:"spaces,omitempty"`
}

type KibanaParameters struct {

	// (Set of String) A base privilege. When specified, the base must be ["all"] or ["read"]. When the base privileges are specified, you are unable to use the "feature" section.
	// A base privilege. When specified, the base must be ["all"] or ["read"]. When the base privileges are specified, you are unable to use the "feature" section.
	// +kubebuilder:validation:Optional
	// +listType=set
	Base []*string `json:"base,omitempty" tf:"base,omitempty"`

	// (Block Set) List of privileges for specific features. When the feature privileges are specified, you are unable to use the "base" section. (see below for nested schema)
	// List of privileges for specific features. When the feature privileges are specified, you are unable to use the "base" section.
	// +kubebuilder:validation:Optional
	Feature []FeatureParameters `json:"feature,omitempty" tf:"feature,omitempty"`

	// (Set of String) The spaces to apply the privileges to. To grant access to all spaces, set to ["*"], or omit the value.
	// The spaces to apply the privileges to. To grant access to all spaces, set to ["*"], or omit the value.
	// +kubebuilder:validation:Optional
	// +listType=set
	Spaces []*string `json:"spaces" tf:"spaces,omitempty"`
}

type SecurityRoleInitParameters struct {

	// (Block Set, Min: 1, Max: 1) Elasticsearch cluster and index privileges. (see below for nested schema)
	// Elasticsearch cluster and index privileges.
	Elasticsearch []ElasticsearchInitParameters `json:"elasticsearch,omitempty" tf:"elasticsearch,omitempty"`

	// (Block Set) The list of objects that specify the Kibana privileges for the role. (see below for nested schema)
	// The list of objects that specify the Kibana privileges for the role.
	Kibana []KibanaInitParameters `json:"kibana,omitempty" tf:"kibana,omitempty"`

	// data.
	// Optional meta-data.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) The name for the role.
	// The name for the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecurityRoleObservation struct {

	// (Block Set, Min: 1, Max: 1) Elasticsearch cluster and index privileges. (see below for nested schema)
	// Elasticsearch cluster and index privileges.
	Elasticsearch []ElasticsearchObservation `json:"elasticsearch,omitempty" tf:"elasticsearch,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block Set) The list of objects that specify the Kibana privileges for the role. (see below for nested schema)
	// The list of objects that specify the Kibana privileges for the role.
	Kibana []KibanaObservation `json:"kibana,omitempty" tf:"kibana,omitempty"`

	// data.
	// Optional meta-data.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) The name for the role.
	// The name for the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecurityRoleParameters struct {

	// (Block Set, Min: 1, Max: 1) Elasticsearch cluster and index privileges. (see below for nested schema)
	// Elasticsearch cluster and index privileges.
	// +kubebuilder:validation:Optional
	Elasticsearch []ElasticsearchParameters `json:"elasticsearch,omitempty" tf:"elasticsearch,omitempty"`

	// (Block Set) The list of objects that specify the Kibana privileges for the role. (see below for nested schema)
	// The list of objects that specify the Kibana privileges for the role.
	// +kubebuilder:validation:Optional
	Kibana []KibanaParameters `json:"kibana,omitempty" tf:"kibana,omitempty"`

	// data.
	// Optional meta-data.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) The name for the role.
	// The name for the role.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// SecurityRoleSpec defines the desired state of SecurityRole
type SecurityRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityRoleParameters `json:"forProvider"`
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
	InitProvider SecurityRoleInitParameters `json:"initProvider,omitempty"`
}

// SecurityRoleStatus defines the observed state of SecurityRole.
type SecurityRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityRole is the Schema for the SecurityRoles API. Creates or updates a Kibana role.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type SecurityRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.elasticsearch) || (has(self.initProvider) && has(self.initProvider.elasticsearch))",message="spec.forProvider.elasticsearch is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SecurityRoleSpec   `json:"spec"`
	Status SecurityRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityRoleList contains a list of SecurityRoles
type SecurityRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityRole `json:"items"`
}

// Repository type metadata.
var (
	SecurityRole_Kind             = "SecurityRole"
	SecurityRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityRole_Kind}.String()
	SecurityRole_KindAPIVersion   = SecurityRole_Kind + "." + CRDGroupVersion.String()
	SecurityRole_GroupVersionKind = CRDGroupVersion.WithKind(SecurityRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityRole{}, &SecurityRoleList{})
}
