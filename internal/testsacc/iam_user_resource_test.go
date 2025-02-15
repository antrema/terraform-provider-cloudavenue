// package testsacc provides the acceptance tests for the provider.
// package testsacc provides the acceptance tests for the provider.
package testsacc

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const TestAccOrgUserResourceConfig = `
resource "cloudavenue_iam_user" "example" {
	name        = "example"
	role_name   = "Organization Administrator"
	password    = "Th!s1sSecur3P@ssword"
}
`

const testAccOrgUserResourceConfigFull = `
resource "cloudavenue_iam_user" "example" {
	name              = "example"
	role_name         = "Organization Administrator"
	password          = "Th!s1sSecur3P@ssword"
	enabled           = true
	email             = "foo@bar.com"
	telephone         = "1234567890"
	full_name         = "Test User"
	take_ownership    = true
	deployed_vm_quota = 10
	stored_vm_quota   = 5
 }
`

// Deprecated
// Used for testing the catalog ACL Resource/Datasource.
const testAccOrgUserResourceConfigForCatalogACL = `
resource "cloudavenue_iam_user" "example" {
	name        = "example"
	role_name   = "Organization Administrator"
	password    = "Th!s1sSecur3P@ssword"
}

resource "cloudavenue_iam_user" "example2" {
	name        = "example2"
	role_name   = "Organization Administrator"
	password    = "Th!s1sSecur3P@ssword"
}
`

func testsOrgUserResourceConfig(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "name", "example"),
		resource.TestCheckResourceAttr(resourceName, "role_name", "Organization Administrator"),
		resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
	)
}

func testsOrgUserResourceConfigFull(resourceName string, isDataSource bool) resource.TestCheckFunc {
	tests := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "name", "example"),
		resource.TestCheckResourceAttr(resourceName, "role_name", "Organization Administrator"),
		resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
		resource.TestCheckResourceAttr(resourceName, "email", "foo@bar.com"),
		resource.TestCheckResourceAttr(resourceName, "telephone", "1234567890"),
		resource.TestCheckResourceAttr(resourceName, "full_name", "Test User"),
		resource.TestCheckResourceAttr(resourceName, "deployed_vm_quota", "10"),
		resource.TestCheckResourceAttr(resourceName, "stored_vm_quota", "5"),
	}

	if !isDataSource {
		tests = append(tests, resource.TestCheckResourceAttr(resourceName, "take_ownership", "true"))
	}

	return resource.ComposeAggregateTestCheckFunc(
		tests...,
	)
}

func TestAccUserResource(t *testing.T) {
	resourceName := "cloudavenue_iam_user.example"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: TestAccOrgUserResourceConfig,
				Check:  testsOrgUserResourceConfig(resourceName),
			},
			{
				Config: testAccOrgUserResourceConfigFull,
				Check:  testsOrgUserResourceConfigFull(resourceName, false),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     "example",
				ImportStateVerify: true,
				// These fields can't be retrieved from user data
				ImportStateVerifyIgnore: []string{"take_ownership", "password"},
			},
		},
	})
}
