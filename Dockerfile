# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

RUN go install github.com/air-verse/air@latest 

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Expose the port that the application listens on
EXPOSE 8080

# Set the command to run the executable
CMD ["./main"]