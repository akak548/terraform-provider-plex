package main

import (
	"github.com/akak548/terraform-provider-plex/plex"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return plex.Provider()
		},
	})
}
