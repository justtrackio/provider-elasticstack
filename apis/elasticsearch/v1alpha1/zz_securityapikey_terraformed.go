/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this SecurityAPIKey
func (mg *SecurityAPIKey) GetTerraformResourceType() string {
	return "elasticstack_elasticsearch_security_api_key"
}

// GetConnectionDetailsMapping for this SecurityAPIKey
func (tr *SecurityAPIKey) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"api_key": "status.atProvider.apiKey", "elasticsearch_connection[*].api_key": "spec.forProvider.elasticsearchConnection[*].apiKeySecretRef", "elasticsearch_connection[*].bearer_token": "spec.forProvider.elasticsearchConnection[*].bearerTokenSecretRef", "elasticsearch_connection[*].endpoints[*]": "spec.forProvider.elasticsearchConnection[*].endpointsSecretRef[*]", "elasticsearch_connection[*].es_client_authentication": "spec.forProvider.elasticsearchConnection[*].esClientAuthenticationSecretRef", "elasticsearch_connection[*].key_data": "spec.forProvider.elasticsearchConnection[*].keyDataSecretRef", "elasticsearch_connection[*].password": "spec.forProvider.elasticsearchConnection[*].passwordSecretRef", "encoded": "status.atProvider.encoded"}
}

// GetObservation of this SecurityAPIKey
func (tr *SecurityAPIKey) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SecurityAPIKey
func (tr *SecurityAPIKey) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SecurityAPIKey
func (tr *SecurityAPIKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SecurityAPIKey
func (tr *SecurityAPIKey) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SecurityAPIKey
func (tr *SecurityAPIKey) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SecurityAPIKey
func (tr *SecurityAPIKey) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this SecurityAPIKey
func (tr *SecurityAPIKey) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this SecurityAPIKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SecurityAPIKey) LateInitialize(attrs []byte) (bool, error) {
	params := &SecurityAPIKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SecurityAPIKey) GetTerraformSchemaVersion() int {
	return 0
}
