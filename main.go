package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
    "github.com/akak548/terraform-provider-plex/plex"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return plex.Provider()
        },
})
}
