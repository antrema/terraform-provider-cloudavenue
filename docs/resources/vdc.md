---
page_title: "cloudavenue_vdc Resource - cloudavenue"
subcategory: "vDC (Virtual Datacenter)"
description: |-
  Provides a Cloud Avenue vDC (Virtual Data Center) resource. This can be used to create, update and delete vDC.
---

# cloudavenue_vdc (Resource)

Provides a Cloud Avenue vDC (Virtual Data Center) resource. This can be used to create, update and delete vDC.
 
 -> Note: For more information about Cloud Avenue vDC, please refer to the [Cloud Avenue documentation](https://wiki.cloudavenue.orange-business.com/wiki/Datacenter_virtuel).

## Example Usage

```terraform
resource "cloudavenue_vdc" "example" {
  name                  = "MyVDC"
  description           = "Example VDC created by Terraform"
  cpu_allocated         = 22000
  memory_allocated      = 30
  cpu_speed_in_mhz      = 2200
  billing_model         = "PAYG"
  disponibility_class   = "ONE-ROOM"
  service_class         = "STD"
  storage_billing_model = "PAYG"

  storage_profiles = [
    {
      class   = "gold"
      default = true
      limit   = 500
    },
    {
      class   = "silver"
      default = false
      limit   = 500
    },
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `billing_model` (String) (ForceNew) Choose Billing model of compute resources. Value must be one of : `PAYG`, `DRAAS`, `RESERVED`.
- `cpu_allocated` (Number) CPU capacity in *MHz* that is committed to be available or used as a limit in PAYG mode. 

 -> Note: Reserved capacity is automatically set according to the service class.
- `cpu_speed_in_mhz` (Number) (ForceNew) Specifies the clock frequency, in Mhz, for any virtual CPU that is allocated to a VM. Force replacement attributes, however you can change the `cpu_speed_in_mhz` attribute only if the `billing_model` is set to **RESERVED**. Value must be at least 1200.
- `disponibility_class` (String) (ForceNew) The disponibility class of the vDC. Value must be one of : `ONE-ROOM`, `DUAL-ROOM`, `HA-DUAL-ROOM`.
- `memory_allocated` (Number) Memory capacity in Gb that is committed to be available or used as a limit in PAYG mode. Value must be between 1 and 500.
- `name` (String) (ForceNew) The name of the vDC. String length must be between 2 and 27.
- `service_class` (String) (ForceNew) The service class of the vDC. Value must be one of : `ECO`, `STD`, `HP`, `VOIP`.
- `storage_billing_model` (String) (ForceNew) Choose Billing model of storage resources. Value must be one of : `PAYG`, `RESERVED`.
- `storage_profiles` (Attributes Set) List of storage profiles for this vDC. Set must contain at least 1 elements. (see [below for nested schema](#nestedatt--storage_profiles))

### Optional

- `description` (String) A description of the vDC.
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) The ID of the vDC.

<a id="nestedatt--storage_profiles"></a>
### Nested Schema for `storage_profiles`

Required:

- `class` (String) The storage class of the storage profile. Value must be one of : `silver`, `silver_r1`, `silver_r2`, `gold`, `gold_r1`, `gold_r2`, `gold_hm`, `platinum3k`, `platinum3k_r1`, `platinum3k_r2`, `platinum3k_hm`, `platinum7k`, `platinum7k_r1`, `platinum7k_r2`, `platinum7k_hm`.
- `default` (Boolean) Set this storage profile as default for this vDC. Only one storage profile can be default per vDC.
- `limit` (Number) Max number in *Gb* of units allocated for this storage profile. Value must be between 500 and 10000.


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `read` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).

## Import

Import is supported using the following syntax:
```shell
# VDC can be imported using the name.

terraform import cloudavenue_vdc.example name
```