package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uptimerobotapi "github.com/vexxhost/terraform-provider-uptimerobot/internal/provider/api"
)

func dataSourceAlertContacts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlertContactsRead,
		Schema: map[string]*schema.Schema{
			"contacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"friendly_name": {Type: schema.TypeString, Computed: true},
						"id":            {Type: schema.TypeString, Computed: true},
						"type":          {Type: schema.TypeString, Computed: true},
						"status":        {Type: schema.TypeString, Computed: true},
						"value":         {Type: schema.TypeString, Computed: true},
					},
				},
			},
		},
	}
}

func dataSourceAlertContactsRead(d *schema.ResourceData, m interface{}) error {
	account, err := m.(uptimerobotapi.UptimeRobotApiClient).GetAccountDetails()
	if err != nil {
		return err
	}
	d.SetId(account.Email)

	alertContacts, err := m.(uptimerobotapi.UptimeRobotApiClient).GetAlertContacts()
	if err != nil {
		return err
	}

	contacts := make([]map[string]interface{}, 0, len(alertContacts))

	for _, alertContact := range alertContacts {
		contact := map[string]interface{}{
			"friendly_name": alertContact.FriendlyName,
			"id":            alertContact.ID,
			"type":          alertContact.Type,
			"status":        alertContact.Status,
			"value":         alertContact.Value,
		}
		contacts = append(contacts, contact)
	}

	if err := d.Set("contacts", contacts); err != nil {
		return err
	}

	return nil
}
