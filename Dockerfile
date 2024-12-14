# Use a newer Go version (1.23 or higher)
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Run 'go mod tidy' to download the necessary dependencies
RUN go mod tidy

# Copy the rest of the application code into the container
COPY . .

# Build the Go application
RUN go build -o myapp .

# Expose the port the app will run on (adjust as needed)
EXPOSE 8080

# Run the application
CMD ["./myapp"]
