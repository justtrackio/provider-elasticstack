/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	clustersettings "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/clustersettings"
	componenttemplate "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/componenttemplate"
	datastream "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/datastream"
	enrichpolicy "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/enrichpolicy"
	index "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/index"
	indexlifecycle "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/indexlifecycle"
	indextemplate "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/indextemplate"
	ingestpipeline "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/ingestpipeline"
	logstashpipeline "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/logstashpipeline"
	script "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/script"
	securityapikey "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/securityapikey"
	securityrole "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/securityrole"
	securityrolemapping "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/securityrolemapping"
	securitysystemuser "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/securitysystemuser"
	securityuser "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/securityuser"
	snapshotlifecycle "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/snapshotlifecycle"
	snapshotrepository "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/snapshotrepository"
	transform "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/transform"
	watch "github.com/justtrackio/provider-elasticstack/internal/controller/elasticsearch/watch"
	agentpolicy "github.com/justtrackio/provider-elasticstack/internal/controller/fleet/agentpolicy"
	integration "github.com/justtrackio/provider-elasticstack/internal/controller/fleet/integration"
	integrationpolicy "github.com/justtrackio/provider-elasticstack/internal/controller/fleet/integrationpolicy"
	output "github.com/justtrackio/provider-elasticstack/internal/controller/fleet/output"
	serverhost "github.com/justtrackio/provider-elasticstack/internal/controller/fleet/serverhost"
	actionconnector "github.com/justtrackio/provider-elasticstack/internal/controller/kibana/actionconnector"
	alertingrule "github.com/justtrackio/provider-elasticstack/internal/controller/kibana/alertingrule"
	securityrolekibana "github.com/justtrackio/provider-elasticstack/internal/controller/kibana/securityrole"
	slo "github.com/justtrackio/provider-elasticstack/internal/controller/kibana/slo"
	space "github.com/justtrackio/provider-elasticstack/internal/controller/kibana/space"
	providerconfig "github.com/justtrackio/provider-elasticstack/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clustersettings.Setup,
		componenttemplate.Setup,
		datastream.Setup,
		enrichpolicy.Setup,
		index.Setup,
		indexlifecycle.Setup,
		indextemplate.Setup,
		ingestpipeline.Setup,
		logstashpipeline.Setup,
		script.Setup,
		securityapikey.Setup,
		securityrole.Setup,
		securityrolemapping.Setup,
		securitysystemuser.Setup,
		securityuser.Setup,
		snapshotlifecycle.Setup,
		snapshotrepository.Setup,
		transform.Setup,
		watch.Setup,
		agentpolicy.Setup,
		integration.Setup,
		integrationpolicy.Setup,
		output.Setup,
		serverhost.Setup,
		actionconnector.Setup,
		alertingrule.Setup,
		securityrolekibana.Setup,
		slo.Setup,
		space.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
