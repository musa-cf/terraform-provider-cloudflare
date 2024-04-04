// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rules

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRulesResultEnvelope struct {
	Result EmailRoutingRulesModel `json:"result,computed"`
}

type EmailRoutingRulesModel struct {
	ZoneIdentifier types.String                      `tfsdk:"zone_identifier" path:"zone_identifier"`
	RuleIdentifier types.String                      `tfsdk:"rule_identifier" path:"rule_identifier"`
	Actions        []*EmailRoutingRulesActionsModel  `tfsdk:"actions" json:"actions"`
	Matchers       []*EmailRoutingRulesMatchersModel `tfsdk:"matchers" json:"matchers"`
	Enabled        types.Bool                        `tfsdk:"enabled" json:"enabled"`
	Name           types.String                      `tfsdk:"name" json:"name"`
	Priority       types.Float64                     `tfsdk:"priority" json:"priority"`
	ID             types.String                      `tfsdk:"id" json:"id,computed"`
	Tag            types.String                      `tfsdk:"tag" json:"tag,computed"`
}

type EmailRoutingRulesActionsModel struct {
	Type  types.String   `tfsdk:"type" json:"type"`
	Value []types.String `tfsdk:"value" json:"value"`
}

type EmailRoutingRulesMatchersModel struct {
	Field types.String `tfsdk:"field" json:"field"`
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}