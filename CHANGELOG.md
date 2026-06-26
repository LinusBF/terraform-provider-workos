# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.4.0](https://github.com/LinusBF/terraform-provider-workos/compare/v2.3.1...v2.4.0) (2026-06-26)


### Features

* add AuthKit infrastructure resources ([2996357](https://github.com/LinusBF/terraform-provider-workos/commit/29963578721da9667c071bbb49dab9e102cd186a))
* add external_id and metadata to organization resource ([4bde9d4](https://github.com/LinusBF/terraform-provider-workos/commit/4bde9d43ca0edf89c938f20797bc74f6e21c352b))
* add external_id and metadata to organization resource ([beb9dbd](https://github.com/LinusBF/terraform-provider-workos/commit/beb9dbdac487ab15e502dd04cedb227d2ae45ff4))
* add external_id lookup to organization data source ([d9f645b](https://github.com/LinusBF/terraform-provider-workos/commit/d9f645b0fced2b2c3db295b5121e45c645309a8e))
* add permission resource, data source, and organization role permission resource ([d9ded58](https://github.com/LinusBF/terraform-provider-workos/commit/d9ded58b3a1a4890872eafcf2109effda3cc6c7a))
* add permission resource, data source, and organization role permission resource ([10fd884](https://github.com/LinusBF/terraform-provider-workos/commit/10fd884c8de8e31aef2c635ad721392bf9dc3da3))
* add WorkOS application and authorization resources ([f5b0f89](https://github.com/LinusBF/terraform-provider-workos/commit/f5b0f8924fc126281dcd0f712bd73703ac2ecd92))
* add WorkOS application and authorization resources ([f2548f4](https://github.com/LinusBF/terraform-provider-workos/commit/f2548f42af972c20de4b27fad6e9b3e35f462548))
* add WorkOS environment role support ([2f9a854](https://github.com/LinusBF/terraform-provider-workos/commit/2f9a85424f55f82940e6e899b88b5ed469e7413b))
* add WorkOS environment role support ([068d22b](https://github.com/LinusBF/terraform-provider-workos/commit/068d22be0370a44c9f1bb134f4f4bfa9bff1a41d))
* add workos_organization_role resource and data source ([bca489a](https://github.com/LinusBF/terraform-provider-workos/commit/bca489a7bf76712d08dc39866c7282623c5e30ad))
* add workos_organization_role resource and data source ([394c45d](https://github.com/LinusBF/terraform-provider-workos/commit/394c45d38fd874252f8d8afe30c6d87141c362df))
* implement missing fields on user objects ([ba1daa4](https://github.com/LinusBF/terraform-provider-workos/commit/ba1daa4d902592b6e6efd46061ef7f9a41a37281))
* implement the missing fields on user objects ([778fb40](https://github.com/LinusBF/terraform-provider-workos/commit/778fb40901085c5ecc762dea707129b9710b1b6b))
* Initial release of WorkOS Terraform Provider v1.0.0 ([0e50a3b](https://github.com/LinusBF/terraform-provider-workos/commit/0e50a3bec5992d443f9a9b5092e7087dbb8a432c))


### Bug Fixes

* add IsUnknown() checks for Optional+Computed fields ([774d8e3](https://github.com/LinusBF/terraform-provider-workos/commit/774d8e323cc0d0856c75487c6d45e28a768e63b1))
* add IsUnknown() checks for Optional+Computed fields ([612617b](https://github.com/LinusBF/terraform-provider-workos/commit/612617bd5cc9b9e62bb65c87ec7a4165abd9aca9))
* align Terraform Registry release assets ([139697b](https://github.com/LinusBF/terraform-provider-workos/commit/139697baf1b2c37f14cbe11f1b26d4e013dba31e))
* **ci:** Add golangci-lint config and tools dependency ([d5bd48d](https://github.com/LinusBF/terraform-provider-workos/commit/d5bd48d3485eaad04d088c7f635a030c667ceb63))
* **ci:** Add Terraform setup and fix goreleaser deprecation ([e8cdc5c](https://github.com/LinusBF/terraform-provider-workos/commit/e8cdc5cf82d326c30cc59a0499a395b013877d78))
* **ci:** Update GoReleaser action to v2 for version 2 config ([7085010](https://github.com/LinusBF/terraform-provider-workos/commit/7085010d57e871153edae6586b3686da4322f005))
* correctly parse nested role object from membership API response ([9d9476c](https://github.com/LinusBF/terraform-provider-workos/commit/9d9476c71e8817fcf9a9ccf6adf9826162e8dd83))
* include registry manifest in release checksums ([80437d6](https://github.com/LinusBF/terraform-provider-workos/commit/80437d6cd93d022337195c49d01ae8deb07132c5))
* pin GitHub Actions to immutable commit SHAs ([2410473](https://github.com/LinusBF/terraform-provider-workos/commit/241047357041aee2d8e677905dfe03708646225b))
* prefix organization role slugs with 'org-' per WorkOS API requirement ([1e700d0](https://github.com/LinusBF/terraform-provider-workos/commit/1e700d098989a3cf1a5727cd57246bdbc490fb5f))
* prevent updated_at drift and allow clearing description on permission resource ([0df2991](https://github.com/LinusBF/terraform-provider-workos/commit/0df2991c8d4ec86bcc174b39518e1da4f82b18cf))
* prevent updated_at drift on re-apply ([36a1b9e](https://github.com/LinusBF/terraform-provider-workos/commit/36a1b9ea3b3e0b500dd563f8c8848986013cca66))
* prevent updated_at drift on re-apply of unchanged state ([1ad7f91](https://github.com/LinusBF/terraform-provider-workos/commit/1ad7f9116917cccb880521d217b5429e52bf1627)), closes [#5](https://github.com/LinusBF/terraform-provider-workos/issues/5)
* regenerate docs to match go generate output ([ea88f08](https://github.com/LinusBF/terraform-provider-workos/commit/ea88f087e52f0630509735cec14064b5186dffb2))
* remove domain uniqueness constraint from organization resource ([27abbbd](https://github.com/LinusBF/terraform-provider-workos/commit/27abbbd0401705b54482dfa08350bd89fdecd669))
* remove unsupported resources and fix acceptance test failures ([8555547](https://github.com/LinusBF/terraform-provider-workos/commit/85555472f4102b5eb4a0e8f832a963bd92281157))
* remove unsupported resources and fix acceptance test failures ([94dcad4](https://github.com/LinusBF/terraform-provider-workos/commit/94dcad47f7ba7d1682f0a8763ff13b05379c1e4d))
* resolve updated_at drift on update and email_verified reset ([3f93787](https://github.com/LinusBF/terraform-provider-workos/commit/3f937874700ecbf6fe6d1fb2cb3c15e9a20baeed))
* send null for removed metadata keys on updates ([988a8ac](https://github.com/LinusBF/terraform-provider-workos/commit/988a8ac94204397ee11828414e3fb6a864bacf4a))
* send null for removed metadata keys on user and organization updates ([aad31e2](https://github.com/LinusBF/terraform-provider-workos/commit/aad31e2abaa0dcf688618e172a47283302e284ff)), closes [#18](https://github.com/LinusBF/terraform-provider-workos/issues/18)
* validate domain uniqueness across organizations ([0658b1f](https://github.com/LinusBF/terraform-provider-workos/commit/0658b1fb8868e097358d77c11887421f8aa883a0))
* validate domain uniqueness across organizations ([14e5d2e](https://github.com/LinusBF/terraform-provider-workos/commit/14e5d2ebcf6d35b6c85286ff7fac8c075523f68e)), closes [#8](https://github.com/LinusBF/terraform-provider-workos/issues/8)

## [Unreleased]

### Features

* add environment role resource and data source for WorkOS Authorization roles
* add AuthKit redirect URI, AuthKit CORS origin, and webhook endpoint resources

## [2.3.1](https://github.com/osodevops/terraform-provider-workos/compare/v2.3.0...v2.3.1) (2026-06-04)


### Bug Fixes

* align Terraform Registry release assets ([139697b](https://github.com/osodevops/terraform-provider-workos/commit/139697baf1b2c37f14cbe11f1b26d4e013dba31e))
* include registry manifest in release checksums ([80437d6](https://github.com/osodevops/terraform-provider-workos/commit/80437d6cd93d022337195c49d01ae8deb07132c5))

## [2.3.0](https://github.com/osodevops/terraform-provider-workos/compare/v2.2.1...v2.3.0) (2026-06-04)


### Features

* add WorkOS application and authorization resources ([f5b0f89](https://github.com/osodevops/terraform-provider-workos/commit/f5b0f8924fc126281dcd0f712bd73703ac2ecd92))
* add WorkOS application and authorization resources ([f2548f4](https://github.com/osodevops/terraform-provider-workos/commit/f2548f42af972c20de4b27fad6e9b3e35f462548))

## [2.2.1](https://github.com/osodevops/terraform-provider-workos/compare/v2.2.0...v2.2.1) (2026-04-09)


### Bug Fixes

* add IsUnknown() checks for Optional+Computed fields ([774d8e3](https://github.com/osodevops/terraform-provider-workos/commit/774d8e323cc0d0856c75487c6d45e28a768e63b1))
* add IsUnknown() checks for Optional+Computed fields ([612617b](https://github.com/osodevops/terraform-provider-workos/commit/612617bd5cc9b9e62bb65c87ec7a4165abd9aca9))
* correctly parse nested role object from membership API response ([9d9476c](https://github.com/osodevops/terraform-provider-workos/commit/9d9476c71e8817fcf9a9ccf6adf9826162e8dd83))
* pin GitHub Actions to immutable commit SHAs ([2410473](https://github.com/osodevops/terraform-provider-workos/commit/241047357041aee2d8e677905dfe03708646225b))
* remove domain uniqueness constraint from organization resource ([27abbbd](https://github.com/osodevops/terraform-provider-workos/commit/27abbbd0401705b54482dfa08350bd89fdecd669))

## [2.2.0](https://github.com/osodevops/terraform-provider-workos/compare/v2.1.0...v2.2.0) (2026-03-25)


### Features

* implement missing fields on user objects ([ba1daa4](https://github.com/osodevops/terraform-provider-workos/commit/ba1daa4d902592b6e6efd46061ef7f9a41a37281))
* implement the missing fields on user objects ([778fb40](https://github.com/osodevops/terraform-provider-workos/commit/778fb40901085c5ecc762dea707129b9710b1b6b))

## [2.1.0](https://github.com/osodevops/terraform-provider-workos/compare/v2.0.0...v2.1.0) (2026-03-13)


### Features

* add external_id and metadata to organization resource ([4bde9d4](https://github.com/osodevops/terraform-provider-workos/commit/4bde9d43ca0edf89c938f20797bc74f6e21c352b))
* add external_id and metadata to organization resource ([beb9dbd](https://github.com/osodevops/terraform-provider-workos/commit/beb9dbdac487ab15e502dd04cedb227d2ae45ff4))
* add external_id lookup to organization data source ([d9f645b](https://github.com/osodevops/terraform-provider-workos/commit/d9f645b0fced2b2c3db295b5121e45c645309a8e))
* add permission resource, data source, and organization role permission resource ([d9ded58](https://github.com/osodevops/terraform-provider-workos/commit/d9ded58b3a1a4890872eafcf2109effda3cc6c7a))
* add permission resource, data source, and organization role permission resource ([10fd884](https://github.com/osodevops/terraform-provider-workos/commit/10fd884c8de8e31aef2c635ad721392bf9dc3da3))


### Bug Fixes

* prevent updated_at drift and allow clearing description on permission resource ([0df2991](https://github.com/osodevops/terraform-provider-workos/commit/0df2991c8d4ec86bcc174b39518e1da4f82b18cf))
* prevent updated_at drift on re-apply ([36a1b9e](https://github.com/osodevops/terraform-provider-workos/commit/36a1b9ea3b3e0b500dd563f8c8848986013cca66))
* prevent updated_at drift on re-apply of unchanged state ([1ad7f91](https://github.com/osodevops/terraform-provider-workos/commit/1ad7f9116917cccb880521d217b5429e52bf1627)), closes [#5](https://github.com/osodevops/terraform-provider-workos/issues/5)
* resolve updated_at drift on update and email_verified reset ([3f93787](https://github.com/osodevops/terraform-provider-workos/commit/3f937874700ecbf6fe6d1fb2cb3c15e9a20baeed))
* validate domain uniqueness across organizations ([0658b1f](https://github.com/osodevops/terraform-provider-workos/commit/0658b1fb8868e097358d77c11887421f8aa883a0))
* validate domain uniqueness across organizations ([14e5d2e](https://github.com/osodevops/terraform-provider-workos/commit/14e5d2ebcf6d35b6c85286ff7fac8c075523f68e)), closes [#8](https://github.com/osodevops/terraform-provider-workos/issues/8)

## [2.0.0] - 2026-02-24

### Added
- `workos_organization_role` resource - Manage organization authorization roles
- `workos_organization_role` data source - Look up organization roles by slug or ID

### Removed
- **BREAKING:** `workos_connection` resource - WorkOS API does not support creating/updating connections via API; use the Dashboard instead. The read-only data source is still available.
- **BREAKING:** `workos_directory` resource - WorkOS API does not support creating/updating directories via API; use the Dashboard instead. The read-only data source is still available.
- **BREAKING:** `workos_webhook` resource - WorkOS has no public webhook management API; use the Dashboard instead.
- **BREAKING:** `allow_profiles_outside_organization` attribute on `workos_organization` resource and data source - WorkOS API no longer accepts this parameter.

### Fixed
- `workos_user` resource: `email_verified` is now always sent on updates, preventing drift when email changes reset verification status
- `workos_organization_membership` resource: `role_slug` is preserved from plan/state when the API omits it from responses
- `workos_user` data source tests: replaced hardcoded placeholder IDs with dynamically created resources
- `workos_organization_role` resource: slug is now prefixed with `org-` per WorkOS API requirement

## [1.0.0] - 2026-02-01

### Added

#### Provider
- Provider configuration with `api_key`, `client_id`, and `base_url` attributes
- Environment variable support: `WORKOS_API_KEY`, `WORKOS_CLIENT_ID`, `WORKOS_BASE_URL`
- Rate limiting with exponential backoff and Retry-After header support
- Comprehensive error handling with typed errors

#### Resources
- `workos_organization` - Manage WorkOS organizations
  - Full CRUD operations
  - Domain management
  - Import support

- `workos_user` - Manage AuthKit users
  - Email and name management
  - Password and password hash support for authentication
  - Email verification status
  - Import support

- `workos_organization_membership` - Manage user-organization associations
  - User and organization linking
  - Role assignment support
  - Import support

#### Data Sources
- `workos_organization` - Look up organizations by ID or domain
- `workos_connection` - Look up SSO connections by ID or organization/type (read-only)
- `workos_directory` - Look up directories by ID or organization (read-only)
- `workos_directory_user` - Look up directory-synced users by ID or email
- `workos_directory_group` - Look up directory-synced groups by ID or name
- `workos_user` - Look up AuthKit users by ID or email

#### Documentation
- Auto-generated documentation using terraform-plugin-docs
- Comprehensive examples for all resources and data sources
- Schema descriptions with Markdown support

### Security
- API keys marked as sensitive and never logged
- User passwords marked as sensitive (write-only)
