package plex

import (
	"fmt"
	plexclient "github.com/akak548/go-plex-client"
	"github.com/hashicorp/terraform/helper/schema"
)

type friendNotFound struct {
	Id string
}

func (f *friendNotFound) Error() string {
	return fmt.Sprintf("Unable to find %v", f.Id)
}

func findFriend(f string, fs *[]plexclient.Friends) (plexclient.Friends, error) {
	return plexclient.Friends{}, nil
}

func resourceFriend() *schema.Resource {
	return &schema.Resource{
		Create: resourceFriendCreate,
		Read:   resourceFriendRead,
		Update: resourceFriendUpdate,
		Delete: resourceFriendDelete,
		Schema: map[string]*schema.Schema{
			"email_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"allow_sync": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Allow users are able to make use of Mobile Sync to sync content to their mobile devices",
			},
			"allow_cameraupload": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Allow users are able to upload media to server",
			},
			"allow_channels": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Allow users are able to make use of Mobile Sync to sync content to their mobile devices",
			},
			"machine_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				Description: "Allow user access to machine_id",
			},
			"library_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    false,
				Description: "Allow user access to plex libraries",
			},
			"restriction_labels": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    false,
				Description: "Allow user access to content of plex libraries",
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

	friend, err := findFriend(d.Id(), &friends)
	if err != nil {
		return fmt.Errorf("Error finding friend %s: Error: %s", d.Id(), err)
	}

	d.Set("allow_sync", friend.AllowSync)
	d.Set("allow_cameraupload", friend.AllowCameraUpload)
	return nil
}

func resourceFriendUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFriendRead(d, m)
}

func resourceFriendDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
