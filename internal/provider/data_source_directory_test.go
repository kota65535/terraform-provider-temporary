package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceTemporaryDirectory(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemporaryDirectory(),
				Check: resource.ComposeTestCheckFunc(
					testAccFilesExists(".terraform-provider-temporary", ".terraform/tmp"),
					resource.TestCheckResourceAttr("data.temporary_directory.main", "id", ".terraform/tmp/main"),
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
		name = "main"
	}
	`)
}

func testAccFilesExists(filename string, dir string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, err := os.Stat(filepath.Join(dir, filename))
		if err != nil {
			return err
		}
		return nil
	}
}
