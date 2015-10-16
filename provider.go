package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://discovery.etcd.io/new",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"etcdiscovery_token": resourceToken(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return d.Get("url"), nil
}
