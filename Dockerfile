FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
COPY cmd ./cmd
COPY internal ./internal

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Expose the port your app runs on
EXPOSE 8080

# Run
CMD ["/app/main"]