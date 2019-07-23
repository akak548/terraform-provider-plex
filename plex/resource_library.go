package plex

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceLibrary() *schema.Resource {
    return &schema.Resource {
        Create: resourceLibraryCreate,
        Read: resourceLibraryRead,
        Update: resourceLibraryRead,
        Delete: resourceLibraryRead,

        Schema: map[string]*schema.Schema {
            "email": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceLibraryCreate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceLibraryRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceLibraryUpdate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceLibraryDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}
