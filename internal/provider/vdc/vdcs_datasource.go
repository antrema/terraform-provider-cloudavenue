package vdc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/client"
)

var (
	_ datasource.DataSource              = &vdcsDataSource{}
	_ datasource.DataSourceWithConfigure = &vdcsDataSource{}
)

func NewVdcsDataSource() datasource.DataSource {
	return &vdcsDataSource{}
}

type vdcsDataSource struct {
	client *client.CloudAvenue
}

type vdcsDataSourceModel struct {
	ID   types.String `tfsdk:"id"`
	Vdcs []vdc        `tfsdk:"vdcs"`
}

type vdc struct {
	VDCName types.String `tfsdk:"vdc_name"`
	VDCUuid types.String `tfsdk:"vdc_uuid"`
}

func (d *vdcsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vdcs"
}

func (d *vdcsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "List all VDC inside an Organization.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vdcs": schema.ListNestedAttribute{
				MarkdownDescription: "VDC list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vdc_name": schema.StringAttribute{
							MarkdownDescription: "VDC name.",
							Computed:            true,
						},
						"vdc_uuid": schema.StringAttribute{
							MarkdownDescription: "VDC UUID.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *vdcsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.CloudAvenue)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.CloudAvenue, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *vdcsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vdcsDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	vdcs, _, err := d.client.APIClient.VDCApi.GetOrgVdcs(d.client.Auth)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vdcs detail, got error: %s", err))
		return
	}

	data = vdcsDataSourceModel{}

	for _, v := range vdcs {
		data.Vdcs = append(data.Vdcs, vdc{
			VDCName: types.StringValue(v.VdcName),
			VDCUuid: types.StringValue(v.VdcUuid),
		})
	}

	data.ID = types.StringValue("frangipane")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
