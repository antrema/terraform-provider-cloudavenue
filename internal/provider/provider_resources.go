package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/alb"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/backup"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/catalog"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/edgegw"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/iam"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/network"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/publicip"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/s3"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/vapp"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/vcda"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/vdc"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/provider/vm"
)

// Resources defines the resources implemented in the provider.
func (p *cloudavenueProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// * ALB
		alb.NewAlbPoolResource,

		// * EDGE GATEWAY
		edgegw.NewEdgeGatewayResource,
		edgegw.NewFirewallResource,
		edgegw.NewPortProfilesResource,
		edgegw.NewSecurityGroupResource,
		edgegw.NewIPSetResource,
		edgegw.NewDhcpForwardingResource,
		edgegw.NewStaticRouteResource,
		edgegw.NewNATRuleResource,
		edgegw.NewVPNIPSecResource,

		// * VDC
		vdc.NewVDCResource,
		vdc.NewACLResource,
		vdc.NewGroupResource,

		// * VCDA
		vcda.NewVCDAIPResource,

		// * PUBLICIP
		publicip.NewPublicIPResource,

		// * VAPP
		vapp.NewVappResource,
		vapp.NewOrgNetworkResource,
		vapp.NewIsolatedNetworkResource,
		vapp.NewACLResource,

		// * CATALOG
		catalog.NewCatalogResource,
		catalog.NewACLResource,

		// * IAM
		iam.NewIAMUserResource,
		iam.NewRoleResource,
		iam.NewTokenResource,

		// * VM
		vm.NewDiskResource,
		vm.NewVMResource,
		vm.NewInsertedMediaResource,
		vm.NewVMAffinityRuleResource,
		vm.NewSecurityTagResource,

		// * NETWORK
		network.NewNetworkRoutedResource,
		network.NewNetworkIsolatedResource,
		network.NewDhcpBindingResource,
		network.NewDhcpResource,

		// * BACKUP
		backup.NewBackupResource,

		// * S3
		s3.NewBucketVersioningConfigurationResource,
		s3.NewBucketResource,
		s3.NewBucketCorsConfigurationResource,
		s3.NewBucketLifecycleConfigurationResource,
		s3.NewBucketWebsiteConfigurationResource,
		s3.NewBucketACLResource,
		s3.NewCredentialResource,
		s3.NewBucketPolicyResource,
	}
}
