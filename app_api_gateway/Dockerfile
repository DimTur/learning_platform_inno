# Build stage
FROM golang:1.23.2-alpine3.20 AS builder

# Installing build dependencies
RUN apk --no-cache add build-base

# Setting up the working directory and copying dependencies
WORKDIR /api-gateway
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .

# Building the application
RUN go build -o api-gateway ./cmd/main.go

# Final stage
FROM alpine:3.20

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /api-gateway

# Copying the binary
COPY --from=builder /api-gateway/api-gateway .

# Setting permissions
RUN chmod +x ./api-gateway \
    && chown -R appuser:appgroup /api-gateway

# Specifying a port for a container
EXPOSE 8000

# Switch to non-root user
USER appuser

# Launch command
ENTRYPOINT ["./api-gateway", "serve", "--config=/api-gateway/config.yaml"]