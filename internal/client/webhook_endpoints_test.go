// Copyright (c) OSO DevOps
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const webhookEndpointFixture = `{
  "id": "webhook_endpoint_01JYQ5B9Q6ZP8K4R2T1V0X9ABC",
  "object": "webhook_endpoint",
  "endpoint_url": "https://api.example.com/workos/webhook",
  "events": ["user.created", "user.updated"],
  "status": "enabled",
  "secret": "whsec_test",
  "created_at": "2026-01-15T12:00:00.000Z",
  "updated_at": "2026-01-15T12:00:00.000Z"
}`

func TestWebhookEndpointsClientCreateUpdateAndDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			if r.URL.Path != "/webhook_endpoints" {
				t.Fatalf("expected /webhook_endpoints, got %s", r.URL.Path)
			}
			var body WebhookEndpointCreateRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("failed to decode create request body: %v", err)
			}
			if body.EndpointURL != "https://api.example.com/workos/webhook" || body.Status != "enabled" || len(body.Events) != 2 {
				t.Fatalf("unexpected create request body: %#v", body)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(webhookEndpointFixture))
		case http.MethodPatch:
			if r.URL.Path != "/webhook_endpoints/webhook_endpoint_01JYQ5B9Q6ZP8K4R2T1V0X9ABC" {
				t.Fatalf("unexpected update path: %s", r.URL.Path)
			}
			var body WebhookEndpointUpdateRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("failed to decode update request body: %v", err)
			}
			if body.EndpointURL != "https://api.example.com/workos/webhook" || body.Status != "disabled" || len(body.Events) != 1 {
				t.Fatalf("unexpected update request body: %#v", body)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(webhookEndpointFixture))
		case http.MethodDelete:
			if r.URL.Path != "/webhook_endpoints/webhook_endpoint_01JYQ5B9Q6ZP8K4R2T1V0X9ABC" {
				t.Fatalf("unexpected delete path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	}))
	defer server.Close()

	client, err := NewClient("sk_test", "", server.URL)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	endpoint, err := client.CreateWebhookEndpoint(context.Background(), &WebhookEndpointCreateRequest{
		EndpointURL: "https://api.example.com/workos/webhook",
		Events:      []string{"user.created", "user.updated"},
		Status:      "enabled",
	})
	if err != nil {
		t.Fatalf("CreateWebhookEndpoint returned error: %v", err)
	}
	if endpoint.ID != "webhook_endpoint_01JYQ5B9Q6ZP8K4R2T1V0X9ABC" || endpoint.Secret != "whsec_test" {
		t.Fatalf("unexpected webhook endpoint response: %#v", endpoint)
	}

	if _, err := client.UpdateWebhookEndpoint(context.Background(), endpoint.ID, &WebhookEndpointUpdateRequest{
		EndpointURL: "https://api.example.com/workos/webhook",
		Events:      []string{"user.created"},
		Status:      "disabled",
	}); err != nil {
		t.Fatalf("UpdateWebhookEndpoint returned error: %v", err)
	}

	if err := client.DeleteWebhookEndpoint(context.Background(), endpoint.ID); err != nil {
		t.Fatalf("DeleteWebhookEndpoint returned error: %v", err)
	}
}

func TestWebhookEndpointsClientListPaginationAndGet(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/webhook_endpoints" {
			t.Fatalf("expected /webhook_endpoints, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("limit") != "100" {
			t.Fatalf("expected limit=100, got %s", r.URL.Query().Get("limit"))
		}

		w.Header().Set("Content-Type", "application/json")
		if requests == 1 {
			if r.URL.Query().Get("after") != "" {
				t.Fatalf("did not expect after on first page, got %s", r.URL.Query().Get("after"))
			}
			_, _ = w.Write([]byte(`{"data": [], "list_metadata": {"after": "cursor_1"}}`))
			return
		}
		if r.URL.Query().Get("after") != "cursor_1" {
			t.Fatalf("expected after=cursor_1, got %s", r.URL.Query().Get("after"))
		}
		_, _ = w.Write([]byte(`{"data": [` + webhookEndpointFixture + `], "list_metadata": {}}`))
	}))
	defer server.Close()

	client, err := NewClient("sk_test", "", server.URL)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	endpoint, err := client.GetWebhookEndpoint(context.Background(), "webhook_endpoint_01JYQ5B9Q6ZP8K4R2T1V0X9ABC")
	if err != nil {
		t.Fatalf("GetWebhookEndpoint returned error: %v", err)
	}
	if endpoint.EndpointURL != "https://api.example.com/workos/webhook" {
		t.Fatalf("unexpected webhook endpoint: %#v", endpoint)
	}
	if requests != 2 {
		t.Fatalf("expected 2 requests, got %d", requests)
	}
}
