---
page_title: "cloudavenue_network_app_port_profile Resource - cloudavenue"
subcategory: "Network"
description: |-
  Provides a NSX-T App Port Profile resource
---

# cloudavenue_network_app_port_profile (Resource)

Provides a NSX-T App Port Profile resource

## Example Usage

```terraform
data "cloudavenue_vdc" "example" {
  name = "VDC_Test"
}

resource "cloudavenue_network_app_port_profile" "example" {
  name        = "example-rule"
  description = "Application port profile for example"
  vdc         = data.cloudavenue_vdc.example.id

  app_ports = [
    {
      protocol = "ICMPv4"
    },
    {
      protocol = "TCP"
      ports = [
        "80",
        "443",
      ]
    },
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_ports` (Attributes List) List of application ports. (see [below for nested schema](#nestedatt--app_ports))
- `name` (String) Application Port Profile name.
- `vdc` (String) (ForceNew) ID of VDC or VDC Group.

### Optional

- `description` (String) Application Port Profile description.

### Read-Only

- `id` (String) The ID of the VM.

<a id="nestedatt--app_ports"></a>
### Nested Schema for `app_ports`

Required:

- `protocol` (String) Protocol. Value must be one of : `ICMPv4`, `ICMPv6`, `TCP`, `UDP`.

Optional:

- `ports` (Set of String) Set of ports or ranges.

## Import

Import is supported using the following syntax:
```shell
terraform import vdc-or-vdc-group-id.RuleName
```