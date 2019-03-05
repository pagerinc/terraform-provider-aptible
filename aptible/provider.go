package aptible

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APTIBLE_EMAIL", nil),
				Description: "Aptible bot email.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APTIBLE_PASSWORD", nil),
				Description: "Aptible bot password.",
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APTIBLE_ENVIRONMENT", nil),
				Description: "Aptible Environment.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"aptible_app": resourceAptibleApp(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"aptible_app": dataSourceAptibleApp(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Println("[INFO] Initializing Aptible provider")
	client := NewClient()
	if email, ok := d.GetOk("email"); ok {
		client.Email = email.(string)
	}
	if password, ok := d.GetOk("password"); ok {
		client.APIKey = password.(string)
	}
	if env, ok := d.GetOk("environment"); ok {
		client.Environment = env.(string)
	}

	if err := client.Login(); err != nil {
		return nil, err
	}

	return client, nil
}
