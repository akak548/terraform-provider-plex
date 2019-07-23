package plex

import (
    "github.com/hashicorp/terraform/helper/schema"
    plexclient "github.com/akak548/go-plex-client"
    "fmt"
)

func findFriend(e string, fs []plexclient.Friends) (plexclient.Friends, error){

    for _, f := range fs {
        
    }
}

func resourceFriend() *schema.Resource {
    return &schema.Resource{
        Create: resourceFriendCreate,
        Read: resourceFriendRead,
        Update: resourceFriendUpdate,
        Delete: resourceFriendDelete,
        Schema: map[string]*schema.Schema {
            "email_address": &schema.Schema {
                Type: schema.TypeString,
                Required: true,
            },
            "allow_sync": &schema.Schema {
                Type: schema.TypeBool,
                Optional: true,
                Default: false,
                Description: "Allow users are able to make use of Mobile Sync to sync content to their mobile devices",
            },
            "allow_cameraupload": &schema.Schema {
                Type: schema.TypeBool,
                Optional: true,
                Default: false,
                Description: "Allow users are able to upload media to server",
            },
        },
    }
}

func resourceFriendCreate(d *schema.ResourceData, m interface{}) error {
    email_address := d.Get("email_address").(string)
    d.SetId(email_address)
    return resourceFriendRead(d, m)
}

func resourceFriendRead(d *schema.ResourceData, m interface{}) error {
    client := m.(*plexclient.Plex)

    friends, err := client.GetFriends()

    if err != nil {
        return fmt.Errorf("Error Getting Plex friends list for %s: Error: %s", d.Id(), err)
    }
    friends
    return nil 
}

func resourceFriendUpdate(d *schema.ResourceData, m interface{}) error {
    return resourceFriendRead(d, m)
}

func resourceFriendDelete(d *schema.ResourceData, m interface{}) error {
    return nil 
}
