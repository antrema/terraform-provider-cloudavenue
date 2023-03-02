---
page_title: "cloudavenue_vm_inserted_media Resource - cloudavenue"
subcategory: "VM (Virtual Machine)"
description: |-
  The inserted_media resource resource for inserting or ejecting media (ISO) file for the VM. Create this resource for inserting the media, and destroy it for ejecting.

---

# cloudavenue_vm_inserted_media (Resource)

The inserted_media resource resource for inserting or ejecting media (ISO) file for the VM. Create this resource for inserting the media, and destroy it for ejecting.

~> Only one media is allowed per VM.

## Example Usage

```terraform
resource "cloudavenue_vm_inserted_media" "example" {
	catalog = "catalog-example"
	name    = "debian-9.9.0-amd64-netinst.iso"
	vapp_name = "vapp-example"
	vm_name   = "vm-example"
  }
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `catalog` (String) The name of the catalog where to find media file
- `name` (String) Media file name in catalog which will be inserted to VM
- `vapp_name` (String) vApp name where VM is located
- `vm_name` (String) VM name where media will be inserted or ejected

### Optional

- `vdc` (String) The name of VDC to use, optional if defined at provider level

### Read-Only

- `id` (String) The ID of the inserted media. This is the vm Id where the media is inserted.

## Import

Import is not supported for this resource