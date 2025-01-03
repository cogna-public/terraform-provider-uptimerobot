package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uptimerobotapi "github.com/vexxhost/terraform-provider-uptimerobot/internal/provider/api"
)

func dataSourceAlertContact() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlertContactRead,
		Schema: map[string]*schema.Schema{
			"friendly_name": {Optional: true, Type: schema.TypeString},
			"id":            {Computed: true, Type: schema.TypeString},
			"type":          {Computed: true, Type: schema.TypeString},
			"status":        {Computed: true, Type: schema.TypeString},
			"value":         {Optional: true, Type: schema.TypeString},
		},
	}
}

func dataSourceAlertContactRead(d *schema.ResourceData, m interface{}) error {
	alertContacts, err := m.(uptimerobotapi.UptimeRobotApiClient).GetAlertContacts()
	if err != nil {
		return err
	}

	friendlyName := d.Get("friendly_name").(string)
	value := d.Get("value").(string)

	var alertContact uptimerobotapi.AlertContact

	if friendlyName != "" {
		for _, a := range alertContacts {
			if friendlyName != "" && a.FriendlyName == friendlyName {
				alertContact = a
				break
			}
		}
	} else if value != "" {
		for _, a := range alertContacts {
			if value != "" && a.Value == value {
				alertContact = a
				break
			}
		}
	}

	if alertContact == (uptimerobotapi.AlertContact{}) {
		if friendlyName != "" {
			return fmt.Errorf("failed to find alert contact by friendly_name %s", friendlyName)
		} else if value != "" {
			return fmt.Errorf("failed to find alert contact by value %s", value)
		}
	}

	d.SetId(alertContact.ID)

	d.Set("friendly_name", alertContact.FriendlyName)
	d.Set("type", alertContact.Type)
	d.Set("status", alertContact.Status)
	d.Set("value", alertContact.Value)

	return nil
}
