# Use the official Golang image for building
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Start a minimal runtime image
FROM alpine:latest

# Install a shell (sh or bash) and networking tools (for health checks)
RUN apk add --no-cache bash curl


# Set environment variable for port
ENV PORT=8080
EXPOSE 8080

# Add a non-root user for security
RUN adduser -u 1000 -D appuser
USER appuser

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /

# Health check using HTTP (recommended)
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:8080/health || exit 1

# Run the Go binary
CMD ["/main"]


