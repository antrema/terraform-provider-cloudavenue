package testsacc

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/pkg/uuid"
)

//go:generate tf-doc-extractor -filename $GOFILE -example-dir ../../../examples -test
const testAccIsolatedNetworkResourceConfig = `
resource "cloudavenue_vapp" "example" {
	name        = "MyVapp"
	description = "This is an example vApp"
  }

resource "cloudavenue_vapp_isolated_network" "example" {
	name       = "MyVappNet"
	vapp_name  = cloudavenue_vapp.example.name
	gateway    = "192.168.10.1"
	netmask	   = "255.255.255.0"
	dns1       = "192.168.10.1"
	dns2       = "192.168.10.3"
	dns_suffix = "myvapp.biz"
	guest_vlan_allowed = true
	retain_ip_mac_enabled = true
  
	static_ip_pool = [{
	  start_address = "192.168.10.51"
	  end_address   = "192.168.10.101"
	  },
	  {
		start_address = "192.168.10.10"
		end_address   = "192.168.10.30"
	}]
}
`

func TestAccIsolatedNetworkResource(t *testing.T) {
	const resourceName = "cloudavenue_vapp_isolated_network.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Apply test
				Config: testAccIsolatedNetworkResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "id", uuid.TestIsType(uuid.Network)),
					resource.TestCheckResourceAttr(resourceName, "vdc", os.Getenv("CLOUDAVENUE_VDC")),
					resource.TestCheckResourceAttr(resourceName, "name", "MyVappNet"),
					resource.TestCheckResourceAttr(resourceName, "vapp_name", "MyVapp"),
					resource.TestCheckResourceAttr(resourceName, "gateway", "192.168.10.1"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr(resourceName, "dns1", "192.168.10.1"),
					resource.TestCheckResourceAttr(resourceName, "dns2", "192.168.10.3"),
					resource.TestCheckResourceAttr(resourceName, "dns_suffix", "myvapp.biz"),
					resource.TestCheckResourceAttr(resourceName, "guest_vlan_allowed", "true"),
					resource.TestCheckResourceAttr(resourceName, "retain_ip_mac_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "static_ip_pool.0.start_address", "192.168.10.51"),
					resource.TestCheckResourceAttr(resourceName, "static_ip_pool.0.end_address", "192.168.10.101"),
					resource.TestCheckResourceAttr(resourceName, "static_ip_pool.1.start_address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "static_ip_pool.1.end_address", "192.168.10.30"),
				),
			},
			// Uncomment if you want to test update or delete this block
			// Update don't work at the moment : https://github.com/vmware/go-vcloud-director/issues/554
			// {
			// 	// Update test
			// 	Config: testAccIsolatedNetworkResourceConfigUpdate,
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttrWith(resourceName, "id", uuid.TestIsType(uuid.Network)),
			// 		resource.TestCheckResourceAttr(resourceName, "vdc", "MyVDC"),
			// 		resource.TestCheckResourceAttr(resourceName, "name", "MyVappNet"),
			// 		resource.TestCheckResourceAttr(resourceName, "vapp_name", "MyVapp"),
			// 		resource.TestCheckResourceAttr(resourceName, "gateway", "192.168.10.1"),
			// 		resource.TestCheckResourceAttr(resourceName, "netmask", "255.255.255.0"),
			// 		resource.TestCheckResourceAttr(resourceName, "dns1", "192.168.10.1"),
			// 		resource.TestCheckNoResourceAttr(resourceName, "dns2"),
			// 		resource.TestCheckResourceAttr(resourceName, "dns_suffix", "myvapp.biz"),
			// 		resource.TestCheckResourceAttr(resourceName, "guest_vlan_allowed", "false"),
			// 		resource.TestCheckResourceAttr(resourceName, "retain_ip_mac_enabled", "false"),
			// 	),
			// },
			// ImportruetState testing
			{
				// Import test with vdc
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "MyVDC.MyVapp.MyVappNet",
			},
			{
				// Import test without vdc
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "MyVapp.MyVappNet",
			},
		},
	})
}
