package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,

		Schema: map[string]*schema.Schema{
			"handle": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// https://www.aptible.com/documentation/enclave/cli/apps-create.html
func resourceAppCreate(d *schema.ResourceData, m interface{}) error {
	handle := d.Get("handle").(string)
	d.SetId(handle)
	return resourceAppRead(d, m)
}

func resourceAppRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAppUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAppRead(d, m)
}

// https://www.aptible.com/documentation/enclave/cli/apps-deprovision.html
func resourceAppDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
