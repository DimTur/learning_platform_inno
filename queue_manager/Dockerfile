# Build stage
FROM golang:1.23.2-alpine3.20 AS builder

# Installing build dependencies
RUN apk --no-cache add build-base

# Setting up the working directory and copying dependencies
WORKDIR /queue-manager
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .

# Building the application
RUN go build -o queue-manager ./cmd/main.go

# Final stage
FROM alpine:3.20

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /queue-manager

# Copying the binary
COPY --from=builder /queue-manager/queue-manager .

# Setting permissions
RUN chmod +x ./queue-manager \
    && chown -R appuser:appgroup /queue-manager

# Switch to non-root user
USER appuser

# Launch command
ENTRYPOINT ["./queue-manager", "serve", "--config=config.yaml"]