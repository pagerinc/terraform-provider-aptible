package aptible

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAptibleApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAptibleAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,

		Schema: map[string]*schema.Schema{
			"handle": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"git_remote": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// https://www.aptible.com/documentation/enclave/cli/apps-create.html
func resourceAptibleAppCreate(d *schema.ResourceData, m interface{}) error {
	handle := d.Get("handle").(string)
	d.SetId(handle)
	return resourceAppRead(d, m)
}

func resourceAppRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	// Attempt to read from an upstream API
	app, err := client.GetAppByHandle(d.Get("handle").(string))

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return nil
	}

	d.SetId(string(app.ID))
	setAppDetails(d, app)
	return nil
}

func resourceAppUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAppRead(d, m)
}

// https://www.aptible.com/documentation/enclave/cli/apps-deprovision.html
func resourceAppDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func setAppDetails(d *schema.ResourceData, app *App) (err error) {
	d.Set("handle", app.Handle)
	d.Set("git_repo", app.GitRemote)
	d.Set("status", app.Status)
	return err
}
