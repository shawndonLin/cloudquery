package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudresourcemanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/container"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/storage"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PluginAutoGeneratedTables() schema.Tables {
	tables := []*schema.Table{
		cloudresourcemanager.Organizations(),
		compute.Addresses(),
		compute.Autoscalers(),
		compute.BackendServices(),
		compute.DiskTypes(),
		compute.Disks(),
		compute.ExternalVpnGateways(),
		compute.ForwardingRules(),
		compute.Instances(),
		compute.SslCertificates(),
		compute.Subnetworks(),
		compute.TargetHttpProxies(),
		compute.UrlMaps(),
		compute.VpnGateways(),
		compute.VpnTunnels(),
		compute.TargetVpnGateways(),
		compute.InstanceGroups(),
		compute.Images(),
		compute.Firewalls(),
		compute.Networks(),
		compute.SslPolicies(),
		compute.Interconnects(),
		compute.InterconnectAttachments(),
		compute.InterconnectLocations(),
		compute.InterconnectRemoteLocations(),
		compute.TargetSslProxies(),
		compute.Projects(),
		compute.Routers(),
		compute.Routes(),
		compute.Zones(),
		container.Clusters(),
		sql.Instances(),
		storage.Buckets(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}
