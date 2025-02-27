# Use the official Golang image as the base image
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install golang-migrate
RUN go install -tags 'netgo' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o api ./cmd/api

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./api"]