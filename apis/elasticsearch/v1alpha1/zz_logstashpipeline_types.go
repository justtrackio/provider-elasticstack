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

type LogstashPipelineElasticsearchConnectionInitParameters struct {

	// (String, Sensitive) API Key to use for authentication to Elasticsearch
	// API Key to use for authentication to Elasticsearch
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// (String, Sensitive) Bearer Token to use for authentication to Elasticsearch
	// Bearer Token to use for authentication to Elasticsearch
	BearerTokenSecretRef *v1.SecretKeySelector `json:"bearerTokenSecretRef,omitempty" tf:"-"`

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

	Endpoints []*string `json:"endpointsSecretRef,omitempty" tf:"-"`

	// (String, Sensitive) ES Client Authentication field to be used with the bearer token
	// ES Client Authentication field to be used with the bearer token
	EsClientAuthenticationSecretRef *v1.SecretKeySelector `json:"esClientAuthenticationSecretRef,omitempty" tf:"-"`

	// (Boolean) Disable TLS certificate validation
	// Disable TLS certificate validation
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// (String, Sensitive) PEM encoded private key for client auth
	// PEM encoded private key for client auth
	KeyDataSecretRef *v1.SecretKeySelector `json:"keyDataSecretRef,omitempty" tf:"-"`

	// (String) Path to a file containing the PEM encoded private key for client auth
	// Path to a file containing the PEM encoded private key for client auth
	KeyFile *string `json:"keyFile,omitempty" tf:"key_file,omitempty"`

	// (String, Sensitive) Password to use for API authentication to Elasticsearch.
	// Password to use for API authentication to Elasticsearch.
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String) User who last updated the pipeline.
	// Username to use for API authentication to Elasticsearch.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LogstashPipelineElasticsearchConnectionObservation struct {

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

	// (String) User who last updated the pipeline.
	// Username to use for API authentication to Elasticsearch.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LogstashPipelineElasticsearchConnectionParameters struct {

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

	// (String) User who last updated the pipeline.
	// Username to use for API authentication to Elasticsearch.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LogstashPipelineInitParameters struct {

	// (String) Description of the pipeline.
	// Description of the pipeline.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []LogstashPipelineElasticsearchConnectionInitParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) Configuration for the pipeline.
	// Configuration for the pipeline.
	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// (Number) Time in milliseconds to wait for each event before sending an undersized batch to pipeline workers.
	// Time in milliseconds to wait for each event before sending an undersized batch to pipeline workers.
	PipelineBatchDelay *float64 `json:"pipelineBatchDelay,omitempty" tf:"pipeline_batch_delay,omitempty"`

	// (Number) The maximum number of events an individual worker thread collects before executing filters and outputs.
	// The maximum number of events an individual worker thread collects before executing filters and outputs.
	PipelineBatchSize *float64 `json:"pipelineBatchSize,omitempty" tf:"pipeline_batch_size,omitempty"`

	// (String) Sets the pipeline default value for ecs_compatibility, a setting that is available to plugins that implement an ECS compatibility mode for use with the Elastic Common Schema.
	// Sets the pipeline default value for ecs_compatibility, a setting that is available to plugins that implement an ECS compatibility mode for use with the Elastic Common Schema.
	PipelineEcsCompatibility *string `json:"pipelineEcsCompatibility,omitempty" tf:"pipeline_ecs_compatibility,omitempty"`

	// (String) Identifier for the pipeline.
	// Identifier for the pipeline.
	PipelineID *string `json:"pipelineId,omitempty" tf:"pipeline_id,omitempty"`

	// (String) Optional JSON metadata about the pipeline.
	// Optional JSON metadata about the pipeline.
	PipelineMetadata *string `json:"pipelineMetadata,omitempty" tf:"pipeline_metadata,omitempty"`

	// (String) Set the pipeline event ordering.
	// Set the pipeline event ordering.
	PipelineOrdered *string `json:"pipelineOrdered,omitempty" tf:"pipeline_ordered,omitempty"`

	// (Boolean) (Beta) Load Java plugins in independent classloaders to isolate their dependencies.
	// (Beta) Load Java plugins in independent classloaders to isolate their dependencies.
	PipelinePluginClassloaders *bool `json:"pipelinePluginClassloaders,omitempty" tf:"pipeline_plugin_classloaders,omitempty"`

	// (Boolean) Forces Logstash to exit during shutdown even if there are still inflight events in memory.
	// Forces Logstash to exit during shutdown even if there are still inflight events in memory.
	PipelineUnsafeShutdown *bool `json:"pipelineUnsafeShutdown,omitempty" tf:"pipeline_unsafe_shutdown,omitempty"`

	// (Number) The number of parallel workers used to run the filter and output stages of the pipeline.
	// The number of parallel workers used to run the filter and output stages of the pipeline.
	PipelineWorkers *float64 `json:"pipelineWorkers,omitempty" tf:"pipeline_workers,omitempty"`

	// (Number) The maximum number of ACKed events before forcing a checkpoint when persistent queues are enabled.
	// The maximum number of ACKed events before forcing a checkpoint when persistent queues are enabled.
	QueueCheckpointAcks *float64 `json:"queueCheckpointAcks,omitempty" tf:"queue_checkpoint_acks,omitempty"`

	// (Boolean) When enabled, Logstash will retry four times per attempted checkpoint write for any checkpoint writes that fail. Any subsequent errors are not retried.
	// When enabled, Logstash will retry four times per attempted checkpoint write for any checkpoint writes that fail. Any subsequent errors are not retried.
	QueueCheckpointRetry *bool `json:"queueCheckpointRetry,omitempty" tf:"queue_checkpoint_retry,omitempty"`

	// (Number) The maximum number of written events before forcing a checkpoint when persistent queues are enabled.
	// The maximum number of written events before forcing a checkpoint when persistent queues are enabled.
	QueueCheckpointWrites *float64 `json:"queueCheckpointWrites,omitempty" tf:"queue_checkpoint_writes,omitempty"`

	// (Boolean) When enabled, Logstash waits until the persistent queue is drained before shutting down.
	// When enabled, Logstash waits until the persistent queue is drained before shutting down.
	QueueDrain *bool `json:"queueDrain,omitempty" tf:"queue_drain,omitempty"`

	// (String) Units for the total capacity of the queue when persistent queues are enabled.
	// Units for the total capacity of the queue when persistent queues are enabled.
	QueueMaxBytes *string `json:"queueMaxBytes,omitempty" tf:"queue_max_bytes,omitempty"`

	// (Number) The maximum number of unread events in the queue when persistent queues are enabled.
	// The maximum number of unread events in the queue when persistent queues are enabled.
	QueueMaxEvents *float64 `json:"queueMaxEvents,omitempty" tf:"queue_max_events,omitempty"`

	// only data files separated into pages.
	// The size of the page data files used when persistent queues are enabled. The queue data consists of append-only data files separated into pages.
	QueuePageCapacity *string `json:"queuePageCapacity,omitempty" tf:"queue_page_capacity,omitempty"`

	// memory queueing, or persisted for disk-based acknowledged queueing.
	// The internal queueing model for event buffering. Options are memory for in-memory queueing, or persisted for disk-based acknowledged queueing.
	QueueType *string `json:"queueType,omitempty" tf:"queue_type,omitempty"`

	// (String) User who last updated the pipeline.
	// User who last updated the pipeline.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LogstashPipelineObservation struct {

	// (String) Description of the pipeline.
	// Description of the pipeline.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []LogstashPipelineElasticsearchConnectionObservation `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) Internal identifier of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Date the pipeline was last updated.
	// Date the pipeline was last updated.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// (String) Configuration for the pipeline.
	// Configuration for the pipeline.
	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// (Number) Time in milliseconds to wait for each event before sending an undersized batch to pipeline workers.
	// Time in milliseconds to wait for each event before sending an undersized batch to pipeline workers.
	PipelineBatchDelay *float64 `json:"pipelineBatchDelay,omitempty" tf:"pipeline_batch_delay,omitempty"`

	// (Number) The maximum number of events an individual worker thread collects before executing filters and outputs.
	// The maximum number of events an individual worker thread collects before executing filters and outputs.
	PipelineBatchSize *float64 `json:"pipelineBatchSize,omitempty" tf:"pipeline_batch_size,omitempty"`

	// (String) Sets the pipeline default value for ecs_compatibility, a setting that is available to plugins that implement an ECS compatibility mode for use with the Elastic Common Schema.
	// Sets the pipeline default value for ecs_compatibility, a setting that is available to plugins that implement an ECS compatibility mode for use with the Elastic Common Schema.
	PipelineEcsCompatibility *string `json:"pipelineEcsCompatibility,omitempty" tf:"pipeline_ecs_compatibility,omitempty"`

	// (String) Identifier for the pipeline.
	// Identifier for the pipeline.
	PipelineID *string `json:"pipelineId,omitempty" tf:"pipeline_id,omitempty"`

	// (String) Optional JSON metadata about the pipeline.
	// Optional JSON metadata about the pipeline.
	PipelineMetadata *string `json:"pipelineMetadata,omitempty" tf:"pipeline_metadata,omitempty"`

	// (String) Set the pipeline event ordering.
	// Set the pipeline event ordering.
	PipelineOrdered *string `json:"pipelineOrdered,omitempty" tf:"pipeline_ordered,omitempty"`

	// (Boolean) (Beta) Load Java plugins in independent classloaders to isolate their dependencies.
	// (Beta) Load Java plugins in independent classloaders to isolate their dependencies.
	PipelinePluginClassloaders *bool `json:"pipelinePluginClassloaders,omitempty" tf:"pipeline_plugin_classloaders,omitempty"`

	// (Boolean) Forces Logstash to exit during shutdown even if there are still inflight events in memory.
	// Forces Logstash to exit during shutdown even if there are still inflight events in memory.
	PipelineUnsafeShutdown *bool `json:"pipelineUnsafeShutdown,omitempty" tf:"pipeline_unsafe_shutdown,omitempty"`

	// (Number) The number of parallel workers used to run the filter and output stages of the pipeline.
	// The number of parallel workers used to run the filter and output stages of the pipeline.
	PipelineWorkers *float64 `json:"pipelineWorkers,omitempty" tf:"pipeline_workers,omitempty"`

	// (Number) The maximum number of ACKed events before forcing a checkpoint when persistent queues are enabled.
	// The maximum number of ACKed events before forcing a checkpoint when persistent queues are enabled.
	QueueCheckpointAcks *float64 `json:"queueCheckpointAcks,omitempty" tf:"queue_checkpoint_acks,omitempty"`

	// (Boolean) When enabled, Logstash will retry four times per attempted checkpoint write for any checkpoint writes that fail. Any subsequent errors are not retried.
	// When enabled, Logstash will retry four times per attempted checkpoint write for any checkpoint writes that fail. Any subsequent errors are not retried.
	QueueCheckpointRetry *bool `json:"queueCheckpointRetry,omitempty" tf:"queue_checkpoint_retry,omitempty"`

	// (Number) The maximum number of written events before forcing a checkpoint when persistent queues are enabled.
	// The maximum number of written events before forcing a checkpoint when persistent queues are enabled.
	QueueCheckpointWrites *float64 `json:"queueCheckpointWrites,omitempty" tf:"queue_checkpoint_writes,omitempty"`

	// (Boolean) When enabled, Logstash waits until the persistent queue is drained before shutting down.
	// When enabled, Logstash waits until the persistent queue is drained before shutting down.
	QueueDrain *bool `json:"queueDrain,omitempty" tf:"queue_drain,omitempty"`

	// (String) Units for the total capacity of the queue when persistent queues are enabled.
	// Units for the total capacity of the queue when persistent queues are enabled.
	QueueMaxBytes *string `json:"queueMaxBytes,omitempty" tf:"queue_max_bytes,omitempty"`

	// (Number) The maximum number of unread events in the queue when persistent queues are enabled.
	// The maximum number of unread events in the queue when persistent queues are enabled.
	QueueMaxEvents *float64 `json:"queueMaxEvents,omitempty" tf:"queue_max_events,omitempty"`

	// only data files separated into pages.
	// The size of the page data files used when persistent queues are enabled. The queue data consists of append-only data files separated into pages.
	QueuePageCapacity *string `json:"queuePageCapacity,omitempty" tf:"queue_page_capacity,omitempty"`

	// memory queueing, or persisted for disk-based acknowledged queueing.
	// The internal queueing model for event buffering. Options are memory for in-memory queueing, or persisted for disk-based acknowledged queueing.
	QueueType *string `json:"queueType,omitempty" tf:"queue_type,omitempty"`

	// (String) User who last updated the pipeline.
	// User who last updated the pipeline.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LogstashPipelineParameters struct {

	// (String) Description of the pipeline.
	// Description of the pipeline.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	// +kubebuilder:validation:Optional
	ElasticsearchConnection []LogstashPipelineElasticsearchConnectionParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) Configuration for the pipeline.
	// Configuration for the pipeline.
	// +kubebuilder:validation:Optional
	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// (Number) Time in milliseconds to wait for each event before sending an undersized batch to pipeline workers.
	// Time in milliseconds to wait for each event before sending an undersized batch to pipeline workers.
	// +kubebuilder:validation:Optional
	PipelineBatchDelay *float64 `json:"pipelineBatchDelay,omitempty" tf:"pipeline_batch_delay,omitempty"`

	// (Number) The maximum number of events an individual worker thread collects before executing filters and outputs.
	// The maximum number of events an individual worker thread collects before executing filters and outputs.
	// +kubebuilder:validation:Optional
	PipelineBatchSize *float64 `json:"pipelineBatchSize,omitempty" tf:"pipeline_batch_size,omitempty"`

	// (String) Sets the pipeline default value for ecs_compatibility, a setting that is available to plugins that implement an ECS compatibility mode for use with the Elastic Common Schema.
	// Sets the pipeline default value for ecs_compatibility, a setting that is available to plugins that implement an ECS compatibility mode for use with the Elastic Common Schema.
	// +kubebuilder:validation:Optional
	PipelineEcsCompatibility *string `json:"pipelineEcsCompatibility,omitempty" tf:"pipeline_ecs_compatibility,omitempty"`

	// (String) Identifier for the pipeline.
	// Identifier for the pipeline.
	// +kubebuilder:validation:Optional
	PipelineID *string `json:"pipelineId,omitempty" tf:"pipeline_id,omitempty"`

	// (String) Optional JSON metadata about the pipeline.
	// Optional JSON metadata about the pipeline.
	// +kubebuilder:validation:Optional
	PipelineMetadata *string `json:"pipelineMetadata,omitempty" tf:"pipeline_metadata,omitempty"`

	// (String) Set the pipeline event ordering.
	// Set the pipeline event ordering.
	// +kubebuilder:validation:Optional
	PipelineOrdered *string `json:"pipelineOrdered,omitempty" tf:"pipeline_ordered,omitempty"`

	// (Boolean) (Beta) Load Java plugins in independent classloaders to isolate their dependencies.
	// (Beta) Load Java plugins in independent classloaders to isolate their dependencies.
	// +kubebuilder:validation:Optional
	PipelinePluginClassloaders *bool `json:"pipelinePluginClassloaders,omitempty" tf:"pipeline_plugin_classloaders,omitempty"`

	// (Boolean) Forces Logstash to exit during shutdown even if there are still inflight events in memory.
	// Forces Logstash to exit during shutdown even if there are still inflight events in memory.
	// +kubebuilder:validation:Optional
	PipelineUnsafeShutdown *bool `json:"pipelineUnsafeShutdown,omitempty" tf:"pipeline_unsafe_shutdown,omitempty"`

	// (Number) The number of parallel workers used to run the filter and output stages of the pipeline.
	// The number of parallel workers used to run the filter and output stages of the pipeline.
	// +kubebuilder:validation:Optional
	PipelineWorkers *float64 `json:"pipelineWorkers,omitempty" tf:"pipeline_workers,omitempty"`

	// (Number) The maximum number of ACKed events before forcing a checkpoint when persistent queues are enabled.
	// The maximum number of ACKed events before forcing a checkpoint when persistent queues are enabled.
	// +kubebuilder:validation:Optional
	QueueCheckpointAcks *float64 `json:"queueCheckpointAcks,omitempty" tf:"queue_checkpoint_acks,omitempty"`

	// (Boolean) When enabled, Logstash will retry four times per attempted checkpoint write for any checkpoint writes that fail. Any subsequent errors are not retried.
	// When enabled, Logstash will retry four times per attempted checkpoint write for any checkpoint writes that fail. Any subsequent errors are not retried.
	// +kubebuilder:validation:Optional
	QueueCheckpointRetry *bool `json:"queueCheckpointRetry,omitempty" tf:"queue_checkpoint_retry,omitempty"`

	// (Number) The maximum number of written events before forcing a checkpoint when persistent queues are enabled.
	// The maximum number of written events before forcing a checkpoint when persistent queues are enabled.
	// +kubebuilder:validation:Optional
	QueueCheckpointWrites *float64 `json:"queueCheckpointWrites,omitempty" tf:"queue_checkpoint_writes,omitempty"`

	// (Boolean) When enabled, Logstash waits until the persistent queue is drained before shutting down.
	// When enabled, Logstash waits until the persistent queue is drained before shutting down.
	// +kubebuilder:validation:Optional
	QueueDrain *bool `json:"queueDrain,omitempty" tf:"queue_drain,omitempty"`

	// (String) Units for the total capacity of the queue when persistent queues are enabled.
	// Units for the total capacity of the queue when persistent queues are enabled.
	// +kubebuilder:validation:Optional
	QueueMaxBytes *string `json:"queueMaxBytes,omitempty" tf:"queue_max_bytes,omitempty"`

	// (Number) The maximum number of unread events in the queue when persistent queues are enabled.
	// The maximum number of unread events in the queue when persistent queues are enabled.
	// +kubebuilder:validation:Optional
	QueueMaxEvents *float64 `json:"queueMaxEvents,omitempty" tf:"queue_max_events,omitempty"`

	// only data files separated into pages.
	// The size of the page data files used when persistent queues are enabled. The queue data consists of append-only data files separated into pages.
	// +kubebuilder:validation:Optional
	QueuePageCapacity *string `json:"queuePageCapacity,omitempty" tf:"queue_page_capacity,omitempty"`

	// memory queueing, or persisted for disk-based acknowledged queueing.
	// The internal queueing model for event buffering. Options are memory for in-memory queueing, or persisted for disk-based acknowledged queueing.
	// +kubebuilder:validation:Optional
	QueueType *string `json:"queueType,omitempty" tf:"queue_type,omitempty"`

	// (String) User who last updated the pipeline.
	// User who last updated the pipeline.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// LogstashPipelineSpec defines the desired state of LogstashPipeline
type LogstashPipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogstashPipelineParameters `json:"forProvider"`
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
	InitProvider LogstashPipelineInitParameters `json:"initProvider,omitempty"`
}

// LogstashPipelineStatus defines the observed state of LogstashPipeline.
type LogstashPipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogstashPipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogstashPipeline is the Schema for the LogstashPipelines API. Creates or updates centrally managed logstash pipelines.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type LogstashPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pipeline) || (has(self.initProvider) && has(self.initProvider.pipeline))",message="spec.forProvider.pipeline is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pipelineId) || (has(self.initProvider) && has(self.initProvider.pipelineId))",message="spec.forProvider.pipelineId is a required parameter"
	Spec   LogstashPipelineSpec   `json:"spec"`
	Status LogstashPipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogstashPipelineList contains a list of LogstashPipelines
type LogstashPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogstashPipeline `json:"items"`
}

// Repository type metadata.
var (
	LogstashPipeline_Kind             = "LogstashPipeline"
	LogstashPipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogstashPipeline_Kind}.String()
	LogstashPipeline_KindAPIVersion   = LogstashPipeline_Kind + "." + CRDGroupVersion.String()
	LogstashPipeline_GroupVersionKind = CRDGroupVersion.WithKind(LogstashPipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&LogstashPipeline{}, &LogstashPipelineList{})
}
