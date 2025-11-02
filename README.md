# Url Shortener

                        ┌────────────────────────────┐
                        │        HTTP Layer          │
                        │  internal/httpapi          │
                        └────────────┬───────────────┘
                                     │
                        ┌────────────▼───────────────┐
                        │       Service Layer        │
                        │   internal/shortener       │
                        └────────────┬───────────────┘
                                     │
                  ┌──────────────────▼──────────────────┐
                  │     Storage + Metrics Layer         │
                  │ internal/store + internal/metrics   │
                  └─────────────────────────────────────┘


* **Handlers** → Parse HTTP requests and responses.
* **Service** → Contains business logic (shortening, resolving, metrics).
* **Store** → In-memory key-value store for URL ↔ code mappings.
* **Metrics** → Tracks top shortened domains.

# Features

* REST API to shorten any long URL
* Deterministic — same long URL returns same short code
* Redirection endpoint (/{code} → original URL)
* Metrics API: top 3 most shortened domains
* Thread-safe in-memory storage
* Unit-tested (store, metrics, service, handlers)

# Tech Stack

| Component        | Technology                           |
| ---------------- | ------------------------------------ |
| Language         | Go 1.25+                             |
| Router           | `net/http` (standard library)        |
| Storage          | In-memory maps with sync.RWMutex     |
| Testing          | Go’s built-in `testing` + `httptest` |
| Containerization | Docker (multi-stage build)           |

# Project Structure

url-shortener/
├── cmd/
│   └── server/
│       └── main.go             # Entry point
├── internal/
│   ├── httpapi/                # REST handlers
│   ├── metrics/                # Metrics tracker
│   ├── shortener/              # Business logic
│   └── store/                  # In-memory storage
├── pkg/
│   └── util/                   # Helper utilities
├── go.mod
├── Dockerfile
└── README.md

# Local Setup (without Docker)
**Prerequisites**
* Install Go 1.25+
* Clone the repository
```bash
git clone https://github.com/shomil999/url-shortener.git
cd url-shortener
```
* Verify Setup
```bash
go version
go mod tidy
```
* Run Locally
```bash
go run ./cmd/server
```
✅ Dockerized for easy local deployment
