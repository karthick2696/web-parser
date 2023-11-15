# Use a lightweight Golang image as the build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy only the necessary files for dependency download
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main

# Use a lightweight Alpine image as the final stage
FROM alpine:latest

WORKDIR /app

# Copy only the built executable from the builder stage
COPY --from=builder /app/main .

# Set the entry point for the container
ENTRYPOINT ["./main"]
