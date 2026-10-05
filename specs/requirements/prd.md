# greeter — PRD

## Problem Statement

Teams building services on this platform need a small, working example of a Go HTTP service built to the organization's own conventions (as documented in `app-factory-kaj/e2e-reference`). Without a minimal, concrete reference, each team re-derives basic scaffolding — routing, JSON response shape, error handling, configuration — from scratch, leading to inconsistent services across the organization.

## Solution

Greeter is a small, standalone Go HTTP service exposing a single endpoint: `GET /hello`, which takes a `name` parameter and returns a JSON greeting addressed to that name. It is built to the organization's Go service conventions so it can serve as a working reference other services can be modeled on.

## Actors

- **Client** — any caller, human or system, that invokes the greeter service's HTTP API to obtain a greeting.

## User Stories

1. As a Client, I want to call `GET /hello` with a `name` value, so that I receive a JSON greeting addressed to that name.
2. As a Client, I want to call `GET /hello` without a `name` value, so that I still receive a sensible default greeting rather than an error. *assumed*

## Product Decisions

- No sign-in: greeter exposes a public, unauthenticated API — it is a minimal reference service with no user accounts or sensitive data. *assumed*
- Response shape: a JSON body with a `message` field carrying the greeting text (e.g. `{"message": "Hello, <name>!"}`), matching the stable-JSON-shape convention the organization's reference documents. *assumed*
- Errors are returned as JSON (`{"error": "..."}`) with an appropriate HTTP status code, per the organization's Go service conventions.
- Implementation follows the organization's Go service conventions documented in `app-factory-kaj/e2e-reference` (project layout, error wrapping, configuration via environment variables with sensible defaults, standard-library `net/http`).

## Out of Scope

- Persistence or storage of any kind — greeter is stateless.
- Authentication, authorization, or per-user behavior.
- Rate limiting, quotas, or multi-tenancy.
- Any endpoint beyond `GET /hello`.
- Internationalization or localization of the greeting text.

## Open Questions

*(none outstanding)*

## Further Notes

This project's idea explicitly directs the implementation to follow the conventions in `app-factory-kaj/e2e-reference`; those are engineering-level conventions (project layout, error handling, testing, configuration) and are carried forward to the design rather than restated here.