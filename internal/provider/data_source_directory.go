package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"path"
)

func dataSourceTemporaryDirectory() *schema.Resource {
	return &schema.Resource{
		Description: "Create a temporary directory. The directory is deleted every `terraform apply`, `plan` and `refresh`.",
		ReadContext: dataSourceTemporaryDirectoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path of the temporary directory created.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Directory name.",
			},
		},
	}
}

func dataSourceTemporaryDirectoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	tmp := m.(*Temporary)

	// Check if path is inside the base dir
	p := path.Join(tmp.BaseDir, name)
	tflog.Info(ctx, "creating temporary directory", map[string]any{"path": p})
	contains, err := ContainsFilePath(tmp.BaseDir, p)
	if err != nil {
		return diag.FromErr(err)
	}
	if !contains {
		return diag.FromErr(fmt.Errorf("cannot create a temporary directory '%s', which is outside the base directory", p))
	}

	// Clean up the directory if exists
	s, err := os.Stat(p)
	if err == nil {
		if !s.IsDir() {
			return diag.FromErr(fmt.Errorf("a non-directory already exists at '%s'", p))
		}
		tflog.Info(ctx, "deleting existing directory", map[string]any{"path": p})
		err = os.RemoveAll(p)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	err = os.MkdirAll(p, 0755)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(p)

	return nil
}
