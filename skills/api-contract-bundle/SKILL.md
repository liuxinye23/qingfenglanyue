---
name: api-contract-bundle
description: Bundle skill for authorized API contract and trust-boundary review. Use when requests involve OpenAPI or Swagger specs, JSON APIs, schema drift, auth declarations, token handling, TLS assumptions, or dependency-driven API risk and Codex should coordinate the relevant review skills.
metadata:
  version: "1.0.0"
  tags:
    - bundle
    - api
    - openapi
    - swagger
    - schema
    - contract
  triggers:
    - openapi
    - swagger
    - api schema
    - contract drift
    - auth declaration
    - tls review
  target_types:
    - api
    - web
  bundle_of:
    - openapi-contract-review
    - token-lifecycle-review
    - tls-configuration-review
    - dependency-risk-review
    - secret-exposure-review
  recommended_tools:
    - api-schema-analyzer
    - http-framework-test
    - nuclei
  role_hints:
    - API security
    - Backend review
  stages:
    - recon
    - verify
    - report
  autoload_priority: 21
---

# API Contract Bundle

Use this bundle when the request is fundamentally about API behavior, contract accuracy, or trust-boundary consistency.

## Workflow

1. Compare the declared contract with the observed runtime behavior.
2. Pull in only the member skills required by the specific trust boundary in scope.
3. Separate contract/documentation defects from exploitable runtime defects.

## Load Order

- Start with `openapi-contract-review` for endpoint, schema, and auth declaration drift.
- Load `token-lifecycle-review` when the API relies on tokens or session-like refresh flows.
- Load `tls-configuration-review` when transport assumptions or certificate posture matter.
- Load `dependency-risk-review` when gateway, parser, or framework components may affect request handling.
- Load `secret-exposure-review` when keys, internal headers, or backend config bleed into responses or docs.

## Output

- Contracted vs observed behavior
- High-risk endpoints
- Trust-boundary defects
- Priority fixes
