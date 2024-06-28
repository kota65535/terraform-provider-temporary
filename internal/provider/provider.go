package provider

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path of the base directory where temporary directories/files are created. If not specified, OS-dependent temporary directory is used.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"temporary_directory": dataSourceTemporaryDirectory(),
		},
		ConfigureContextFunc: configure(),
	}
}

func configure() func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(cxt context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		baseDir := d.Get("base").(string)
		if baseDir == "" {
			baseDir = os.TempDir()
		}

		s, err := os.Stat(baseDir)
		if err == nil {
			if !s.IsDir() {
				return nil, diag.FromErr(fmt.Errorf("a non-directory already exists at '%s'", baseDir))
			}
		} else {
			err = os.MkdirAll(baseDir, 0755)
			if err != nil {
				return nil, diag.FromErr(err)
			}
		}

		tmp := NewTemporary(baseDir)

		return tmp, nil
	}
}
