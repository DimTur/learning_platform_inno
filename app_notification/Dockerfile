# Build stage
FROM golang:1.23.2-alpine3.20 AS builder

# Installing build dependencies
RUN apk --no-cache add build-base

# Setting up the working directory and copying dependencies
WORKDIR /notification
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .

# Building the application
RUN go build -o notification ./cmd/main.go

# Final stage
FROM alpine:3.20

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /notification

# Copying the binary
COPY --from=builder /notification/notification .

# Setting permissions
RUN chmod +x ./notification \
    && chown -R appuser:appgroup /notification

# Specifying a port for a container
EXPOSE 8003

# Switch to non-root user
USER appuser

# Launch command
ENTRYPOINT ["./notification", "serve", "--config=/notification/config.yaml"]