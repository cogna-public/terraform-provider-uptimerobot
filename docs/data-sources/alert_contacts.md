---
page_title: "uptimerobot_alert_contacts Data Source - terraform-provider-uptimerobot"
subcategory: ""
description: |-
  Get all alert contacts
---

# Data Source: uptimerobot_alert_contacts

Use this data source to list all alert contacts on the account

## Example Usage

```hcl
# get by friendly name
data "uptimerobot_alert_contacts" "this" {}
```

## Attributes Reference

Each alert contact will have the following attributes:

* `friendly_name` - the friendly name of the alert contact
* `id` - the id of the alert contact
* `type` - the type of the alert contact
* `status` - the status of the alert contact
* `value` - the value of the alert contact (e.g email address or Slack webhook URL)

