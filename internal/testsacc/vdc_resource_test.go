package testsacc

import (
	"context"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/internal/helpers/testsacc"
	"github.com/orange-cloudavenue/terraform-provider-cloudavenue/pkg/uuid"
)

var _ testsacc.TestACC = &VDCResource{}

const (
	VDCResourceName = testsacc.ResourceName("cloudavenue_vdc")
)

type VDCResource struct{}

func NewVDCResourceTest() testsacc.TestACC {
	return &VDCResource{}
}

// GetResourceName returns the name of the resource.
func (r *VDCResource) GetResourceName() string {
	return VDCResourceName.String()
}

func (r *VDCResource) DependenciesConfig() (resp testsacc.DependenciesConfigResponse) {
	return
}

func (r *VDCResource) Tests(ctx context.Context) map[testsacc.TestName]func(ctx context.Context, resourceName string) testsacc.Test {
	return map[testsacc.TestName]func(ctx context.Context, resourceName string) testsacc.Test{
		// * Test One (example)
		"example": func(_ context.Context, resourceName string) testsacc.Test {
			return testsacc.Test{
				CommonChecks: []resource.TestCheckFunc{
					resource.TestCheckResourceAttrWith(resourceName, "id", uuid.TestIsType(uuid.VDC)),
					resource.TestCheckResourceAttr(resourceName, "billing_model", "PAYG"),
					resource.TestCheckResourceAttr(resourceName, "disponibility_class", "ONE-ROOM"),
					resource.TestCheckResourceAttr(resourceName, "service_class", "STD"),
					resource.TestCheckResourceAttr(resourceName, "storage_billing_model", "PAYG"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.class", "gold"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.default", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.limit", "500"),
				},
				// ! Create testing
				Create: testsacc.TFConfig{
					TFConfig: testsacc.GenerateFromTemplate(resourceName, `
					resource "cloudavenue_vdc" "example" {
						name                  = {{ generate . "name" }}
						description           = {{ generate . "description" "longString"}}
						cpu_allocated         = 22000
						memory_allocated      = 30
						cpu_speed_in_mhz      = 2200
						billing_model         = "PAYG"
						disponibility_class   = "ONE-ROOM"
						service_class         = "STD"
						storage_billing_model = "PAYG"
					  
						storage_profiles = [{
						  class   = "gold"
						  default = true
						  limit   = 500
						}]
					  
					}`),
					Checks: []resource.TestCheckFunc{
						resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
						resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
						resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22000"),
						resource.TestCheckResourceAttr(resourceName, "memory_allocated", "30"),
						resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2200"),
					},
				},
				// ! Updates testing
				Updates: []testsacc.TFConfig{
					// Update description
					{
						TFConfig: testsacc.GenerateFromTemplate(resourceName, `
						resource "cloudavenue_vdc" "example" {
							name                  = {{ get . "name" }}
							description           = {{ generate . "description" "longString"}}
							cpu_allocated         = 22000
							memory_allocated      = 30
							cpu_speed_in_mhz      = 2200
							billing_model         = "PAYG"
							disponibility_class   = "ONE-ROOM"
							service_class         = "STD"
							storage_billing_model = "PAYG"
						  
							storage_profiles = [{
							  class   = "gold"
							  default = true
							  limit   = 500
							}]
						  
						}`),
						Checks: []resource.TestCheckFunc{
							resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
							resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
							resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22000"),
							resource.TestCheckResourceAttr(resourceName, "memory_allocated", "30"),
							resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2200"),
						},
					},
					// Update cpu_allocated
					{
						TFConfig: testsacc.GenerateFromTemplate(resourceName, `
						resource "cloudavenue_vdc" "example" {
							name                  = {{ get . "name" }}
							description           = {{ get . "description" }}
							cpu_allocated         = 22500
							memory_allocated      = 30
							cpu_speed_in_mhz      = 2200
							billing_model         = "PAYG"
							disponibility_class   = "ONE-ROOM"
							service_class         = "STD"
							storage_billing_model = "PAYG"

							storage_profiles = [{
								class   = "gold"
								default = true
								limit   = 500
							  }]
						}`),
						Checks: []resource.TestCheckFunc{
							resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
							resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
							resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22500"),
							resource.TestCheckResourceAttr(resourceName, "memory_allocated", "30"),
							resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2200"),
						},
					},
					// Update memory_allocated
					{
						TFConfig: testsacc.GenerateFromTemplate(resourceName, `
						resource "cloudavenue_vdc" "example" {
							name                  = {{ get . "name" }}
							description           = {{ get . "description" }}
							cpu_allocated         = 22500
							memory_allocated      = 40
							cpu_speed_in_mhz      = 2200
							billing_model         = "PAYG"
							disponibility_class   = "ONE-ROOM"
							service_class         = "STD"
							storage_billing_model = "PAYG"

							storage_profiles = [{
								class   = "gold"
								default = true
								limit   = 500
							  }]
						}`),
						Checks: []resource.TestCheckFunc{
							resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
							resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
							resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22500"),
							resource.TestCheckResourceAttr(resourceName, "memory_allocated", "40"),
							resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2200"),
						},
					},
					// Update cpu_speed_in_mhz
					// NOTE : This generate resource replacement
					{
						TFConfig: testsacc.GenerateFromTemplate(resourceName, `
						resource "cloudavenue_vdc" "example" {
							name                  = {{ get . "name" }}
							description           = {{ get . "description" }}
							cpu_allocated         = 22000
							memory_allocated      = 30
							cpu_speed_in_mhz      = 2300
							billing_model         = "PAYG"
							disponibility_class   = "ONE-ROOM"
							service_class         = "STD"
							storage_billing_model = "PAYG"
						  
							storage_profiles = [{
							  class   = "gold"
							  default = true
							  limit   = 500
							}]
						  
						}`),
						TFAdvanced: testsacc.TFAdvanced{
							PlanOnly:           true,
							ExpectNonEmptyPlan: true,
							ExpectError:        regexp.MustCompile("You can change the cpu_speed_in_mhz attribute only if the billing_model is set to RESERVED."),
						},
					},
				},
				// ! Imports testing
				Imports: []testsacc.TFImport{
					{
						ImportStateID:     testsacc.GetValueFromTemplate(resourceName, "name"),
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			}
		},
		"example_reserved": func(_ context.Context, resourceName string) testsacc.Test {
			return testsacc.Test{
				CommonChecks: []resource.TestCheckFunc{
					resource.TestCheckResourceAttrWith(resourceName, "id", uuid.TestIsType(uuid.VDC)),
					resource.TestCheckResourceAttr(resourceName, "billing_model", "RESERVED"),
					resource.TestCheckResourceAttr(resourceName, "disponibility_class", "ONE-ROOM"),
					resource.TestCheckResourceAttr(resourceName, "service_class", "STD"),
					resource.TestCheckResourceAttr(resourceName, "storage_billing_model", "PAYG"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.class", "gold"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.default", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.limit", "500"),
				},
				// ! Create testing
				Create: testsacc.TFConfig{
					TFConfig: testsacc.GenerateFromTemplate(resourceName, `
					resource "cloudavenue_vdc" "example_reserved" {
						name                  = {{ generate . "name" }}
						description           = {{ generate . "description" "longString"}}
						cpu_allocated         = 22000
						memory_allocated      = 30
						cpu_speed_in_mhz      = 2200
						billing_model         = "RESERVED"
						disponibility_class   = "ONE-ROOM"
						service_class         = "STD"
						storage_billing_model = "PAYG"
					  
						storage_profiles = [{
						  class   = "gold"
						  default = true
						  limit   = 500
						}]
					  
					}`),
					Checks: []resource.TestCheckFunc{
						resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
						resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
						resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22000"),
						resource.TestCheckResourceAttr(resourceName, "memory_allocated", "30"),
						resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2200"),
					},
				},
				// ! Updates testing
				Updates: []testsacc.TFConfig{
					// Update cpu_speed_in_mhz
					{
						TFConfig: testsacc.GenerateFromTemplate(resourceName, `
						resource "cloudavenue_vdc" "example_reserved" {
							name                  = {{ get . "name" }}
							description           = {{ generate . "description" "longString"}}
							cpu_allocated         = 22000
							memory_allocated      = 30
							cpu_speed_in_mhz      = 2300
							billing_model         = "RESERVED"
							disponibility_class   = "ONE-ROOM"
							service_class         = "STD"
							storage_billing_model = "PAYG"
						  
							storage_profiles = [{
							  class   = "gold"
							  default = true
							  limit   = 500
							}]
						  
						}`),
						Checks: []resource.TestCheckFunc{
							resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
							resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
							resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22000"),
							resource.TestCheckResourceAttr(resourceName, "memory_allocated", "30"),
							resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2300"),
						},
					},
				},
			}
		},
		// This is used to test vdc_group resource
		"example_vdc_group_1": func(_ context.Context, resourceName string) testsacc.Test {
			return testsacc.Test{
				CommonChecks: []resource.TestCheckFunc{
					resource.TestCheckResourceAttrWith(resourceName, "id", uuid.TestIsType(uuid.VDC)),
					resource.TestCheckResourceAttr(resourceName, "billing_model", "PAYG"),
					resource.TestCheckResourceAttr(resourceName, "disponibility_class", "ONE-ROOM"),
					resource.TestCheckResourceAttr(resourceName, "service_class", "STD"),
					resource.TestCheckResourceAttr(resourceName, "storage_billing_model", "PAYG"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.class", "gold"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.default", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.limit", "500"),
				},
				// ! Create testing
				Create: testsacc.TFConfig{
					TFConfig: testsacc.GenerateFromTemplate(resourceName, `
					resource "cloudavenue_vdc" "example_vdc_group_1" {
						name                  = {{ generate . "name" }}
						description           = {{ generate . "description" "longString"}}
						cpu_allocated         = 22000
						memory_allocated      = 30
						cpu_speed_in_mhz      = 2200
						billing_model         = "PAYG"
						disponibility_class   = "ONE-ROOM"
						service_class         = "STD"
						storage_billing_model = "PAYG"
					  
						storage_profiles = [{
						  class   = "gold"
						  default = true
						  limit   = 500
						}]
					  
					}`),
					Checks: []resource.TestCheckFunc{
						resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
						resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
						resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22000"),
						resource.TestCheckResourceAttr(resourceName, "memory_allocated", "30"),
						resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2200"),
					},
				},
			}
		},
		"example_vdc_group_2": func(_ context.Context, resourceName string) testsacc.Test {
			return testsacc.Test{
				CommonChecks: []resource.TestCheckFunc{
					resource.TestCheckResourceAttrWith(resourceName, "id", uuid.TestIsType(uuid.VDC)),
					resource.TestCheckResourceAttr(resourceName, "billing_model", "PAYG"),
					resource.TestCheckResourceAttr(resourceName, "disponibility_class", "ONE-ROOM"),
					resource.TestCheckResourceAttr(resourceName, "service_class", "STD"),
					resource.TestCheckResourceAttr(resourceName, "storage_billing_model", "PAYG"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.class", "gold"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.default", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_profiles.0.limit", "500"),
				},
				// ! Create testing
				Create: testsacc.TFConfig{
					TFConfig: testsacc.GenerateFromTemplate(resourceName, `
					resource "cloudavenue_vdc" "example_vdc_group_2" {
						name                  = {{ generate . "name" }}
						description           = {{ generate . "description" "longString"}}
						cpu_allocated         = 22000
						memory_allocated      = 30
						cpu_speed_in_mhz      = 2200
						billing_model         = "PAYG"
						disponibility_class   = "ONE-ROOM"
						service_class         = "STD"
						storage_billing_model = "PAYG"
					  
						storage_profiles = [{
						  class   = "gold"
						  default = true
						  limit   = 500
						}]
					  
					}`),
					Checks: []resource.TestCheckFunc{
						resource.TestCheckResourceAttr(resourceName, "name", testsacc.GetValueFromTemplate(resourceName, "name")),
						resource.TestCheckResourceAttr(resourceName, "description", testsacc.GetValueFromTemplate(resourceName, "description")),
						resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "22000"),
						resource.TestCheckResourceAttr(resourceName, "memory_allocated", "30"),
						resource.TestCheckResourceAttr(resourceName, "cpu_speed_in_mhz", "2200"),
					},
				},
			}
		},
	}
}

func TestAccVDCResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps:                    testsacc.GenerateTests(&VDCResource{}),
	})
}
