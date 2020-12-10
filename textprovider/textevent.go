package textprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func textEvent() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventCreate,
		Read:   resourceEventRead,
		Update: resourceEventUpdate,
		Delete: resourceEventDelete,

		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stored_content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEventCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId("1")
	return resourceEventRead(d, meta)
}

func resourceEventRead(d *schema.ResourceData, meta interface{}) error {
	d.Set("stored_content", d.Get("content"))
	return nil
}

func resourceEventUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceEventDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
