package main

import (
	"github.com/GSLabDev/terraform-provider-nsx/nsx"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return nsx.Provider()
		},
	})
}
