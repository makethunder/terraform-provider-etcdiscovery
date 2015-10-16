package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
)

func resourceToken() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},

		Create: resourceTokenSet,
		Read:   resourceTokenRead,
		Delete: resourceTokenDelete,
	}
}

func resourceTokenSet(d *schema.ResourceData, meta interface{}) error {
	url := meta.(string)
	size := d.Get("key")

	if size != nil {
		url += fmt.Sprintf("?size=%d", size.(int))
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not get discovery token: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read body of discovery token request: %s", err)
	}

	d.SetId(string(body))
	return nil
}

func resourceTokenRead(d *schema.ResourceData, meta interface{}) error {
	// we could conceivably fetch the token and set some computed attributes
	// about the cluster, but I can't think of a use case for that. So, not
	// much to do here.
	return nil
}

func resourceTokenDelete(d *schema.ResourceData, meta interface{}) error {
	// discovery tokens automatically expire and there's no way to explicitly
	// delete them, so just the resource from Terraform's state.
	d.SetId("")
	return nil
}
