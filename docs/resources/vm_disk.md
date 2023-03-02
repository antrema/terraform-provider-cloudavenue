---
page_title: "cloudavenue_vm_disk Resource - cloudavenue"
subcategory: "VM"
description: |-
  The disk resource allows you to manage an disk of a VM .
---

# cloudavenue_vm_disk (Resource)

The disk resource allows you to manage an disk of a VM .

## Example Usage

```terraform
resource "cloudavenue_vm_disk" "example" {
	vapp_name       = "vapp_test3"
	vm_name         = "TestRomain"
	allow_vm_reboot = true
	internal_disk = {
		bus_type      = "sata"
		size_in_mb    = "500"
		bus_number    = 0
		unit_number   = 1
	}
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vapp_name` (String) The vApp this VM disk belongs to.
- `vm_name` (String) VM in vApp in which disk is created.

### Optional

- `allow_vm_reboot` (Boolean) Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail.
- `internal_disk` (Attributes) A block to define disk. Multiple can be used. (see [below for nested schema](#nestedatt--internal_disk))
- `vdc` (String) The name of VDC to use, optional if defined at provider level.

### Read-Only

- `id` (String) The ID of the internal disk.

<a id="nestedatt--internal_disk"></a>
### Nested Schema for `internal_disk`

Required:

- `bus_number` (Number) The number of the `SCSI` or `IDE` controller itself.
- `bus_type` (String) The type of disk controller. Possible values: `ide`, `parallel` (LSI Logic Parallel SCSI), `sas` (LSI Logic SAS SCSI), `paravirtual` (Paravirtual SCSI), `sata`, `nvme`.
- `size_in_mb` (Number) The size of the disk in MB.
- `unit_number` (Number) The device number on the `SCSI` or `IDE` controller of the disk.

Optional:

- `storage_profile` (String) Storage profile to override the VM default one.

Read-Only:

- `id` (String) The ID of the internal disk.

## Import

Import is supported using the following syntax:
```shell
# if vdc is not specified, the default vdc will be used
terraform import cloudavenue_vm_disk.example vapp_name.vm_name.id

# if vdc is specified, the vdc will be used
terraform import cloudavenue_vm_disk.example vdc.vapp_name.vm_name.id
```