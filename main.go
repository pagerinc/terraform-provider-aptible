package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pagerinc/terraform-provider-aptible/aptible"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aptible.Provider,
	})
}
