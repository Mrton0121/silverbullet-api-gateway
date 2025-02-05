# Start with the official Golang image
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o app .

# Use a minimal image for the final container
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Expose the necessary port
EXPOSE 8080

# Run the application
CMD ["./app"]
