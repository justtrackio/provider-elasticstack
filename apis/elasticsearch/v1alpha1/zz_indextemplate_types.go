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

type IndexTemplateDataStreamInitParameters struct {

	// (Boolean) If true, the data stream supports custom routing. Defaults to false. Available only in 8.x
	// If `true`, the data stream supports custom routing. Defaults to `false`. Available only in **8.x**
	AllowCustomRouting *bool `json:"allowCustomRouting,omitempty" tf:"allow_custom_routing,omitempty"`

	// (Boolean) If true, the data stream is hidden.
	// If true, the data stream is hidden.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`
}

type IndexTemplateDataStreamObservation struct {

	// (Boolean) If true, the data stream supports custom routing. Defaults to false. Available only in 8.x
	// If `true`, the data stream supports custom routing. Defaults to `false`. Available only in **8.x**
	AllowCustomRouting *bool `json:"allowCustomRouting,omitempty" tf:"allow_custom_routing,omitempty"`

	// (Boolean) If true, the data stream is hidden.
	// If true, the data stream is hidden.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`
}

type IndexTemplateDataStreamParameters struct {

	// (Boolean) If true, the data stream supports custom routing. Defaults to false. Available only in 8.x
	// If `true`, the data stream supports custom routing. Defaults to `false`. Available only in **8.x**
	// +kubebuilder:validation:Optional
	AllowCustomRouting *bool `json:"allowCustomRouting,omitempty" tf:"allow_custom_routing,omitempty"`

	// (Boolean) If true, the data stream is hidden.
	// If true, the data stream is hidden.
	// +kubebuilder:validation:Optional
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`
}

type IndexTemplateElasticsearchConnectionInitParameters struct {

	// encoded custom Certificate Authority certificate
	// PEM-encoded custom Certificate Authority certificate
	CAData *string `json:"caData,omitempty" tf:"ca_data,omitempty"`

	// (String) Path to a custom Certificate Authority certificate
	// Path to a custom Certificate Authority certificate
	CAFile *string `json:"caFile,omitempty" tf:"ca_file,omitempty"`

	// (String) PEM encoded certificate for client auth
	// PEM encoded certificate for client auth
	CertData *string `json:"certData,omitempty" tf:"cert_data,omitempty"`

	// (String) Path to a file containing the PEM encoded certificate for client auth
	// Path to a file containing the PEM encoded certificate for client auth
	CertFile *string `json:"certFile,omitempty" tf:"cert_file,omitempty"`

	// (Boolean) Disable TLS certificate validation
	// Disable TLS certificate validation
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// (String) Path to a file containing the PEM encoded private key for client auth
	// Path to a file containing the PEM encoded private key for client auth
	KeyFile *string `json:"keyFile,omitempty" tf:"key_file,omitempty"`

	// (String) Username to use for API authentication to Elasticsearch.
	// Username to use for API authentication to Elasticsearch.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type IndexTemplateElasticsearchConnectionObservation struct {

	// encoded custom Certificate Authority certificate
	// PEM-encoded custom Certificate Authority certificate
	CAData *string `json:"caData,omitempty" tf:"ca_data,omitempty"`

	// (String) Path to a custom Certificate Authority certificate
	// Path to a custom Certificate Authority certificate
	CAFile *string `json:"caFile,omitempty" tf:"ca_file,omitempty"`

	// (String) PEM encoded certificate for client auth
	// PEM encoded certificate for client auth
	CertData *string `json:"certData,omitempty" tf:"cert_data,omitempty"`

	// (String) Path to a file containing the PEM encoded certificate for client auth
	// Path to a file containing the PEM encoded certificate for client auth
	CertFile *string `json:"certFile,omitempty" tf:"cert_file,omitempty"`

	// (Boolean) Disable TLS certificate validation
	// Disable TLS certificate validation
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// (String) Path to a file containing the PEM encoded private key for client auth
	// Path to a file containing the PEM encoded private key for client auth
	KeyFile *string `json:"keyFile,omitempty" tf:"key_file,omitempty"`

	// (String) Username to use for API authentication to Elasticsearch.
	// Username to use for API authentication to Elasticsearch.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type IndexTemplateElasticsearchConnectionParameters struct {

	// (String, Sensitive) API Key to use for authentication to Elasticsearch
	// API Key to use for authentication to Elasticsearch
	// +kubebuilder:validation:Optional
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// (String, Sensitive) Bearer Token to use for authentication to Elasticsearch
	// Bearer Token to use for authentication to Elasticsearch
	// +kubebuilder:validation:Optional
	BearerTokenSecretRef *v1.SecretKeySelector `json:"bearerTokenSecretRef,omitempty" tf:"-"`

	// encoded custom Certificate Authority certificate
	// PEM-encoded custom Certificate Authority certificate
	// +kubebuilder:validation:Optional
	CAData *string `json:"caData,omitempty" tf:"ca_data,omitempty"`

	// (String) Path to a custom Certificate Authority certificate
	// Path to a custom Certificate Authority certificate
	// +kubebuilder:validation:Optional
	CAFile *string `json:"caFile,omitempty" tf:"ca_file,omitempty"`

	// (String) PEM encoded certificate for client auth
	// PEM encoded certificate for client auth
	// +kubebuilder:validation:Optional
	CertData *string `json:"certData,omitempty" tf:"cert_data,omitempty"`

	// (String) Path to a file containing the PEM encoded certificate for client auth
	// Path to a file containing the PEM encoded certificate for client auth
	// +kubebuilder:validation:Optional
	CertFile *string `json:"certFile,omitempty" tf:"cert_file,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointsSecretRef *[]v1.SecretKeySelector `json:"endpointsSecretRef,omitempty" tf:"-"`

	// (String, Sensitive) ES Client Authentication field to be used with the bearer token
	// ES Client Authentication field to be used with the bearer token
	// +kubebuilder:validation:Optional
	EsClientAuthenticationSecretRef *v1.SecretKeySelector `json:"esClientAuthenticationSecretRef,omitempty" tf:"-"`

	// (Boolean) Disable TLS certificate validation
	// Disable TLS certificate validation
	// +kubebuilder:validation:Optional
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// (String, Sensitive) PEM encoded private key for client auth
	// PEM encoded private key for client auth
	// +kubebuilder:validation:Optional
	KeyDataSecretRef *v1.SecretKeySelector `json:"keyDataSecretRef,omitempty" tf:"-"`

	// (String) Path to a file containing the PEM encoded private key for client auth
	// Path to a file containing the PEM encoded private key for client auth
	// +kubebuilder:validation:Optional
	KeyFile *string `json:"keyFile,omitempty" tf:"key_file,omitempty"`

	// (String, Sensitive) Password to use for API authentication to Elasticsearch.
	// Password to use for API authentication to Elasticsearch.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String) Username to use for API authentication to Elasticsearch.
	// Username to use for API authentication to Elasticsearch.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type IndexTemplateInitParameters struct {

	// (List of String) An ordered list of component template names.
	// An ordered list of component template names.
	ComposedOf []*string `json:"composedOf,omitempty" tf:"composed_of,omitempty"`

	// (Block List, Max: 1) If this object is included, the template is used to create data streams and their backing indices. Supports an empty object. (see below for nested schema)
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	DataStream []IndexTemplateDataStreamInitParameters `json:"dataStream,omitempty" tf:"data_stream,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []IndexTemplateElasticsearchConnectionInitParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (Set of String) Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	IndexPatterns []*string `json:"indexPatterns,omitempty" tf:"index_patterns,omitempty"`

	// (String) Optional user metadata about the index template.
	// Optional user metadata about the index template.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the index template to create.
	// Name of the index template to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Priority to determine index template precedence when a new data stream or index is created.
	// Priority to determine index template precedence when a new data stream or index is created.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (Block List, Max: 1) Template to be applied. It may optionally include an aliases, mappings, or settings configuration. (see below for nested schema)
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template []IndexTemplateTemplateInitParameters `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Version number used to manage index templates externally.
	// Version number used to manage index templates externally.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type IndexTemplateObservation struct {

	// (List of String) An ordered list of component template names.
	// An ordered list of component template names.
	ComposedOf []*string `json:"composedOf,omitempty" tf:"composed_of,omitempty"`

	// (Block List, Max: 1) If this object is included, the template is used to create data streams and their backing indices. Supports an empty object. (see below for nested schema)
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	DataStream []IndexTemplateDataStreamObservation `json:"dataStream,omitempty" tf:"data_stream,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []IndexTemplateElasticsearchConnectionObservation `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) Internal identifier of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	IndexPatterns []*string `json:"indexPatterns,omitempty" tf:"index_patterns,omitempty"`

	// (String) Optional user metadata about the index template.
	// Optional user metadata about the index template.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the index template to create.
	// Name of the index template to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Priority to determine index template precedence when a new data stream or index is created.
	// Priority to determine index template precedence when a new data stream or index is created.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (Block List, Max: 1) Template to be applied. It may optionally include an aliases, mappings, or settings configuration. (see below for nested schema)
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template []IndexTemplateTemplateObservation `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Version number used to manage index templates externally.
	// Version number used to manage index templates externally.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type IndexTemplateParameters struct {

	// (List of String) An ordered list of component template names.
	// An ordered list of component template names.
	// +kubebuilder:validation:Optional
	ComposedOf []*string `json:"composedOf,omitempty" tf:"composed_of,omitempty"`

	// (Block List, Max: 1) If this object is included, the template is used to create data streams and their backing indices. Supports an empty object. (see below for nested schema)
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	// +kubebuilder:validation:Optional
	DataStream []IndexTemplateDataStreamParameters `json:"dataStream,omitempty" tf:"data_stream,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	// +kubebuilder:validation:Optional
	ElasticsearchConnection []IndexTemplateElasticsearchConnectionParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (Set of String) Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	// +kubebuilder:validation:Optional
	IndexPatterns []*string `json:"indexPatterns,omitempty" tf:"index_patterns,omitempty"`

	// (String) Optional user metadata about the index template.
	// Optional user metadata about the index template.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the index template to create.
	// Name of the index template to create.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Priority to determine index template precedence when a new data stream or index is created.
	// Priority to determine index template precedence when a new data stream or index is created.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (Block List, Max: 1) Template to be applied. It may optionally include an aliases, mappings, or settings configuration. (see below for nested schema)
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	// +kubebuilder:validation:Optional
	Template []IndexTemplateTemplateParameters `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Version number used to manage index templates externally.
	// Version number used to manage index templates externally.
	// +kubebuilder:validation:Optional
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type IndexTemplateTemplateInitParameters struct {

	// (Block Set) Alias to add. (see below for nested schema)
	// Alias to add.
	Alias []TemplateAliasInitParameters `json:"alias,omitempty" tf:"alias,omitempty"`

	// (String) Mapping for fields in the index.
	// Mapping for fields in the index.
	Mappings *string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// modules.html#index-modules-settings
	// Configuration options for the index. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type IndexTemplateTemplateObservation struct {

	// (Block Set) Alias to add. (see below for nested schema)
	// Alias to add.
	Alias []TemplateAliasObservation `json:"alias,omitempty" tf:"alias,omitempty"`

	// (String) Mapping for fields in the index.
	// Mapping for fields in the index.
	Mappings *string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// modules.html#index-modules-settings
	// Configuration options for the index. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type IndexTemplateTemplateParameters struct {

	// (Block Set) Alias to add. (see below for nested schema)
	// Alias to add.
	// +kubebuilder:validation:Optional
	Alias []TemplateAliasParameters `json:"alias,omitempty" tf:"alias,omitempty"`

	// (String) Mapping for fields in the index.
	// Mapping for fields in the index.
	// +kubebuilder:validation:Optional
	Mappings *string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// modules.html#index-modules-settings
	// Configuration options for the index. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
	// +kubebuilder:validation:Optional
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type TemplateAliasInitParameters struct {

	// (String) Query used to limit documents the alias can access.
	// Query used to limit documents the alias can access.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	// Value used to route indexing operations to a specific shard. If specified, this overwrites the `routing` value for indexing operations.
	IndexRouting *string `json:"indexRouting,omitempty" tf:"index_routing,omitempty"`

	// (Boolean) If true, the alias is hidden.
	// If true, the alias is hidden.
	IsHidden *bool `json:"isHidden,omitempty" tf:"is_hidden,omitempty"`

	// (Boolean) If true, the index is the write index for the alias.
	// If true, the index is the write index for the alias.
	IsWriteIndex *bool `json:"isWriteIndex,omitempty" tf:"is_write_index,omitempty"`

	// (String) Name of the index template to create.
	// The alias name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Value used to route indexing and search operations to a specific shard.
	// Value used to route indexing and search operations to a specific shard.
	Routing *string `json:"routing,omitempty" tf:"routing,omitempty"`

	// (String) Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	SearchRouting *string `json:"searchRouting,omitempty" tf:"search_routing,omitempty"`
}

type TemplateAliasObservation struct {

	// (String) Query used to limit documents the alias can access.
	// Query used to limit documents the alias can access.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	// Value used to route indexing operations to a specific shard. If specified, this overwrites the `routing` value for indexing operations.
	IndexRouting *string `json:"indexRouting,omitempty" tf:"index_routing,omitempty"`

	// (Boolean) If true, the alias is hidden.
	// If true, the alias is hidden.
	IsHidden *bool `json:"isHidden,omitempty" tf:"is_hidden,omitempty"`

	// (Boolean) If true, the index is the write index for the alias.
	// If true, the index is the write index for the alias.
	IsWriteIndex *bool `json:"isWriteIndex,omitempty" tf:"is_write_index,omitempty"`

	// (String) Name of the index template to create.
	// The alias name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Value used to route indexing and search operations to a specific shard.
	// Value used to route indexing and search operations to a specific shard.
	Routing *string `json:"routing,omitempty" tf:"routing,omitempty"`

	// (String) Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	SearchRouting *string `json:"searchRouting,omitempty" tf:"search_routing,omitempty"`
}

type TemplateAliasParameters struct {

	// (String) Query used to limit documents the alias can access.
	// Query used to limit documents the alias can access.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	// Value used to route indexing operations to a specific shard. If specified, this overwrites the `routing` value for indexing operations.
	// +kubebuilder:validation:Optional
	IndexRouting *string `json:"indexRouting,omitempty" tf:"index_routing,omitempty"`

	// (Boolean) If true, the alias is hidden.
	// If true, the alias is hidden.
	// +kubebuilder:validation:Optional
	IsHidden *bool `json:"isHidden,omitempty" tf:"is_hidden,omitempty"`

	// (Boolean) If true, the index is the write index for the alias.
	// If true, the index is the write index for the alias.
	// +kubebuilder:validation:Optional
	IsWriteIndex *bool `json:"isWriteIndex,omitempty" tf:"is_write_index,omitempty"`

	// (String) Name of the index template to create.
	// The alias name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Value used to route indexing and search operations to a specific shard.
	// Value used to route indexing and search operations to a specific shard.
	// +kubebuilder:validation:Optional
	Routing *string `json:"routing,omitempty" tf:"routing,omitempty"`

	// (String) Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// +kubebuilder:validation:Optional
	SearchRouting *string `json:"searchRouting,omitempty" tf:"search_routing,omitempty"`
}

// IndexTemplateSpec defines the desired state of IndexTemplate
type IndexTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IndexTemplateParameters `json:"forProvider"`
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
	InitProvider IndexTemplateInitParameters `json:"initProvider,omitempty"`
}

// IndexTemplateStatus defines the observed state of IndexTemplate.
type IndexTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IndexTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IndexTemplate is the Schema for the IndexTemplates API. Creates or updates an index template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type IndexTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.indexPatterns) || (has(self.initProvider) && has(self.initProvider.indexPatterns))",message="spec.forProvider.indexPatterns is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IndexTemplateSpec   `json:"spec"`
	Status IndexTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IndexTemplateList contains a list of IndexTemplates
type IndexTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IndexTemplate `json:"items"`
}

// Repository type metadata.
var (
	IndexTemplate_Kind             = "IndexTemplate"
	IndexTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IndexTemplate_Kind}.String()
	IndexTemplate_KindAPIVersion   = IndexTemplate_Kind + "." + CRDGroupVersion.String()
	IndexTemplate_GroupVersionKind = CRDGroupVersion.WithKind(IndexTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&IndexTemplate{}, &IndexTemplateList{})
}
