---
page_title: "cloudavenue_s3_bucket_versioning_configuration Data Source - cloudavenue"
subcategory: "S3 (Object Storage)"
description: |-
  The cloudavenue_s3_bucket_versioning_configuration data source allows you to retrieve information about an S3 bucket's versioning configuration.
---

# cloudavenue_s3_bucket_versioning_configuration (Data Source)

The `cloudavenue_s3_bucket_versioning_configuration` data source allows you to retrieve information about an S3 bucket's versioning configuration.

## Example Usage

```terraform
data "cloudavenue_s3_bucket_versioning_configuration" "example" {
  bucket = cloudavenue_s3_bucket.example.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `bucket` (String) The name of the bucket.

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) The ID is a bucket name.
- `status` (String) Versioning state of the bucket.

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
