package provider

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	uptimerobotapi "github.com/vexxhost/terraform-provider-uptimerobot/internal/provider/api"
)

func TestUptimeRobotDataResourceMonitor_http_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http monitor"
	var Type = "http"
	var URL = "https://google.com"
	var URL2 = "https://yahoo.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL2),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_keyword_monitor(t *testing.T) {
	var FriendlyName = "TF Test: keyword"
	var Type = "keyword"
	var URL = "https://google.com"
	var KeywordType = "not exists"
	var KeywordType2 = "exists"
	var KeywordValue = "yahoo"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					keyword_type  = "%s"
					keyword_value = "%s"
				}
				`, FriendlyName, Type, URL, KeywordType, KeywordValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_type", KeywordType),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_value", KeywordValue),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					keyword_type  = "%s"
					keyword_value = "%s"
				}
				`, FriendlyName, Type, URL, KeywordType2, KeywordValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_type", KeywordType2),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_http_port_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http port monitor"
	var Type = "port"
	var URL = "google.com"
	var URL2 = "yahoo.com"
	var SubType = "http"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					sub_type      = "%s"
				}
				`, FriendlyName, Type, URL, SubType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "sub_type", SubType),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					sub_type      = "%s"
				}
				`, FriendlyName, Type, URL2, SubType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL2),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_port_monitor(t *testing.T) {
	var FriendlyName = "TF Test: custom port monitor"
	var Type = "port"
	var URL = "google.com"
	var SubType = "custom"
	var Port = 8080
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					sub_type      = "%s"
					port          = %d
				}
				`, FriendlyName, Type, URL, SubType, Port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "sub_type", SubType),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "port", fmt.Sprintf(`%d`, Port)),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_ignore_ssl_errors(t *testing.T) {
	var FriendlyName = "TF Test:  custom ignore ssl errors"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name     = "%s"
					type              = "%s"
					url               = "%s"
					ignore_ssl_errors = true
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "ignore_ssl_errors", "true"),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}
func TestUptimeRobotDataResourceMonitor_custom_alert_contact_threshold_and_recurrence(t *testing.T) {
	t.Skip("UptimeRobot response: you can use only 1 active alert contact")
	var FriendlyName = "TF Test: custom alert contact threshold & recurrence"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test" {
					friendly_name = "SRE Team"
					type          = "e-mail"
					value         = "sre@vexxhost.com"
				}
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = uptimerobot_alert_contact.test.id
						threshold  = 0
						recurrence = 0
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "1"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.recurrence", "0"),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_alert_contacts(t *testing.T) {
	t.Skip("UptimeRobot response: you can use only 1 active alert contact")
	var FriendlyName = "TF Test: custom alert contacts"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test1" {
					friendly_name = "Test 1"
					type          = "e-mail"
					value         = "test1@example.com"
				}

				resource "uptimerobot_alert_contact" "test2" {
					friendly_name = "Test 2"
					type          = "e-mail"
					value         = "test2@example.com"
				}

				resource "uptimerobot_alert_contact" "test3" {
					friendly_name = "Test 3"
					type          = "e-mail"
					value         = "test3@example.com"
				}

				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = uptimerobot_alert_contact.test1.id
						threshold  = 0
						recurrence = 0
					}
					alert_contact {
						id         = uptimerobot_alert_contact.test2.id
						threshold  = 0
						recurrence = 0
					}
					alert_contact {
						id         = uptimerobot_alert_contact.test3.id
						threshold  = 0
						recurrence = 0
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "3"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.recurrence", "0"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.1.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.1.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.1.recurrence", "0"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.2.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.2.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.2.recurrence", "0"),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_http_headers(t *testing.T) {
	var FriendlyName = "TF Test:  custom http headers"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					custom_http_headers = {
						// Accept-Language = "en"
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "custom_http_headers.%", "0"),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_ping_monitor(t *testing.T) {
	var FriendlyName = "TF Test: ping monitor"
	var Type = "ping"
	var URL = "1.1.1.1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_interval(t *testing.T) {
	var FriendlyName = "TF Test: custom interval"
	var Type = "ping"
	var URL = "1.1.1.1"
	var Interval = 300
	var Interval2 = 360
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					interval      = %d
				}
				`, FriendlyName, Type, URL, Interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "interval", fmt.Sprintf(`%d`, Interval)),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					interval      = %d
				}
				`, FriendlyName, Type, URL, Interval2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "interval", fmt.Sprintf(`%d`, Interval2)),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_http_method(t *testing.T) {
	var FriendlyName = "TF Test: http method monitor"
	var Type = "http"
	var URL = "https://httpbin.org/post"
	var Method = "POST"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_method    = "%s"
				}
				`, FriendlyName, Type, URL, Method),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_method", Method),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_http_auth_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http auth monitor"
	var Type = "http"
	var Username = "tester"
	var Password = "secret"
	var AuthType = "basic"
	var AuthType2 = "digest"
	var URL = fmt.Sprintf("https://httpbin.org/basic-auth/%s/%s", Username, Password)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_username  = "%s"
					http_password  = "%s"
					http_auth_type = "%s"
				}
				`, FriendlyName, Type, URL, Username, Password, AuthType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_username", Username),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_password", Password),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_auth_type", AuthType),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_auth_type issue
				// ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_username  = "%s"
					http_password  = "%s"
					http_auth_type = "%s"
				}
				`, FriendlyName, Type, URL, Username, Password, AuthType2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_username", Username),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_password", Password),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_auth_type", AuthType2),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_auth_type issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_default_alert_contact(t *testing.T) {
	var FriendlyName = "TF Test: using the default alert contact"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

				data "uptimerobot_account" "account" {}

				data "uptimerobot_alert_contact" "default" {
				  value = data.uptimerobot_account.account.email
				}

				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = data.uptimerobot_alert_contact.default.id
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "1"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
				),
			},
			{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckMonitorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(uptimerobotapi.UptimeRobotApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "uptimerobot_monitor" {
			continue
		}

		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = client.GetMonitor(id)

		if err == nil {
			return fmt.Errorf("Monitor still exists")
		}

		// Verify the error is what we want
		if strings.Contains(err.Error(), "test") {
			return err
		}
	}

	return nil
}
