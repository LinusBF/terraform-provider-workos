// Copyright (c) OSO DevOps
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/osodevops/terraform-provider-workos/internal/client"
)

var _ resource.Resource = &WebhookEndpointResource{}
var _ resource.ResourceWithImportState = &WebhookEndpointResource{}

func NewWebhookEndpointResource() resource.Resource {
	return &WebhookEndpointResource{}
}

type WebhookEndpointResource struct {
	client *client.Client
}

type WebhookEndpointResourceModel struct {
	ID          types.String `tfsdk:"id"`
	EndpointURL types.String `tfsdk:"endpoint_url"`
	Events      types.Set    `tfsdk:"events"`
	Status      types.String `tfsdk:"status"`
	Secret      types.String `tfsdk:"secret"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

func (r *WebhookEndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook_endpoint"
}

func (r *WebhookEndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a WorkOS webhook endpoint.",
		MarkdownDescription: `
Manages a WorkOS webhook endpoint in the current WorkOS environment.

WorkOS generates the webhook signing secret. The ` + "`secret`" + ` attribute is computed
and sensitive. WorkOS may only return this value when the endpoint is created; when
later read or update responses omit it, the provider preserves the prior state value.

Terraform state and plan files may contain the webhook signing secret. Use a secure
remote backend, restrict state access, and do not commit state or plan files.

## Example Usage

` + "```hcl" + `
resource "workos_webhook_endpoint" "app" {
  endpoint_url = "https://api.example.com/workos/webhook"
  events = [
    "user.created",
    "user.updated",
    "organization.created",
    "organization.updated",
  ]
  status = "enabled"
}
` + "```" + `

## Import

Webhook endpoints can be imported using their WorkOS ID:

` + "```shell" + `
terraform import workos_webhook_endpoint.example webhook_endpoint_01H00000000000000000000000
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "The unique identifier of the webhook endpoint.",
				MarkdownDescription: "The unique identifier of the webhook endpoint.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"endpoint_url": schema.StringAttribute{
				Description:         "The HTTPS URL WorkOS sends webhook events to.",
				MarkdownDescription: "The HTTPS URL WorkOS sends webhook events to.",
				Required:            true,
			},
			"events": schema.SetAttribute{
				Description:         "The complete set of event types delivered to the endpoint.",
				MarkdownDescription: "The complete set of event types delivered to the endpoint.",
				Required:            true,
				ElementType:         types.StringType,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"status": schema.StringAttribute{
				Description:         "Whether the endpoint is enabled or disabled.",
				MarkdownDescription: "Whether the endpoint is `enabled` or `disabled`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("enabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"secret": schema.StringAttribute{
				Description:         "The WorkOS-generated webhook signing secret.",
				MarkdownDescription: "The WorkOS-generated webhook signing secret. This value is sensitive and may only be returned by WorkOS at creation time.",
				Computed:            true,
				Sensitive:           true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Description:         "The timestamp when the webhook endpoint was created.",
				MarkdownDescription: "The timestamp when the webhook endpoint was created (RFC3339 format).",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				Description:         "The timestamp when the webhook endpoint was last updated.",
				MarkdownDescription: "The timestamp when the webhook endpoint was last updated (RFC3339 format).",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					useStateForUnknownIfConfigUnchanged{
						configAttributes: []path.Path{
							path.Root("endpoint_url"),
							path.Root("events"),
							path.Root("status"),
						},
					},
				},
			},
		},
	}
}

func (r *WebhookEndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c
}

func (r *WebhookEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WebhookEndpointResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	events, diags := stringSetFromTerraform(ctx, plan.Events)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, err := r.client.CreateWebhookEndpoint(ctx, &client.WebhookEndpointCreateRequest{
		EndpointURL: plan.EndpointURL.ValueString(),
		Events:      events,
		Status:      plan.Status.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Webhook Endpoint", "Could not create webhook endpoint: "+err.Error())
		return
	}

	webhookEndpointToState(ctx, &plan, endpoint, types.StringNull(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WebhookEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WebhookEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, err := r.client.GetWebhookEndpoint(ctx, state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			tflog.Info(ctx, "Webhook endpoint not found, removing from state", map[string]any{"id": state.ID.ValueString()})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Webhook Endpoint", "Could not read webhook endpoint: "+err.Error())
		return
	}

	webhookEndpointToState(ctx, &state, endpoint, state.Secret, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WebhookEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WebhookEndpointResourceModel
	var state WebhookEndpointResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	events, diags := stringSetFromTerraform(ctx, plan.Events)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, err := r.client.UpdateWebhookEndpoint(ctx, state.ID.ValueString(), &client.WebhookEndpointUpdateRequest{
		EndpointURL: plan.EndpointURL.ValueString(),
		Events:      events,
		Status:      plan.Status.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Webhook Endpoint", "Could not update webhook endpoint: "+err.Error())
		return
	}

	webhookEndpointToState(ctx, &plan, endpoint, state.Secret, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WebhookEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WebhookEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteWebhookEndpoint(ctx, state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			tflog.Info(ctx, "Webhook endpoint already deleted", map[string]any{"id": state.ID.ValueString()})
			return
		}
		resp.Diagnostics.AddError("Error Deleting Webhook Endpoint", "Could not delete webhook endpoint: "+err.Error())
	}
}

func (r *WebhookEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func stringSetFromTerraform(ctx context.Context, value types.Set) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	if value.IsNull() || value.IsUnknown() {
		return nil, diags
	}

	var values []string
	diags.Append(value.ElementsAs(ctx, &values, false)...)
	sort.Strings(values)
	return values, diags
}

func webhookEndpointToState(ctx context.Context, state *WebhookEndpointResourceModel, endpoint *client.WebhookEndpoint, priorSecret types.String, diags *diag.Diagnostics) {
	state.ID = types.StringValue(endpoint.ID)
	state.EndpointURL = types.StringValue(endpoint.EndpointURL)
	state.CreatedAt = timeString(endpoint.CreatedAt)
	state.UpdatedAt = timeString(endpoint.UpdatedAt)

	if endpoint.Status != "" {
		state.Status = types.StringValue(endpoint.Status)
	} else if state.Status.IsNull() || state.Status.IsUnknown() {
		state.Status = types.StringNull()
	}

	if endpoint.Secret != "" {
		state.Secret = types.StringValue(endpoint.Secret)
	} else if !priorSecret.IsNull() && !priorSecret.IsUnknown() {
		state.Secret = priorSecret
	} else {
		state.Secret = types.StringNull()
	}

	events, eventDiags := types.SetValueFrom(ctx, types.StringType, endpoint.Events)
	diags.Append(eventDiags...)
	state.Events = events
}
