# ----------------------------------------
# Stage 1: Build stage
# ----------------------------------------
FROM golang:1.25-alpine AS builder


# Install Git (needed for Go modules)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker caching)
COPY go.mod ./

# Download dependencies (will be cached unless go.mod/go.sum change)
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-shortener ./cmd/server

# ----------------------------------------
# Stage 2: Runtime stage (small image)
# ----------------------------------------
FROM alpine:3.19

# Create a non-root user for security
RUN adduser -D appuser

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/url-shortener .

# Expose port (for documentation)
EXPOSE 8080

# Environment variables (can be overridden)
ENV PORT=8080
ENV BASE_URL=http://localhost:8080

# Run as non-root user
USER appuser

# Command to run the app
ENTRYPOINT ["./url-shortener"]
