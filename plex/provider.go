package plex

import (
	plexclient "github.com/akak548/go-plex-client"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PLEX_API_TOKEN", nil),
				Description: "API Token for Plex Media Server",
			},
			"api_address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PLEX_API_ADDRESS", nil),
				Description: "Address for Plex Media Server",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"plex_friend": resourceFriend(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	return plexclient.New(
		data.Get("api_address").(string),
		data.Get("api_token").(string),
	)
}
