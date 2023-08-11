# Start from the Go official base image.
FROM golang:1.20 as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the go.work and go.work.sum files to the Docker context.
COPY ../go.work ../go.work.sum ./

# Copy microservice's Go mod and sum files.
COPY go.mod ./genexample/

# Download all Go dependencies. This will consider the go.work file.
RUN go mod download

# Copy the source code of the microservice into the container.
COPY . genexample/

# Build the application using the go.work file.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./genexample

# Start a new stage from a minimal Alpine image.
FROM alpine:3.18

# Set the working directory.
WORKDIR /root/

# Copy the binary from the `builder` stage.
COPY --from=builder /app/main .

# Expose the port the app runs on.
EXPOSE 8080

# Command to run the application.
CMD ["./main"]