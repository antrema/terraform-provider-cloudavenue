// Package provider provides the CloudAvenue Terraform Provider.
package provider

import (
	"context"
	"errors"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	casdk "github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	clientnetbackup "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/netbackup"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/client"
)

const VCDVersion = "37.2"

// Ensure the implementation satisfies the expected interfaces.
var _ provider.Provider = &cloudavenueProvider{}

// cloudavenueProvider is the provider implementation.
type cloudavenueProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &cloudavenueProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *cloudavenueProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudavenue"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *cloudavenueProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerSchema(ctx)
}

func (p *cloudavenueProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config cloudavenueProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cloudAvenue := client.CloudAvenue{
		CloudAvenueVersion: p.version,
		// Legacy SDK Cloudavenue
		URL: func() string {
			url := findValue(config.URL, "CLOUDAVENUE_URL")
			if url == "" {
				url = "https://console1.cloudavenue.orange-business.com"
			}
			return url
		}(),
		User:             findValue(config.User, "CLOUDAVENUE_USER"),
		Password:         findValue(config.Password, "CLOUDAVENUE_PASSWORD"),
		Org:              findValue(config.Org, "CLOUDAVENUE_ORG"),
		VDC:              findValue(config.VDC, "CLOUDAVENUE_VDC"),
		TerraformVersion: req.TerraformVersion,
		VCDVersion:       VCDVersion,

		// This is a new SDK Cloudavenue
		CAVSDKOpts: &casdk.ClientOpts{
			Netbackup: &clientnetbackup.Opts{
				Endpoint: findValue(config.NetBackupURL, "NETBACKUP_URL"),
				Username: findValue(config.NetBackupUser, "NETBACKUP_USER"),
				Password: findValue(config.NetBackupPassword, "NETBACKUP_PASSWORD"),
			},
			CloudAvenue: &clientcloudavenue.Opts{
				Endpoint: func() string {
					url := findValue(config.URL, "CLOUDAVENUE_URL")
					if url == "" {
						url = "https://console1.cloudavenue.orange-business.com"
					}
					return url
				}(),
				Username:   findValue(config.User, "CLOUDAVENUE_USER"),
				Password:   findValue(config.Password, "CLOUDAVENUE_PASSWORD"),
				Org:        findValue(config.Org, "CLOUDAVENUE_ORG"),
				VDC:        findValue(config.VDC, "CLOUDAVENUE_VDC"),
				VCDVersion: VCDVersion,
			},
		},
	}

	const valueNotEmpty = "If either is already set, ensure the value is not empty."

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if cloudAvenue.User == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("user"),
			"Missing Cloud Avenue API User",
			"The provider cannot create the Cloud Avenue API client as there is a missing or empty value for the Cloud Avenue API user. "+
				"Set the host value in the configuration or use the CLOUDAVENUE_USER environment variable. "+
				valueNotEmpty,
		)
	}
	if cloudAvenue.Password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Cloud Avenue API Password",
			"The provider cannot create the Cloud Avenue API client as there is a missing or empty value for the Cloud Avenue API password. "+
				"Set the host value in the configuration or use the CLOUDAVENUE_PASSWORD environment variable. "+
				valueNotEmpty,
		)
	}
	if cloudAvenue.Org == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("org"),
			"Missing Cloud Avenue API Org",
			"The provider cannot create the Cloud Avenue API client as there is a missing or empty value for the Cloud Avenue API org. "+
				"Set the host value in the configuration or use the CLOUDAVENUE_ORG environment variable. "+
				valueNotEmpty,
		)
	}

	if cloudAvenue.CAVSDKOpts.Netbackup.Username == "" && cloudAvenue.CAVSDKOpts.Netbackup.Password != "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("netbackup_user"),
			"Missing NetBackup API User",
			"The provider cannot create the NetBackup API client as there is a missing or empty value for the NetBackup API user. "+
				"Set the host value in the configuration or use the NETBACKUP_USER environment variable. "+
				valueNotEmpty,
		)
	}
	if cloudAvenue.CAVSDKOpts.Netbackup.Password == "" && cloudAvenue.CAVSDKOpts.Netbackup.Username != "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("netbackup_password"),
			"Missing NetBackup API Password",
			"The provider cannot create the NetBackup API client as there is a missing or empty value for the NetBackup API password. "+
				"Set the host value in the configuration or use the NETBACKUP_PASSWORD environment variable. "+
				valueNotEmpty,
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	const (
		summaryErrorAPICAV = "Unable to Create Cloud Avenue API Client"
		summaryErrorVCD    = "Unable to Create VMWare VCD Client"
	)

	cA, err := cloudAvenue.New()
	if err != nil {
		switch {
		case errors.Is(err, client.ErrAuthFailed):
			resp.Diagnostics.AddError(summaryErrorAPICAV, err.Error())
			return
		case errors.Is(err, client.ErrTokenEmpty):
			resp.Diagnostics.AddError(summaryErrorAPICAV, "Cloud Avenue Client Error: empty token")
			return
		case errors.Is(err, client.ErrConfigureVmware):
			resp.Diagnostics.AddError(summaryErrorVCD, "VMWare VCD Client Error: "+err.Error())
			return
		case errors.Is(err, client.ErrVCDVersionEmpty):
			resp.Diagnostics.AddError(summaryErrorVCD, "VMWare VCD version is empty")
			return
		case errors.Is(err, client.ErrConfigureNetBackup):
			resp.Diagnostics.AddError("Unable to Configure NetBackup Client", err.Error())
			return
		default:
			resp.Diagnostics.AddError(summaryErrorAPICAV, "unknown error: "+err.Error())
			return
		}
	}

	// Make the CloudAvenue client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = cA
	resp.ResourceData = cA
}

func findValue(tfValue basetypes.StringValue, envName string) string {
	if tfValue.IsNull() {
		return os.Getenv(envName)
	}
	return tfValue.ValueString()
}
