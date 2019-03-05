package aptible

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAptibleApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAptibleAppRead,
		Schema: map[string]*schema.Schema{
			"handle": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"git_repo": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAptibleAppRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	app, err := client.GetAppByHandle(d.Get("handle").(string))

	if err != nil {
		return err
	}

	d.SetId(string(app.ID))
	setAppDetails(d, app)

	return nil
}
