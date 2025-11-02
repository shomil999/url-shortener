# Url Shortener

* A simple yet production-grade **URL Shortener Service** built in **Go (Golang)** — similar to Bitly.  
It provides REST APIs to shorten URLs, redirect users, and view analytics about the most shortened domains.

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

```
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
```

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

# Run with Docker

```
docker pull shomil99/url-shortener:v1
```

* If above doesn't work, then build locally
```
docker build -t shomil/url-shortener:v1 .
docker run -p 8080:8080 -e BASE_URL=http://localhost:8080 shomil/url-shortener:v1
```
```
🚀 Server started at http://localhost:8080
```

# API Endpoints

| Method | Endpoint                      | Description                        |
| ------ | ----------------------------- | ---------------------------------- |
| `POST` | `/api/v1/shorten`             | Shorten a long URL                 |
| `GET`  | `/{code}`                     | Redirect to original URL           |
| `GET`  | `/api/v1/metrics/top-domains` | Fetch top 3 most shortened domains |

# Example Usage

* Use postman or use following curl commands.
  
* **Shorten URL**
```
curl -X POST http://localhost:8080/api/v1/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.udemy.com/course/kubernetes"}'
```
**Response**
```json
{
  "short_url": "http://localhost:8080/aBc123",
  "code": "aBc123"
}
```

**Redirect**
```
curl -i http://localhost:8080/aBc123
```
** Response **
```
HTTP/1.1 302 Found
Location: https://www.udemy.com/course/kubernetes
```

**Metrics**
```
curl http://localhost:8080/api/v1/metrics/top-domains
```
** Response **
```json
[
  {"domain":"udemy.com","count":3},
  {"domain":"youtube.com","count":2},
  {"domain":"wikipedia.org","count":1}
]
```
