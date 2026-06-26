// Copyright (c) OSO DevOps
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// WebhookEndpoint represents a WorkOS webhook endpoint.
type WebhookEndpoint struct {
	ID          string    `json:"id"`
	Object      string    `json:"object"`
	EndpointURL string    `json:"endpoint_url"`
	Events      []string  `json:"events"`
	Status      string    `json:"status"`
	Secret      string    `json:"secret,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// WebhookEndpointCreateRequest represents the request to create a webhook endpoint.
type WebhookEndpointCreateRequest struct {
	EndpointURL string   `json:"endpoint_url"`
	Events      []string `json:"events"`
	Status      string   `json:"status,omitempty"`
}

// WebhookEndpointUpdateRequest represents the request to update a webhook endpoint.
type WebhookEndpointUpdateRequest struct {
	EndpointURL string   `json:"endpoint_url,omitempty"`
	Events      []string `json:"events,omitempty"`
	Status      string   `json:"status,omitempty"`
}

// WebhookEndpointListResponse represents the response from listing webhook endpoints.
type WebhookEndpointListResponse struct {
	Data         []WebhookEndpoint `json:"data"`
	ListMetadata ListMetadata      `json:"list_metadata"`
}

// CreateWebhookEndpoint creates a webhook endpoint.
func (c *Client) CreateWebhookEndpoint(ctx context.Context, req *WebhookEndpointCreateRequest) (*WebhookEndpoint, error) {
	var endpoint WebhookEndpoint
	err := c.Post(ctx, "/webhook_endpoints", req, &endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook endpoint: %w", err)
	}
	return &endpoint, nil
}

// ListWebhookEndpoints lists webhook endpoints.
func (c *Client) ListWebhookEndpoints(ctx context.Context) (*WebhookEndpointListResponse, error) {
	var all WebhookEndpointListResponse
	params := url.Values{}
	applyDefaultPagination(params)

	for {
		var page WebhookEndpointListResponse
		err := c.Get(ctx, pathWithQuery("/webhook_endpoints", params), &page)
		if err != nil {
			return nil, fmt.Errorf("failed to list webhook endpoints: %w", err)
		}

		all.Data = append(all.Data, page.Data...)
		all.ListMetadata = page.ListMetadata
		if page.ListMetadata.After == "" {
			break
		}
		params.Set("after", page.ListMetadata.After)
	}

	return &all, nil
}

// GetWebhookEndpoint finds a webhook endpoint by ID using the list endpoint.
func (c *Client) GetWebhookEndpoint(ctx context.Context, id string) (*WebhookEndpoint, error) {
	resp, err := c.ListWebhookEndpoints(ctx)
	if err != nil {
		return nil, err
	}

	for _, endpoint := range resp.Data {
		if endpoint.ID == id {
			return &endpoint, nil
		}
	}

	return nil, &APIError{
		StatusCode: 404,
		Message:    fmt.Sprintf("no webhook endpoint found with ID: %s", id),
	}
}

// UpdateWebhookEndpoint updates a webhook endpoint.
func (c *Client) UpdateWebhookEndpoint(ctx context.Context, id string, req *WebhookEndpointUpdateRequest) (*WebhookEndpoint, error) {
	var endpoint WebhookEndpoint
	err := c.Patch(ctx, "/webhook_endpoints/"+url.PathEscape(id), req, &endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to update webhook endpoint: %w", err)
	}
	return &endpoint, nil
}

// DeleteWebhookEndpoint deletes a webhook endpoint.
func (c *Client) DeleteWebhookEndpoint(ctx context.Context, id string) error {
	err := c.Delete(ctx, "/webhook_endpoints/"+url.PathEscape(id))
	if err != nil {
		return fmt.Errorf("failed to delete webhook endpoint: %w", err)
	}
	return nil
}
