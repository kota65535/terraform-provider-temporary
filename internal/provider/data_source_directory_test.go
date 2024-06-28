package provider

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
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
					resource.TestCheckResourceAttrWith("data.temporary_directory.main", "id", func(s string) error {
						info, err := os.Stat(s)
						require.NoError(t, err)
						if !info.IsDir() {
							return fmt.Errorf("'%s' should be a directory", s)
						}
						if info.Name() != "main" {
							return fmt.Errorf("directory name should be 'main', got '%s'", info.Name())
						}
						return nil
					}),
				),
			},
			{
				Config: testAccDataSourceTemporaryDirectoryWithBase(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith("data.temporary_directory.main", "id", func(s string) error {
						info, err := os.Stat(s)
						require.NoError(t, err)
						if !info.IsDir() {
							return fmt.Errorf("'%s' should be a directory", s)
						}
						expectedId := ".terraform/tmp/main"
						if s != expectedId {
							return fmt.Errorf("id should be '%s', got '%s'", expectedId, s)
						}
						return nil
					}),
				),
			},
		},
	})
}

func testAccDataSourceTemporaryDirectory() string {
	return fmt.Sprintf(`
		data "temporary_directory" "main" {
			name = "main"
		}
	`)
}

func testAccDataSourceTemporaryDirectoryWithBase() string {
	return fmt.Sprintf(`
    provider "temporary" {
		  base = "${path.root}/.terraform/tmp"
    }

		data "temporary_directory" "main" {
			name = "main"
		}
	`)
}
