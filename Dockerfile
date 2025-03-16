# Use the official Golang image as the base image
FROM golang:1.24.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 9998
EXPOSE 9998

# Command to run the application
CMD ["./main"]