package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTemporaryDirectory(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemporaryDirectory(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.temporary_directory.main", "id", ".terraform/tmp/foo/bar"),
				),
			},
		},
	})
}

func testAccDataSourceTemporaryDirectory() string {
	return fmt.Sprintf(`
  provider "temporary" {
		base = "${path.root}/.terraform/tmp"
  }

	data "temporary_directory" "main" {
		name = "foo/bar"
	}
	`)
}
