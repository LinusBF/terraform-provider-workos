// Copyright (c) OSO DevOps
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/osodevops/terraform-provider-workos/internal/client"
)

func TestRedirectURIResourceSchema(t *testing.T) {
	resp := &resource.SchemaResponse{}
	NewRedirectURIResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected schema diagnostics: %v", resp.Diagnostics)
	}

	uri, ok := resp.Schema.Attributes["uri"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("expected uri to be a string attribute")
	}
	if !uri.Required {
		t.Fatal("expected uri to be required")
	}
	if len(uri.PlanModifiers) == 0 {
		t.Fatal("expected uri to require replacement when changed")
	}
}

func TestCORSOriginResourceSchema(t *testing.T) {
	resp := &resource.SchemaResponse{}
	NewCORSOriginResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected schema diagnostics: %v", resp.Diagnostics)
	}

	origin, ok := resp.Schema.Attributes["origin"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("expected origin to be a string attribute")
	}
	if !origin.Required {
		t.Fatal("expected origin to be required")
	}
	if len(origin.PlanModifiers) == 0 {
		t.Fatal("expected origin to require replacement when changed")
	}
}

func TestWebhookEndpointResourceSchema(t *testing.T) {
	resp := &resource.SchemaResponse{}
	NewWebhookEndpointResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected schema diagnostics: %v", resp.Diagnostics)
	}

	events, ok := resp.Schema.Attributes["events"].(schema.SetAttribute)
	if !ok {
		t.Fatalf("expected events to be a set attribute")
	}
	if !events.Required {
		t.Fatal("expected events to be required")
	}
	if !events.ElementType.Equal(types.StringType) {
		t.Fatalf("expected events element type to be string, got %s", events.ElementType)
	}

	secret, ok := resp.Schema.Attributes["secret"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("expected secret to be a string attribute")
	}
	if !secret.Computed || !secret.Sensitive {
		t.Fatal("expected secret to be computed and sensitive")
	}
}

func TestWebhookEndpointToStatePreservesPriorSecret(t *testing.T) {
	state := &WebhookEndpointResourceModel{}
	var diagnostics diag.Diagnostics

	webhookEndpointToState(context.Background(), state, &client.WebhookEndpoint{
		ID:          "webhook_endpoint_123",
		EndpointURL: "https://api.example.com/workos/webhook",
		Events:      []string{"user.updated", "user.created"},
		Status:      "enabled",
		CreatedAt:   time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC),
	}, types.StringValue("whsec_previous"), &diagnostics)

	if diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diagnostics)
	}
	if state.Secret.ValueString() != "whsec_previous" {
		t.Fatalf("expected prior secret to be preserved, got %q", state.Secret.ValueString())
	}
	if state.Events.IsNull() || state.Events.IsUnknown() {
		t.Fatal("expected events to be set")
	}
}
