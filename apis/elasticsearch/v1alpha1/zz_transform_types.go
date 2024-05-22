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

type DestinationInitParameters struct {

	// (String) The destination index for the transform.
	// The destination index for the transform.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String) The unique identifier for an ingest pipeline.
	// The unique identifier for an ingest pipeline.
	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`
}

type DestinationObservation struct {

	// (String) The destination index for the transform.
	// The destination index for the transform.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String) The unique identifier for an ingest pipeline.
	// The unique identifier for an ingest pipeline.
	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`
}

type DestinationParameters struct {

	// (String) The destination index for the transform.
	// The destination index for the transform.
	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// (String) The unique identifier for an ingest pipeline.
	// The unique identifier for an ingest pipeline.
	// +kubebuilder:validation:Optional
	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`
}

type RetentionPolicyInitParameters struct {

	// (Block List, Min: 1, Max: 1) Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option. (see below for nested schema)
	// Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option.
	Time []TimeInitParameters `json:"time,omitempty" tf:"time,omitempty"`
}

type RetentionPolicyObservation struct {

	// (Block List, Min: 1, Max: 1) Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option. (see below for nested schema)
	// Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option.
	Time []TimeObservation `json:"time,omitempty" tf:"time,omitempty"`
}

type RetentionPolicyParameters struct {

	// (Block List, Min: 1, Max: 1) Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option. (see below for nested schema)
	// Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option.
	// +kubebuilder:validation:Optional
	Time []TimeParameters `json:"time" tf:"time,omitempty"`
}

type SourceInitParameters struct {

	// (List of String) The source indices for the transform.
	// The source indices for the transform.
	Indices []*string `json:"indices,omitempty" tf:"indices,omitempty"`

	// (String) A query clause that retrieves a subset of data from the source index.
	// A query clause that retrieves a subset of data from the source index.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// time runtime fields that can be used by the transform.
	// Definitions of search-time runtime fields that can be used by the transform.
	RuntimeMappings *string `json:"runtimeMappings,omitempty" tf:"runtime_mappings,omitempty"`
}

type SourceObservation struct {

	// (List of String) The source indices for the transform.
	// The source indices for the transform.
	Indices []*string `json:"indices,omitempty" tf:"indices,omitempty"`

	// (String) A query clause that retrieves a subset of data from the source index.
	// A query clause that retrieves a subset of data from the source index.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// time runtime fields that can be used by the transform.
	// Definitions of search-time runtime fields that can be used by the transform.
	RuntimeMappings *string `json:"runtimeMappings,omitempty" tf:"runtime_mappings,omitempty"`
}

type SourceParameters struct {

	// (List of String) The source indices for the transform.
	// The source indices for the transform.
	// +kubebuilder:validation:Optional
	Indices []*string `json:"indices" tf:"indices,omitempty"`

	// (String) A query clause that retrieves a subset of data from the source index.
	// A query clause that retrieves a subset of data from the source index.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// time runtime fields that can be used by the transform.
	// Definitions of search-time runtime fields that can be used by the transform.
	// +kubebuilder:validation:Optional
	RuntimeMappings *string `json:"runtimeMappings,omitempty" tf:"runtime_mappings,omitempty"`
}

type SyncInitParameters struct {

	// (Block List, Min: 1, Max: 1) Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option. (see below for nested schema)
	// Specifies that the transform uses a time field to synchronize the source and destination indices. This is currently the only supported option.
	Time []SyncTimeInitParameters `json:"time,omitempty" tf:"time,omitempty"`
}

type SyncObservation struct {

	// (Block List, Min: 1, Max: 1) Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option. (see below for nested schema)
	// Specifies that the transform uses a time field to synchronize the source and destination indices. This is currently the only supported option.
	Time []SyncTimeObservation `json:"time,omitempty" tf:"time,omitempty"`
}

type SyncParameters struct {

	// (Block List, Min: 1, Max: 1) Specifies that the transform uses a time field to set the retention policy. This is currently the only supported option. (see below for nested schema)
	// Specifies that the transform uses a time field to synchronize the source and destination indices. This is currently the only supported option.
	// +kubebuilder:validation:Optional
	Time []SyncTimeParameters `json:"time" tf:"time,omitempty"`
}

type SyncTimeInitParameters struct {

	// (String) The time delay between the current time and the latest input data time. The default value is 60s.
	// The time delay between the current time and the latest input data time. The default value is 60s.
	Delay *string `json:"delay,omitempty" tf:"delay,omitempty"`

	// (String) The date field that is used to calculate the age of the document.
	// The date field that is used to identify new documents in the source.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`
}

