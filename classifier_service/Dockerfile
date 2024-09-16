# Use the official Golang image as a build environment
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY classifier_service/go.mod classifier_service/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY classifier_service/ .

# Build the Go app
RUN go build -o classifier_service main.go

# Use a minimal Docker image for production
FROM gcr.io/distroless/base-debian10

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/classifier_service /classifier_service

# Expose the gRPC port
EXPOSE 50051

# Run the Go application
ENTRYPOINT ["/classifier_service"]
