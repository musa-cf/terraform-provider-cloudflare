// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &WorkerSecretDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WorkerSecretDataSource{}

func (r WorkerSecretDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of this secret, this is what will be to access it inside the Worker.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of secret to put.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("secret_text"),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"dispatch_namespace": schema.StringAttribute{
						Description: "Name of the Workers for Platforms dispatch namespace.",
						Required:    true,
					},
					"script_name": schema.StringAttribute{
						Description: "Name of the script, used in URLs and route configuration.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *WorkerSecretDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WorkerSecretDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}