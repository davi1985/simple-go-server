---
# Simple Go Server 🚀

A lightweight Go web application built to practice idiomatic project architecture, standard HTTP patterns, and native unit testing practices.
---

## Key Patterns & Learnings

This project implements core conventions and best practices widely adopted by the Go community:

- **Idiomatic Project Layout (`cmd/` & `internal/`)**:
- Isolates the executable entry point inside `cmd/server/`.
- Encapsulates private application logic inside `internal/handler/`, preventing unauthorized imports from external packages.

- **Standard HTTP Handlers**:
- Uses Go's native `net/http` package to handle routing, requests, and responses.
- Separates concerns between serving static assets and handling dynamic routes (`/hello`, `/form`).

- **Framework-Free Unit Testing**:
- Utilizes native `testing` and `net/http/httptest.ResponseRecorder` for mocking HTTP requests without third-party dependencies.
- Applies **Black-box Testing** (`package handler_test`) to test package behaviors through their public interface.
- Implements **Table-Driven Tests** to cleanly validate multiple scenarios and HTTP methods with minimal code repetition.

---

## Project Structure

```text
simple-go-server/
├── cmd/
│   └── server/
│       └── main.go          # Entry point & server initialization
├── internal/
│   └── handler/
│       ├── handler.go       # HTTP Handlers implementation
│       └── handler_test.go  # Unit tests for HTTP handlers
├── static/                  # Static web files (HTML/Forms)
└── go.mod                   # Go module definition

```

---

## How to Run

### 1. Start the Server

From the root directory, run:

```bash
go run ./cmd/server/main.go

```

The server will start at `http://localhost:8080`.

### 2. Run Tests

To execute the unit test suite:

```bash
go test ./...

```

For verbose output detailing each test case execution:

```bash
go test -v ./...

```