type SyncTimeObservation struct {

	// (String) The time delay between the current time and the latest input data time. The default value is 60s.
	// The time delay between the current time and the latest input data time. The default value is 60s.
	Delay *string `json:"delay,omitempty" tf:"delay,omitempty"`

	// (String) The date field that is used to calculate the age of the document.
	// The date field that is used to identify new documents in the source.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`
}

type SyncTimeParameters struct {

	// (String) The time delay between the current time and the latest input data time. The default value is 60s.
	// The time delay between the current time and the latest input data time. The default value is 60s.
	// +kubebuilder:validation:Optional
	Delay *string `json:"delay,omitempty" tf:"delay,omitempty"`

	// (String) The date field that is used to calculate the age of the document.
	// The date field that is used to identify new documents in the source.
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`
}

type TimeInitParameters struct {

	// (String) The date field that is used to calculate the age of the document.
	// The date field that is used to calculate the age of the document.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String) Specifies the maximum age of a document in the destination index.
	// Specifies the maximum age of a document in the destination index.
	MaxAge *string `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type TimeObservation struct {

	// (String) The date field that is used to calculate the age of the document.
	// The date field that is used to calculate the age of the document.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String) Specifies the maximum age of a document in the destination index.
	// Specifies the maximum age of a document in the destination index.
	MaxAge *string `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type TimeParameters struct {

	// (String) The date field that is used to calculate the age of the document.
	// The date field that is used to calculate the age of the document.
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// (String) Specifies the maximum age of a document in the destination index.
	// Specifies the maximum age of a document in the destination index.
	// +kubebuilder:validation:Optional
	MaxAge *string `json:"maxAge" tf:"max_age,omitempty"`
}

type TransformInitParameters struct {

	// (Boolean) Specifies whether the transform checkpoint ranges should be optimized for performance.
	// Specifies whether the transform checkpoint ranges should be optimized for performance.
	AlignCheckpoints *bool `json:"alignCheckpoints,omitempty" tf:"align_checkpoints,omitempty"`

	// (Boolean) Defines if dates in the output should be written as ISO formatted string (default) or as millis since epoch.
	// Defines if dates in the output should be written as ISO formatted string (default) or as millis since epoch.
	DatesAsEpochMillis *bool `json:"datesAsEpochMillis,omitempty" tf:"dates_as_epoch_millis,omitempty"`

	// (Boolean) Specifies whether the transform should deduce the destination index mappings from the transform config.
	// Specifies whether the transform should deduce the destination index mappings from the transform config.
	DeduceMappings *bool `json:"deduceMappings,omitempty" tf:"deduce_mappings,omitempty"`

	// (Boolean) When true, deferrable validations are not run upon creation, but rather when the transform is started. This behavior may be desired if the source index does not exist until after the transform is created. Default is false
	// When true, deferrable validations are not run upon creation, but rather when the transform is started. This behavior may be desired if the source index does not exist until after the transform is created. Default is `false`
	DeferValidation *bool `json:"deferValidation,omitempty" tf:"defer_validation,omitempty"`

	// (String) Free text description of the transform.
	// Free text description of the transform.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Min: 1, Max: 1) The destination for the transform. (see below for nested schema)
	// The destination for the transform.
	Destination []DestinationInitParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// (Number) Specifies a limit on the number of input documents per second. Default (unset) value disables throttling.
	// Specifies a limit on the number of input documents per second. Default (unset) value disables throttling.
	DocsPerSecond *float64 `json:"docsPerSecond,omitempty" tf:"docs_per_second,omitempty"`

	// (Boolean) Controls wether the transform should be started or stopped. Default is false (stopped).
	// Controls wether the transform should be started or stopped. Default is `false` (stopped).
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The interval between checks for changes in the source indices when the transform is running continuously. Defaults to 1m.
	// The interval between checks for changes in the source indices when the transform is running continuously. Defaults to `1m`.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String) The latest method transforms the data by finding the latest document for each unique key. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// The latest method transforms the data by finding the latest document for each unique key. JSON definition expected. Either 'pivot' or 'latest' must be present.
	Latest *string `json:"latest,omitempty" tf:"latest,omitempty"`

	// (Number) Defines the initial page size to use for the composite aggregation for each checkpoint. Default is 500.
	// Defines the initial page size to use for the composite aggregation for each checkpoint. Default is 500.
	MaxPageSearchSize *float64 `json:"maxPageSearchSize,omitempty" tf:"max_page_search_size,omitempty"`

	// (String) Defines optional transform metadata.
	// Defines optional transform metadata.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the transform you wish to create.
	// Name of the transform you wish to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// level setting num_transform_failure_retries.
	// Defines the number of retries on a recoverable failure before the transform task is marked as failed. The default value is the cluster-level setting num_transform_failure_retries.
	NumFailureRetries *float64 `json:"numFailureRetries,omitempty" tf:"num_failure_retries,omitempty"`

	// (String) The pivot method transforms the data by aggregating and grouping it. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// The pivot method transforms the data by aggregating and grouping it. JSON definition expected. Either 'pivot' or 'latest' must be present.
	Pivot *string `json:"pivot,omitempty" tf:"pivot,omitempty"`

	// (Block List, Max: 1) Defines a retention policy for the transform. (see below for nested schema)
	// Defines a retention policy for the transform.
	RetentionPolicy []RetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// (Block List, Min: 1, Max: 1) The source of the data for the transform. (see below for nested schema)
	// The source of the data for the transform.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`

	// (Block List, Max: 1) Defines the properties transforms require to run continuously. (see below for nested schema)
	// Defines the properties transforms require to run continuously.
	Sync []SyncInitParameters `json:"sync,omitempty" tf:"sync,omitempty"`

	// (String) Period to wait for a response from Elastisearch when performing any management operation. If no response is received before the timeout expires, the operation fails and returns an error. Defaults to 30s.
	// Period to wait for a response from Elastisearch when performing any management operation. If no response is received before the timeout expires, the operation fails and returns an error. Defaults to `30s`.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Boolean) In unattended mode, the transform retries indefinitely in case of an error which means the transform never fails.
	// In unattended mode, the transform retries indefinitely in case of an error which means the transform never fails.
	Unattended *bool `json:"unattended,omitempty" tf:"unattended,omitempty"`
}

type TransformObservation struct {

	// (Boolean) Specifies whether the transform checkpoint ranges should be optimized for performance.
	// Specifies whether the transform checkpoint ranges should be optimized for performance.
	AlignCheckpoints *bool `json:"alignCheckpoints,omitempty" tf:"align_checkpoints,omitempty"`

	// (Boolean) Defines if dates in the output should be written as ISO formatted string (default) or as millis since epoch.
	// Defines if dates in the output should be written as ISO formatted string (default) or as millis since epoch.
	DatesAsEpochMillis *bool `json:"datesAsEpochMillis,omitempty" tf:"dates_as_epoch_millis,omitempty"`

	// (Boolean) Specifies whether the transform should deduce the destination index mappings from the transform config.
	// Specifies whether the transform should deduce the destination index mappings from the transform config.
	DeduceMappings *bool `json:"deduceMappings,omitempty" tf:"deduce_mappings,omitempty"`

	// (Boolean) When true, deferrable validations are not run upon creation, but rather when the transform is started. This behavior may be desired if the source index does not exist until after the transform is created. Default is false
	// When true, deferrable validations are not run upon creation, but rather when the transform is started. This behavior may be desired if the source index does not exist until after the transform is created. Default is `false`
	DeferValidation *bool `json:"deferValidation,omitempty" tf:"defer_validation,omitempty"`

	// (String) Free text description of the transform.
	// Free text description of the transform.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Min: 1, Max: 1) The destination for the transform. (see below for nested schema)
	// The destination for the transform.
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// (Number) Specifies a limit on the number of input documents per second. Default (unset) value disables throttling.
	// Specifies a limit on the number of input documents per second. Default (unset) value disables throttling.
	DocsPerSecond *float64 `json:"docsPerSecond,omitempty" tf:"docs_per_second,omitempty"`

	// (Boolean) Controls wether the transform should be started or stopped. Default is false (stopped).
	// Controls wether the transform should be started or stopped. Default is `false` (stopped).
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The interval between checks for changes in the source indices when the transform is running continuously. Defaults to 1m.
	// The interval between checks for changes in the source indices when the transform is running continuously. Defaults to `1m`.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String) Internal identifier of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The latest method transforms the data by finding the latest document for each unique key. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// The latest method transforms the data by finding the latest document for each unique key. JSON definition expected. Either 'pivot' or 'latest' must be present.
	Latest *string `json:"latest,omitempty" tf:"latest,omitempty"`

	// (Number) Defines the initial page size to use for the composite aggregation for each checkpoint. Default is 500.
	// Defines the initial page size to use for the composite aggregation for each checkpoint. Default is 500.
	MaxPageSearchSize *float64 `json:"maxPageSearchSize,omitempty" tf:"max_page_search_size,omitempty"`

	// (String) Defines optional transform metadata.
	// Defines optional transform metadata.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the transform you wish to create.
	// Name of the transform you wish to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// level setting num_transform_failure_retries.
	// Defines the number of retries on a recoverable failure before the transform task is marked as failed. The default value is the cluster-level setting num_transform_failure_retries.
	NumFailureRetries *float64 `json:"numFailureRetries,omitempty" tf:"num_failure_retries,omitempty"`

	// (String) The pivot method transforms the data by aggregating and grouping it. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// The pivot method transforms the data by aggregating and grouping it. JSON definition expected. Either 'pivot' or 'latest' must be present.
	Pivot *string `json:"pivot,omitempty" tf:"pivot,omitempty"`

	// (Block List, Max: 1) Defines a retention policy for the transform. (see below for nested schema)
	// Defines a retention policy for the transform.
	RetentionPolicy []RetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// (Block List, Min: 1, Max: 1) The source of the data for the transform. (see below for nested schema)
	// The source of the data for the transform.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// (Block List, Max: 1) Defines the properties transforms require to run continuously. (see below for nested schema)
	// Defines the properties transforms require to run continuously.
	Sync []SyncObservation `json:"sync,omitempty" tf:"sync,omitempty"`

	// (String) Period to wait for a response from Elastisearch when performing any management operation. If no response is received before the timeout expires, the operation fails and returns an error. Defaults to 30s.
	// Period to wait for a response from Elastisearch when performing any management operation. If no response is received before the timeout expires, the operation fails and returns an error. Defaults to `30s`.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Boolean) In unattended mode, the transform retries indefinitely in case of an error which means the transform never fails.
	// In unattended mode, the transform retries indefinitely in case of an error which means the transform never fails.
	Unattended *bool `json:"unattended,omitempty" tf:"unattended,omitempty"`
}

type TransformParameters struct {

	// (Boolean) Specifies whether the transform checkpoint ranges should be optimized for performance.
	// Specifies whether the transform checkpoint ranges should be optimized for performance.
	// +kubebuilder:validation:Optional
	AlignCheckpoints *bool `json:"alignCheckpoints,omitempty" tf:"align_checkpoints,omitempty"`

	// (Boolean) Defines if dates in the output should be written as ISO formatted string (default) or as millis since epoch.
	// Defines if dates in the output should be written as ISO formatted string (default) or as millis since epoch.
	// +kubebuilder:validation:Optional
	DatesAsEpochMillis *bool `json:"datesAsEpochMillis,omitempty" tf:"dates_as_epoch_millis,omitempty"`

	// (Boolean) Specifies whether the transform should deduce the destination index mappings from the transform config.
	// Specifies whether the transform should deduce the destination index mappings from the transform config.
	// +kubebuilder:validation:Optional
	DeduceMappings *bool `json:"deduceMappings,omitempty" tf:"deduce_mappings,omitempty"`

	// (Boolean) When true, deferrable validations are not run upon creation, but rather when the transform is started. This behavior may be desired if the source index does not exist until after the transform is created. Default is false
	// When true, deferrable validations are not run upon creation, but rather when the transform is started. This behavior may be desired if the source index does not exist until after the transform is created. Default is `false`
	// +kubebuilder:validation:Optional
	DeferValidation *bool `json:"deferValidation,omitempty" tf:"defer_validation,omitempty"`

	// (String) Free text description of the transform.
	// Free text description of the transform.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Min: 1, Max: 1) The destination for the transform. (see below for nested schema)
	// The destination for the transform.
	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// (Number) Specifies a limit on the number of input documents per second. Default (unset) value disables throttling.
	// Specifies a limit on the number of input documents per second. Default (unset) value disables throttling.
	// +kubebuilder:validation:Optional
	DocsPerSecond *float64 `json:"docsPerSecond,omitempty" tf:"docs_per_second,omitempty"`

	// (Boolean) Controls wether the transform should be started or stopped. Default is false (stopped).
	// Controls wether the transform should be started or stopped. Default is `false` (stopped).
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The interval between checks for changes in the source indices when the transform is running continuously. Defaults to 1m.
	// The interval between checks for changes in the source indices when the transform is running continuously. Defaults to `1m`.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String) The latest method transforms the data by finding the latest document for each unique key. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// The latest method transforms the data by finding the latest document for each unique key. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// +kubebuilder:validation:Optional
	Latest *string `json:"latest,omitempty" tf:"latest,omitempty"`

	// (Number) Defines the initial page size to use for the composite aggregation for each checkpoint. Default is 500.
	// Defines the initial page size to use for the composite aggregation for each checkpoint. Default is 500.
	// +kubebuilder:validation:Optional
	MaxPageSearchSize *float64 `json:"maxPageSearchSize,omitempty" tf:"max_page_search_size,omitempty"`

	// (String) Defines optional transform metadata.
	// Defines optional transform metadata.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the transform you wish to create.
	// Name of the transform you wish to create.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// level setting num_transform_failure_retries.
	// Defines the number of retries on a recoverable failure before the transform task is marked as failed. The default value is the cluster-level setting num_transform_failure_retries.
	// +kubebuilder:validation:Optional
	NumFailureRetries *float64 `json:"numFailureRetries,omitempty" tf:"num_failure_retries,omitempty"`

	// (String) The pivot method transforms the data by aggregating and grouping it. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// The pivot method transforms the data by aggregating and grouping it. JSON definition expected. Either 'pivot' or 'latest' must be present.
	// +kubebuilder:validation:Optional
	Pivot *string `json:"pivot,omitempty" tf:"pivot,omitempty"`

	// (Block List, Max: 1) Defines a retention policy for the transform. (see below for nested schema)
	// Defines a retention policy for the transform.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// (Block List, Min: 1, Max: 1) The source of the data for the transform. (see below for nested schema)
	// The source of the data for the transform.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// (Block List, Max: 1) Defines the properties transforms require to run continuously. (see below for nested schema)
	// Defines the properties transforms require to run continuously.
	// +kubebuilder:validation:Optional
	Sync []SyncParameters `json:"sync,omitempty" tf:"sync,omitempty"`

	// (String) Period to wait for a response from Elastisearch when performing any management operation. If no response is received before the timeout expires, the operation fails and returns an error. Defaults to 30s.
	// Period to wait for a response from Elastisearch when performing any management operation. If no response is received before the timeout expires, the operation fails and returns an error. Defaults to `30s`.
	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Boolean) In unattended mode, the transform retries indefinitely in case of an error which means the transform never fails.
	// In unattended mode, the transform retries indefinitely in case of an error which means the transform never fails.
	// +kubebuilder:validation:Optional
	Unattended *bool `json:"unattended,omitempty" tf:"unattended,omitempty"`
}

// TransformSpec defines the desired state of Transform
type TransformSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransformParameters `json:"forProvider"`
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
	InitProvider TransformInitParameters `json:"initProvider,omitempty"`
}

// TransformStatus defines the observed state of Transform.
type TransformStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransformObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Transform is the Schema for the Transforms API. Manages transforms. Transforms enable you to convert existing Elasticsearch indices into summarized indices.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type Transform struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   TransformSpec   `json:"spec"`
	Status TransformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransformList contains a list of Transforms
type TransformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Transform `json:"items"`
}

// Repository type metadata.
var (
	Transform_Kind             = "Transform"
	Transform_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Transform_Kind}.String()
	Transform_KindAPIVersion   = Transform_Kind + "." + CRDGroupVersion.String()
	Transform_GroupVersionKind = CRDGroupVersion.WithKind(Transform_Kind)
)

func init() {
	SchemeBuilder.Register(&Transform{}, &TransformList{})
}
