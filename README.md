# cloudrift

A small REST API in Go, built as a portfolio demo of a cleanly layered service: handler → service → repository, with dependency injection and explicit error handling.

## Structure

```
api/                  # Entry point: wiring, HTTP server
internal/
├── user/             # Users: model, service, in-memory repo, handlers
├── order/            # Orders: model, service, in-memory repo, handlers
├── product/          # Products (work in progress, not yet exposed)
└── platform/httpx/   # Shared JSON response helpers
```

Each domain depends only on a `Repository` interface, so the in-memory storage can be swapped (e.g. for PostgreSQL) without touching business logic.

## Endpoints

| Method | Path           | Description     |
|--------|----------------|-----------------|
| GET    | `/users`       | List users      |
| GET    | `/users/{id}`  | Get a user      |
| POST   | `/users`       | Create a user   |
| POST   | `/orders`      | Place an order  |
| GET    | `/orders/{id}` | Get an order    |

Responses are wrapped as `{"data": ...}`, errors as `{"error": "..."}` with `400`, `404` or `500`.

```bash
curl -X POST http://localhost:8080/users -d '{"name": "Ada Lovelace"}'
```

## Run

Requires Go 1.27+.

```bash
go run ./api    # http://localhost:8080
```

---

Work in progress — not intended for production use.
