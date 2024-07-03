// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r HostnameTLSSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"setting_id": schema.StringAttribute{
				Description: "The TLS Setting name.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ciphers", "min_tls_version", "http2"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description:   "The hostname for which the tls settings are set.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value": schema.Float64Attribute{
				Description: "The tls setting value.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "This is the time the tls setting was originally created for this hostname.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Description: "This is the time the tls setting was updated.",
				Computed:    true,
			},
		},
	}
}