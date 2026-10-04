# cloudrift

A small REST API written in Go, built as a demo project to show how a production-grade API service is structured: clean layering, dependency injection, explicit error handling and testable components.

## Architecture

```
cloudrift/
├── api/              # Entry point: wiring, HTTP server, routing
└── internal/
    └── user/         # User domain
        ├── user.go       # Domain model, request types, domain errors
        ├── service.go    # Business logic and validation
        ├── memory.go     # In-memory repository (thread-safe)
        └── handler.go    # HTTP handlers
```

The code follows a layered design:

- **Handler**: translates HTTP requests and responses, maps domain errors to status codes.
- **Service**: holds the business rules and depends only on a `Repository` interface.
- **Repository**: stores data. The in-memory implementation can be replaced (for example with PostgreSQL) without touching the service layer.

## Endpoints

| Method | Path          | Description        | Success |
|--------|---------------|--------------------|---------|
| GET    | `/users`      | List all users     | `200`   |
| GET    | `/users/{id}` | Get a single user  | `200`   |
| POST   | `/users`      | Create a new user  | `201`   |

Example:

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Ada Lovelace"}'
```

```json
{ "id": "3f1c9a2e-...", "name": "Ada Lovelace" }
```

Errors return an appropriate status code (`400` for invalid input, `404` if a user does not exist).

## Getting started

Requirements: Go 1.27+

```bash
go run ./api
```

The server listens on `http://localhost:8080`.

Run the tests:

```bash
go test ./...
```

## Status

Work in progress. This project is meant as a portfolio demo, not for production use.
