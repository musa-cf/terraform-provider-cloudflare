---
page_title: "cloudflare_access_service_tokens Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_access_service_tokens (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `max_items` (Number) Max items to fetch, default: 1000
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Optional:

- `client_id` (String) The Client ID for the service token. Access will check for this value in the `CF-Access-Client-ID` request header.
- `created_at` (String)
- `duration` (String) The duration for how long the service token will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The default is 1 year in hours (8760h).
- `id` (String) The ID of the service token.
- `name` (String) The name of the service token.
- `updated_at` (String)

