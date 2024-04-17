// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_jobs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r LogpushJobsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "Unique id of the job.",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"destination_conf": schema.StringAttribute{
				Description: "Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.",
				Required:    true,
			},
			"dataset": schema.StringAttribute{
				Description: "Name of the dataset.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Flag that indicates if the job is enabled.",
				Optional:    true,
			},
			"frequency": schema.StringAttribute{
				Description: "The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("high", "low"),
				},
				Default: stringdefault.StaticString("high"),
			},
			"logpull_options": schema.StringAttribute{
				Description: "This field is deprecated. Use `output_options` instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Optional human readable job name. Not unique. Cloudflare suggests that you set this to a meaningful string, like the domain name, to make it easier to identify your job.",
				Optional:    true,
			},
			"output_options": schema.SingleNestedAttribute{
				Description: "The structured replacement for `logpull_options`. When including this field, the `logpull_option` field will be ignored.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"batch_prefix": schema.StringAttribute{
						Description: "String to be prepended before each batch.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(""),
					},
					"batch_suffix": schema.StringAttribute{
						Description: "String to be appended after each batch.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(""),
					},
					"cve_2021_4428": schema.BoolAttribute{
						Description: "If set to true, will cause all occurrences of `${` in the generated files to be replaced with `x{`.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"field_delimiter": schema.StringAttribute{
						Description: "String to join fields. This field be ignored when `record_template` is set.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(","),
					},
					"field_names": schema.ListAttribute{
						Description: "List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"output_type": schema.StringAttribute{
						Description: "Specifies the output type, such as `ndjson` or `csv`. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ndjson", "csv"),
						},
						Default: stringdefault.StaticString("ndjson"),
					},
					"record_delimiter": schema.StringAttribute{
						Description: "String to be inserted in-between the records as separator.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(""),
					},
					"record_prefix": schema.StringAttribute{
						Description: "String to be prepended before each record.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("{"),
					},
					"record_suffix": schema.StringAttribute{
						Description: "String to be appended after each record.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("}\n"),
					},
					"record_template": schema.StringAttribute{
						Description: "String to use as template for each record instead of the default comma-separated list. All fields used in the template must be present in `field_names` as well, otherwise they will end up as null. Format as a Go `text/template` without any standard functions, like conditionals, loops, sub-templates, etc.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(""),
					},
					"sample_rate": schema.Float64Attribute{
						Description: "Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current `sample_interval` of the data.",
						Computed:    true,
						Optional:    true,
						Default:     float64default.StaticFloat64(1),
					},
					"timestamp_format": schema.StringAttribute{
						Description: "String to specify the format for timestamps, such as `unixnano`, `unix`, or `rfc3339`.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("unixnano", "unix", "rfc3339"),
						},
						Default: stringdefault.StaticString("unixnano"),
					},
				},
			},
			"ownership_challenge": schema.StringAttribute{
				Description: "Ownership challenge token to prove destination ownership.",
				Optional:    true,
			},
		},
	}
}