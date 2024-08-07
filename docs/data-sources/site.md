---
page_title: "jamfpro_site"
description: |-
  
---

# jamfpro_site (Data Source)


## Example Usage
```terraform
data "jamfpro_site" "site_001_data" {
  id = jamfpro_site.site_001.id
}

output "jamfpro_site_001_id" {
  value = data.jamfpro_site.site_001_data.id
}

output "jamfpro_site_001_name" {
  value = data.jamfpro_site.site_001_data.name
}

data "jamfpro_sites" "site_002_data" {
  id = jamfpro_site.site_002.id
}

output "jamfpro_site_002_id" {
  value = data.jamfpro_site.site_002_data.id
}

output "jamfpro_site_002_name" {
  value = data.jamfpro_site.site_002_data.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The unique identifier of the Jamf Pro site.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `name` (String) The unique name of the Jamf Pro site.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)