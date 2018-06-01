package nsx

import (
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider for NSX
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "NSX Server Port",
				DefaultFunc: schema.EnvDefaultFunc("NSX_SERVER_PORT", nil),
			},
			"nsx_username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "NSX username",
				DefaultFunc: schema.EnvDefaultFunc("NSX_SERVER_USERNAME", nil),
			},
			"nsx_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "NSX Password for provided user_name",
				DefaultFunc: schema.EnvDefaultFunc("NSX_SERVER_PASSWORD", nil),
			},
			"nsx_server_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "NSX Server IP Address",
				DefaultFunc: schema.EnvDefaultFunc("NSX_SERVER_IP", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"nsx_add_virtual_machine_security_group": resourceNsxAddVirtualMachineSecurityGroup(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config, err := NsxConfig(d)
	if err != nil {
		log.Println("[ERROR] no connection was established with nsx server.")
		os.Exit(1)
	}
	return config, nil
}
