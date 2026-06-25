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

output "app_webhook_endpoint_id" {
  value = workos_webhook_endpoint.app.id
}
