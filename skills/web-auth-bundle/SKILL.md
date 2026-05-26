---
name: web-auth-bundle
description: Bundle skill for authorized web and API authentication assessment. Use when requests involve login flows, sessions, cookies, bearer tokens, refresh tokens, MFA, CSRF boundaries, or mixed browser/API auth and Codex should coordinate multiple auth-focused skills.
metadata:
  version: "1.0.0"
  tags:
    - bundle
    - web
    - api
    - auth
    - session
    - jwt
  triggers:
    - login flow
    - session cookie
    - bearer token
    - refresh token
    - mfa
    - csrf
    - web auth
  target_types:
    - web
    - api
  bundle_of:
    - auth-session-review
    - token-lifecycle-review
    - secret-exposure-review
    - security-headers-review
  depends_on:
    - openapi-contract-review
  recommended_tools:
    - http-framework-test
    - jwt-analyzer
    - nuclei
  role_hints:
    - Web application
    - API security
  stages:
    - recon
    - verify
    - report
  autoload_priority: 22
---

# Web Auth Bundle

Use this bundle to coordinate the listed member skills for an auth-heavy assessment.

## Workflow

1. Map the authentication surface first: login, logout, refresh, password reset, MFA, admin switches, and any browser/API token crossover.
2. Load only the member skills that match the concrete issue instead of pulling every auth-related skill blindly.
3. Prioritize evidence that distinguishes browser session state, API bearer state, and backend authorization state.

## Load Order

- Start with `auth-session-review` for login, logout, cookie, and CSRF boundaries.
- Load `token-lifecycle-review` when refresh, expiry, rotation, revocation, or replay behavior matters.
- Load `secret-exposure-review` when tokens, keys, or env leaks appear in responses, docs, or frontend assets.
- Load `security-headers-review` when browser trust boundaries and cookie/header hardening matter.
- Load `openapi-contract-review` when the API contract may be understating auth or schema constraints.

## Output

- Auth surface inventory
- Member skills actually loaded
- Key evidence by state transition
- Reproducible findings and fixes
