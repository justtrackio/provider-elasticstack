/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"elasticstack_elasticsearch_cluster_settings":      config.IdentifierFromProvider,
	"elasticstack_elasticsearch_component_template":    config.IdentifierFromProvider,
	"elasticstack_elasticsearch_data_stream":           config.IdentifierFromProvider,
	"elasticstack_elasticsearch_enrich_policy":         config.IdentifierFromProvider,
	"elasticstack_elasticsearch_index":                 config.IdentifierFromProvider,
	"elasticstack_elasticsearch_index_lifecycle":       config.IdentifierFromProvider,
	"elasticstack_elasticsearch_index_template":        config.IdentifierFromProvider,
	"elasticstack_elasticsearch_ingest_pipeline":       config.IdentifierFromProvider,
	"elasticstack_elasticsearch_logstash_pipeline":     config.IdentifierFromProvider,
	"elasticstack_elasticsearch_script":                config.IdentifierFromProvider,
	"elasticstack_elasticsearch_security_api_key":      config.IdentifierFromProvider,
	"elasticstack_elasticsearch_security_role":         config.IdentifierFromProvider,
	"elasticstack_elasticsearch_security_role_mapping": config.IdentifierFromProvider,
	"elasticstack_elasticsearch_security_system_user":  config.IdentifierFromProvider,
	"elasticstack_elasticsearch_security_user":         config.IdentifierFromProvider,
	"elasticstack_elasticsearch_snapshot_lifecycle":    config.IdentifierFromProvider,
	"elasticstack_elasticsearch_snapshot_repository":   config.IdentifierFromProvider,
	"elasticstack_elasticsearch_transform":             config.IdentifierFromProvider,
	"elasticstack_elasticsearch_watch":                 config.IdentifierFromProvider,
	"elasticstack_fleet_agent_policy":                  config.IdentifierFromProvider,
	"elasticstack_fleet_integration":                   config.IdentifierFromProvider,
	"elasticstack_fleet_integration_policy":            config.IdentifierFromProvider,
	"elasticstack_fleet_output":                        config.IdentifierFromProvider,
	"elasticstack_fleet_server_host":                   config.IdentifierFromProvider,
	"elasticstack_kibana_action_connector":             config.IdentifierFromProvider,
	"elasticstack_kibana_alerting_rule":                config.IdentifierFromProvider,
	// "elasticstack_kibana_data_view":                    config.IdentifierFromProvider,
	// "elasticstack_kibana_import_saved_objects":         config.IdentifierFromProvider,
	"elasticstack_kibana_security_role": config.IdentifierFromProvider,
	"elasticstack_kibana_slo":           config.IdentifierFromProvider,
	"elasticstack_kibana_space":         config.IdentifierFromProvider,

	// Name is a parameter and it is also used to import the resource.
	"github_repository": config.NameAsIdentifier,
	// The import ID consists of several parameters. We'll use branch name as
	// the external name.
	"github_branch": config.TemplatedStringAsIdentifier("branch", "{{ .parameters.repository }}:{{ .external_name }}:{{ .parameters.source_branch }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
