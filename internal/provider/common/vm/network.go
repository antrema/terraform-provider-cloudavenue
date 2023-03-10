package vm

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	govcdtypes "github.com/vmware/go-vcloud-director/v2/types/v56"

	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/helpers/boolpm"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/helpers/stringpm"
)

type Networks []Network

type Network struct {
	Type             types.String `tfsdk:"type"`
	IPAllocationMode types.String `tfsdk:"ip_allocation_mode"`
	Name             types.String `tfsdk:"name"`
	IP               types.String `tfsdk:"ip"`
	IsPrimary        types.Bool   `tfsdk:"is_primary"`
	Mac              types.String `tfsdk:"mac"`
	AdapterType      types.String `tfsdk:"adapter_type"`
	Connected        types.Bool   `tfsdk:"connected"`
}

// NetworkAttrType is the type of the network attribute.
func NetworkAttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"type":               types.StringType,
		"ip_allocation_mode": types.StringType,
		"name":               types.StringType,
		"ip":                 types.StringType,
		"is_primary":         types.BoolType,
		"mac":                types.StringType,
		"adapter_type":       types.StringType,
		"connected":          types.BoolType,
	}
}

// ObjectType returns the type of the network object.
func (n *Network) ObjectType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: NetworkAttrType(),
	}
}

// ObjectType returns the type of the networks object.
func (n *Networks) ObjectType() types.ObjectType {
	x := Network{}
	return x.ObjectType()
}

// ToPlan converts a Network to a plan.
func (n *Networks) ToPlan() (basetypes.ListValue, diag.Diagnostics) {
	if n == nil || len(*n) == 0 {
		return types.ListNull(n.ObjectType()), diag.Diagnostics{}
	}

	return types.ListValueFrom(context.Background(), n.ObjectType(), n)
}

// NetworksFromPlan converts a plan to a Networks struct.
func NetworksFromPlan(x basetypes.ListValue) (networks *Networks, err error) {
	if x.IsNull() {
		return nil, errors.New("the ListValue is null")
	}

	var net *Networks

	diag := x.ElementsAs(context.Background(), net, false)
	if diag.HasError() {
		return nil, errors.New(diag[0].Detail())
	}

	return net, nil
}

// NetworkSchema returns the schema for the network
func NetworkSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"type": schema.StringAttribute{
			MarkdownDescription: "Network type to use: `vapp`, `org` or `none`. Use `vapp` for vApp network, `org` to attach Org VDC network. `none` for empty NIC.",
			Required:            true,
			// TODO : Add validator
		},
		"ip_allocation_mode": schema.StringAttribute{
			MarkdownDescription: "IP allocation mode: `DHCP`, `POOL`, `MANUAL` or `NONE`. ",
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("DHCP", "POOL", "MANUAL", "NONE"),
			},
			PlanModifiers: []planmodifier.String{
				stringpm.SetDefaultEmptyString(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name of the network this VM should connect to. Always required except for `type` `NONE`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringpm.SetDefault("none"),
			},
			// Not force new because it can be changed in-place
			// TODO : Add validator
		},
		"ip": schema.StringAttribute{
			MarkdownDescription: "IP of the VM. Settings depend on `ip_allocation_mode`. Omitted or empty for DHCP, POOL, NONE. Required for MANUAL",
			Optional:            true,
			Computed:            true,
			// TODO : Add validator
		},
		"is_primary": schema.BoolAttribute{
			MarkdownDescription: "Set to true if network interface should be primary. First network card in the list will be primary by default",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolpm.SetDefault(false),
			},
		},
		"mac": schema.StringAttribute{
			MarkdownDescription: "MAC address of the VM. Optional and autogenerated by default.",
			Optional:            true,
			Computed:            true,
			// TODO : Add validator
		},
		"adapter_type": schema.StringAttribute{
			MarkdownDescription: "The type of vNic to create on this interface. One of: `VMXNET3`, `E1000`, `E1000E`, `PCNet32`, `SRIOVETHERNETCARD`.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("VMXNET3", "E1000", "E1000E", "PCNet32", "SRIOVETHERNETCARD"),
			},
		},
		"connected": schema.BoolAttribute{
			MarkdownDescription: "Set to true if network should be connected or false otherwise. Default is `true`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolpm.SetDefault(true),
			},
		},
	}
}

// NetworksRead returns network configuration for saving into statefile
func NetworksRead(vm *govcd.VM) (*Networks, error) {
	vapp, err := vm.GetParentVApp()
	if err != nil {
		return nil, fmt.Errorf("error getting vApp: %w", err)
	}

	// Determine type for all networks in vApp
	vAppNetworkConfig, err := vapp.GetNetworkConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting vApp networks: %w", err)
	}
	// If vApp network is "isolated" and has no ParentNetwork - it is a vApp network.
	// https://code.vmware.com/apis/72/vcloud/doc/doc/types/NetworkConfigurationType.html
	vAppNetworkTypes := make(map[string]string, 0)
	for _, netConfig := range vAppNetworkConfig.NetworkConfig {
		switch {
		case netConfig.NetworkName == govcdtypes.NoneNetwork:
			vAppNetworkTypes[netConfig.NetworkName] = govcdtypes.NoneNetwork
		case govcd.IsVappNetwork(netConfig.Configuration):
			vAppNetworkTypes[netConfig.NetworkName] = "vapp"
		default:
			vAppNetworkTypes[netConfig.NetworkName] = "org"
		}
	}

	nets := make(Networks, 0)
	// Sort NIC cards by their virtual slot numbers as the API returns them in random order
	sort.SliceStable(vm.VM.NetworkConnectionSection.NetworkConnection, func(i, j int) bool {
		return vm.VM.NetworkConnectionSection.NetworkConnection[i].NetworkConnectionIndex <
			vm.VM.NetworkConnectionSection.NetworkConnection[j].NetworkConnectionIndex
	})

	for _, vmNet := range vm.VM.NetworkConnectionSection.NetworkConnection {
		singleNIC := Network{
			IPAllocationMode: types.StringValue(vmNet.IPAddressAllocationMode),
			IP:               types.StringValue(vmNet.IPAddress),
			Mac:              types.StringValue(vmNet.MACAddress),
			AdapterType:      types.StringValue(vmNet.NetworkAdapterType),
			Connected:        types.BoolValue(vmNet.IsConnected),
			IsPrimary:        types.BoolValue(false),
			Type:             types.StringValue(vAppNetworkTypes[vmNet.Network]),
		}

		if vmNet.Network != govcdtypes.NoneNetwork {
			singleNIC.Name = types.StringValue(vmNet.Network)
		}

		if vmNet.NetworkConnectionIndex == vm.VM.NetworkConnectionSection.PrimaryNetworkConnectionIndex {
			singleNIC.IsPrimary = types.BoolValue(true)
		}

		nets = append(nets, singleNIC)
	}

	return &nets, nil
}
