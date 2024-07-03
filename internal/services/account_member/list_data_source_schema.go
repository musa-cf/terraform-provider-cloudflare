// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &AccountMembersDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccountMembersDataSource{}

func (r AccountMembersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order results.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"order": schema.StringAttribute{
				Description: "Field to order results by.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("user.first_name", "user.last_name", "user.email", "status"),
				},
			},
			"page": schema.Float64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "Maximum number of results per page.",
				Computed:    true,
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "A member's status in the account.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("accepted", "pending", "rejected"),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Membership identifier tag.",
							Computed:    true,
						},
						"policies": schema.ListNestedAttribute{
							Description: "Access policy for the membership",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Policy identifier.",
										Computed:    true,
									},
									"access": schema.StringAttribute{
										Description: "Allow or deny operations against the resources.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("allow", "deny"),
										},
									},
									"permission_groups": schema.ListNestedAttribute{
										Description: "A set of permission groups that are specified to the policy.",
										Computed:    true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description: "Identifier of the group.",
													Computed:    true,
												},
												"meta": schema.StringAttribute{
													Description: "Attributes associated to the permission group.",
													Computed:    true,
												},
												"name": schema.StringAttribute{
													Description: "Name of the group.",
													Computed:    true,
												},
											},
										},
									},
									"resource_groups": schema.ListNestedAttribute{
										Description: "A list of resource groups that the policy applies to.",
										Computed:    true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description: "Identifier of the group.",
													Computed:    true,
												},
												"scope": schema.ListNestedAttribute{
													Description: "The scope associated to the resource group",
													Computed:    true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description: "This is a combination of pre-defined resource name and identifier (like Account ID etc.)",
																Computed:    true,
															},
															"objects": schema.ListNestedAttribute{
																Description: "A list of scope objects for additional context.",
																Computed:    true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description: "This is a combination of pre-defined resource name and identifier (like Zone ID etc.)",
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"meta": schema.StringAttribute{
													Description: "Attributes associated to the resource group.",
													Computed:    true,
												},
												"name": schema.StringAttribute{
													Description: "Name of the resource group.",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"roles": schema.ListNestedAttribute{
							Description: "Roles assigned to this Member.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Role identifier tag.",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Description: "Description of role's permissions.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Role Name.",
										Computed:    true,
									},
									"permissions": schema.ListAttribute{
										Description: "Access permissions for this User.",
										Computed:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"status": schema.StringAttribute{
							Description: "A member's status in the account.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("accepted", "pending"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *AccountMembersDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccountMembersDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}