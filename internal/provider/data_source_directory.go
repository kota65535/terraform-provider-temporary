package provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"path"
)

func dataSourceTemporaryDirectory() *schema.Resource {
	return &schema.Resource{
		Description: "Create a temporary directory.",
		ReadContext: dataSourceTemporaryDirectoryRead,
		Schema: map[string]*schema.Schema{
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

	p := path.Join(tmp.BaseDir, name)

	s, err := os.Stat(p)
	if err != nil {
		err = os.Mkdir(p, 0777)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		if !s.IsDir() {
			err = errors.New(fmt.Sprintf("non-directory exists in %s", p))
			return diag.FromErr(err)
		}
	}

	d.SetId(p)

	return nil
}
