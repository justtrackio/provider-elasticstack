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

// GetTerraformResourceType returns Terraform resource type for this SecurityRole
func (mg *SecurityRole) GetTerraformResourceType() string {
	return "elasticstack_elasticsearch_security_role"
}

// GetConnectionDetailsMapping for this SecurityRole
func (tr *SecurityRole) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"elasticsearch_connection[*].api_key": "elasticsearchConnection[*].apiKeySecretRef", "elasticsearch_connection[*].bearer_token": "elasticsearchConnection[*].bearerTokenSecretRef", "elasticsearch_connection[*].endpoints[*]": "elasticsearchConnection[*].endpointsSecretRef[*]", "elasticsearch_connection[*].es_client_authentication": "elasticsearchConnection[*].esClientAuthenticationSecretRef", "elasticsearch_connection[*].key_data": "elasticsearchConnection[*].keyDataSecretRef", "elasticsearch_connection[*].password": "elasticsearchConnection[*].passwordSecretRef"}
}

// GetObservation of this SecurityRole
func (tr *SecurityRole) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SecurityRole
func (tr *SecurityRole) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SecurityRole
func (tr *SecurityRole) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SecurityRole
func (tr *SecurityRole) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SecurityRole
func (tr *SecurityRole) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SecurityRole
func (tr *SecurityRole) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this SecurityRole
func (tr *SecurityRole) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
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

// LateInitialize this SecurityRole using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SecurityRole) LateInitialize(attrs []byte) (bool, error) {
	params := &SecurityRoleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SecurityRole) GetTerraformSchemaVersion() int {
	return 0
}
