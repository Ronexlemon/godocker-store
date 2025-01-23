# Start from the Go base image
FROM golang:alpine AS builder

# Enable Go modules
ENV GO111MODULE=on

# Install git and other dependencies
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /client

# Copy go.mod and go.sum files
COPY go.mod go.sum ./



# Download Go modules dependencies
#RUN go mod tidy

# Copy the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# Start a new stage from scratch
FROM scratch

# Set the working directory in the new stage
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /client/bin/main .

# Expose the application port
EXPOSE 8080

# Run the executable
CMD ["./main"]
