---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dcloud_topologies Data Source - terraform-provider-dcloud"
subcategory: ""
description: |-
  All the topologies owned or shared to the authenticated user
---

# dcloud_topologies (Data Source)

All the topologies owned or shared to the authenticated user



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `topologies` (List of Object) (see [below for nested schema](#nestedatt--topologies))

<a id="nestedatt--topologies"></a>
### Nested Schema for `topologies`

Read-Only:

- `datacenter` (String)
- `description` (String)
- `name` (String)
- `notes` (String)
- `status` (String)
- `uid` (String)


