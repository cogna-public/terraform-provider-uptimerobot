package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	uptimerobotapi "github.com/vexxhost/terraform-provider-uptimerobot/internal/provider/api"
)

// TODO - you can use only 1 active alert contact
func TestUptimeRobotDataResourceAlertContact_email(t *testing.T) {
	t.Skip("UptimeRobot response: you can use only 1 active alert contact")
	var email = "louay+tftest@alakkad.me"
	var friendlyName = "TF Test: Email"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAlertContactDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test" {
					friendly_name = "%s"
					type          = "%s"
					value         = "%s"
				}
				`, friendlyName, "e-mail", email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "type", "e-mail"),
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "value", email),
				),
			},
			{
				ResourceName:      "uptimerobot_alert_contact.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceAlertContact_update_email(t *testing.T) {
	t.Skip("UptimeRobot response: you can use only 1 active alert contact")
	var email = "louay+tftest@alakkad.me"
	var email2 = "louay+tftest2@alakkad.me"
	var friendlyName = "TF Test: Email"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAlertContactDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test" {
					friendly_name = "%s"
					type          = "%s"
					value         = "%s"
				}
				`, friendlyName, "e-mail", email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "type", "e-mail"),
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "value", email),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test" {
					friendly_name = "%s"
					type          = "%s"
					value         = "%s"
				}
				`, friendlyName, "e-mail", email2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "value", email2),
				),
			},
			{
				ResourceName:      "uptimerobot_alert_contact.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceAlertContact_sms(t *testing.T) {
	t.Skip("API seems to reject this")

	var tel = "00447870000000"
	var friendlyName = "TF Test: SMS"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAlertContactDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test" {
					friendly_name = "%s"
					type          = "%s"
					value         = "%s"
				}
				`, friendlyName, "sms", tel),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "type", "sms"),
					resource.TestCheckResourceAttr("uptimerobot_alert_contact.test", "value", tel),
				),
			},
			{
				ResourceName:      "uptimerobot_alert_contact.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAlertContactDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(uptimerobotapi.UptimeRobotApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "uptimerobot_alert_contact" {
			continue
		}

		id := rs.Primary.ID

		_, err := client.GetAlertContact(id)

		if err == nil {
			return fmt.Errorf("Alert contact still exists")
		}

		// Verify the error is what we want
		if strings.Contains(err.Error(), "test") {
			return err
		}
	}

	return nil
}
