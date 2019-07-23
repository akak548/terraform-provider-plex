package plex

import (
        "github.com/hashicorp/terraform/helper/schema"
        plexclient "github.com/akak548/go-plex-client"
)

func Provider() *schema.Provider {
        return &schema.Provider{
            Schema: map[string]*schema.Schema{
                "api_token":{
                    Type:       schema.TypeString,
                    Required:    true,
                    DefaultFunc: schema.EnvDefaultFunc("PLEX_API_TOKEN", nil),
                },
                "api_address":{
                    Type:       schema.TypeString,
                    Required:    true,
                    DefaultFunc: schema.EnvDefaultFunc("PLEX_API_TOKEN", nil),
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
