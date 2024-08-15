# Use the official Go image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application's source code
COPY . .

# Build the Go application
RUN go build -o /radar_service cmd/main.go

# Command to run the application
CMD ["/radar_service"]

# Expose the port on which your Go API will run
EXPOSE 8080

