# Start from the Go official base image.
FROM golang:1.20 as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy Go mod and sum files.
# go.sum isn't here yet because there are no deps
# COPY go.mod go.sum ./
COPY go.mod ./

# Download all Go dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Start a new stage from a minimal Alpine image.
FROM alpine:latest

# Set the working directory.
WORKDIR /root/

# Copy the binary from the `builder` stage.
COPY --from=builder /app/main .

# Expose the port the app runs on.
EXPOSE 8080

# Command to run the application.
CMD ["./main"]