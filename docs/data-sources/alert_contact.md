---
page_title: "uptimerobot_alert_contact Data Source - terraform-provider-uptimerobot"
subcategory: ""
description: |-
  Get an alert contact
---

# Data Source: uptimerobot_alert_contact

Use this data source to get information about an alert contact by friendly_name or value

## Example Usage

```hcl
# get by friendly name
data "uptimerobot_alert_contact" "this" {
    friendly_name = "My Alert Contact"
}
# get by value
data "uptimerobot_alert_contact" "this" {
    value = "foo@bar.com"
}
```

## Attributes Reference

* `friendly_name` - the friendly name of the alert contact
* `id` - the id of the alert contact
* `type` - the type of the alert contact
* `status` - the status of the alert contact
* `value` - the value of the alert contact (e.g email address or Slack webhook URL)

