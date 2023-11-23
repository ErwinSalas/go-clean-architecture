# Build Stage
FROM golang:1.20 AS builder

WORKDIR /app

# Copy the source code into the container
COPY . .

# Ensure modules are downloaded and tidy up
RUN go mod tidy

# Build the Go app and place the binary in the /app directory
RUN CGO_ENABLED=0 GOOS=linux go build -o go-bank-api ./cmd/api/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/go-bank-api .
COPY /ssl/server.crt /app/server.crt
COPY /ssl/server.key /app/server.key
# Expose port 50052
EXPOSE 50052

# Run the executable
CMD ["./go-bank-api"]
