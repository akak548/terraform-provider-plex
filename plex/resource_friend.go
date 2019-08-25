package plex

import (
	"fmt"
	plexclient "github.com/akak548/go-plex-client"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

type friendNotFound struct {
	Id string
}

func (f friendNotFound) Error() string {
	return fmt.Sprintf("Unable to find %v", f.Id)
}

func findFriend(f string, fs []plexclient.Friends) (plexclient.Friends, error) {

	for _, friend := range fs {
		if friend.Username == f || friend.Email == f {
			return friend, nil
		}
	}
	return plexclient.Friends{}, friendNotFound{f}
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
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_sync": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     false,
				Description: "Allow users are able to make use of Mobile Sync to sync content to their mobile devices",
			},
			"allow_cameraupload": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     false,
				Description: "Allow users are able to upload media to server",
			},
			"allow_channels": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     false,
				Description: "Allow users are able to make use of Mobile Sync to sync content to their mobile devices",
			},
			"machine_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Allow user access to machine_id",
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceFriendCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*plexclient.Plex)

	params := plexclient.InviteFriendParams{
		MachineID: d.Get("machine_id").(string),
	}

	if v, ok := d.GetOk("username"); ok {
		params.UsernameOrEmail = v.(string)
	} else if v, ok := d.GetOk("email_address"); ok {
		params.UsernameOrEmail = v.(string)
	} else {
		return fmt.Errorf("Error: Must provide username or email_adddress for %s", d.Id())
	}

	fmt.Println(params)
	newFriendId, err := client.InviteFriend(params)

	if err != nil {
		return fmt.Errorf("Error inviting friend %d/%s: %s", newFriendId, d.Get("email_address").(string), err)
	}

	d.SetId(params.UsernameOrEmail)
	d.Set("user_id", fmt.Sprintf("%d", d.Set("test_id", newFriendId)))
	d.Set("invite", "pending")
	return resourceFriendRead(d, m)
}

func resourceFriendRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*plexclient.Plex)

	invite_state := d.Get("invite")

	if invite_state == "pending" {

	} else {

	}

	friends, err := client.GetFriends()
	if err != nil {
		log.Printf("[WARN] No friends found: %s", d.Id())
		d.SetId("")
		return nil
	}

	friend, err := findFriend(d.Id(), friends)
	if err != nil {
		return fmt.Errorf("Error finding friend %s: Error: %s", d.Id(), err)
	}

	d.Set("email_address", friend.Email)
	d.Set("username", friend.Username)
	d.Set("allow_sync", friend.AllowSync)
	d.Set("allow_cameraupload", friend.AllowCameraUpload)
	d.Set("allow_channels", friend.AllowChannels)
	d.Set("machine_id", friend.Server.MachineIdentifier)
	return nil
}

func resourceFriendUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFriendRead(d, m)
}

func resourceFriendDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
