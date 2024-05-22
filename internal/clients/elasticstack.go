/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/justtrackio/provider-elasticstack/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal elasticstack credentials as JSON"
	elasticsearchEndpoints  = "endpoints"
	elasticsearchUsername   = "username"
	elasticsearchPassword   = "password"
	elasticsearchApiKey     = "api_key"
	kibanaEndpoints         = "endpoints"
	kibanaUsername          = "username"
	kibanaPassword          = "password"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}

		// Elasticsearch configuration
		elasticsearchConfig := buildElasticsearchConfig(creds)
		ps.Configuration["elasticsearch"] = elasticsearchConfig

		// Kibana configuration
		kibanaConfig := buildKibanaConfig(creds)
		ps.Configuration["kibana"] = kibanaConfig

		return ps, nil
	}
}

// buildElasticsearchConfig constructs the Elasticsearch configuration from credentials
func buildElasticsearchConfig(creds map[string]string) map[string]any {
	config := map[string]any{}
	if endpoints, ok := creds[elasticsearchEndpoints]; ok {
		config["endpoints"] = endpoints
	}
	if username, ok := creds[elasticsearchUsername]; ok {
		config["username"] = username
	}
	if password, ok := creds[elasticsearchPassword]; ok {
		config["password"] = password
	}
	if apiKey, ok := creds[elasticsearchApiKey]; ok {
		config["api_key"] = apiKey
	}
	return config
}

// buildKibanaConfig constructs the Kibana configuration from credentials
func buildKibanaConfig(creds map[string]string) map[string]any {
	config := map[string]any{}
	if endpoints, ok := creds[kibanaEndpoints]; ok {
		config["endpoints"] = endpoints
	}
	if username, ok := creds[kibanaUsername]; ok {
		config["username"] = username
	}
	if password, ok := creds[kibanaPassword]; ok {
		config["password"] = password
	}
	return config
}
