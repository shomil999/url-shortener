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


Handlers → Parse HTTP requests and responses.
Service → Contains business logic (shortening, resolving, metrics).
Store → In-memory key-value store for URL ↔ code mappings.
Metrics → Tracks top shortened domains.
