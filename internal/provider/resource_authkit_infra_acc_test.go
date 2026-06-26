// Copyright (c) OSO DevOps
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRedirectURIResource_Basic(t *testing.T) {
	uri := fmt.Sprintf("https://example.com/tf-acc-%d/callback", time.Now().UnixNano())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRedirectURIResourceConfig(uri),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("workos_redirect_uri.test", "uri", uri),
					resource.TestCheckResourceAttrSet("workos_redirect_uri.test", "id"),
					resource.TestCheckResourceAttrSet("workos_redirect_uri.test", "created_at"),
					resource.TestCheckResourceAttrSet("workos_redirect_uri.test", "updated_at"),
				),
			},
			{
				ResourceName:      "workos_redirect_uri.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCORSOriginResource_Basic(t *testing.T) {
	origin := fmt.Sprintf("https://tf-acc-%d.example.com", time.Now().UnixNano())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCORSOriginResourceConfig(origin),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("workos_cors_origin.test", "origin", origin),
					resource.TestCheckResourceAttrSet("workos_cors_origin.test", "id"),
					resource.TestCheckResourceAttrSet("workos_cors_origin.test", "created_at"),
					resource.TestCheckResourceAttrSet("workos_cors_origin.test", "updated_at"),
				),
			},
			{
				ResourceName:      "workos_cors_origin.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccWebhookEndpointResource_Basic(t *testing.T) {
	url := fmt.Sprintf("https://example.com/tf-acc-%d/workos/webhook", time.Now().UnixNano())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWebhookEndpointResourceConfig(url, "enabled", `"user.created", "user.updated"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("workos_webhook_endpoint.test", "endpoint_url", url),
					resource.TestCheckResourceAttr("workos_webhook_endpoint.test", "status", "enabled"),
					resource.TestCheckResourceAttr("workos_webhook_endpoint.test", "events.#", "2"),
					resource.TestCheckResourceAttrSet("workos_webhook_endpoint.test", "id"),
					resource.TestCheckResourceAttrSet("workos_webhook_endpoint.test", "created_at"),
					resource.TestCheckResourceAttrSet("workos_webhook_endpoint.test", "updated_at"),
				),
			},
			{
				Config: testAccWebhookEndpointResourceConfig(url, "disabled", `"organization.created", "user.created"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("workos_webhook_endpoint.test", "endpoint_url", url),
					resource.TestCheckResourceAttr("workos_webhook_endpoint.test", "status", "disabled"),
					resource.TestCheckResourceAttr("workos_webhook_endpoint.test", "events.#", "2"),
				),
			},
			{
				ResourceName:            "workos_webhook_endpoint.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret"},
			},
		},
	})
}

func testAccRedirectURIResourceConfig(uri string) string {
	return fmt.Sprintf(`
resource "workos_redirect_uri" "test" {
  uri = %[1]q
}
`, uri)
}

func testAccCORSOriginResourceConfig(origin string) string {
	return fmt.Sprintf(`
resource "workos_cors_origin" "test" {
  origin = %[1]q
}
`, origin)
}

func testAccWebhookEndpointResourceConfig(url, status, events string) string {
	return fmt.Sprintf(`
resource "workos_webhook_endpoint" "test" {
  endpoint_url = %[1]q
  events       = [%[3]s]
  status       = %[2]q
}
`, url, status, events)
}
