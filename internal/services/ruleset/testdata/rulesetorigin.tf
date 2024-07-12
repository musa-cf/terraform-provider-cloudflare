
  resource "cloudflare_ruleset" "%[1]s" {
    zone_id  = "%[3]s"
    name        = "%[2]s"
    description = "%[1]s ruleset description"
    kind        = "zone"
    phase       = "http_request_origin"

    rules =[ {
      action = "route"
      action_parameters = {
    host_header = "%[1]s.%[4]s"
        origin = {
    host = "%[1]s.%[4]s"
          port = 80
  }
        sni = {
    value = "%[1]s.%[4]s"
  }
  }
      expression = "(http.request.uri.path matches \"^/api/\")"
      description = "example http request origin"
      enabled = true
    }]
  }