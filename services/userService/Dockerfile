# Use the official Golang image to create a build artifact.
FROM golang:1.18 as builder

WORKDIR /app

# Copy go mod and sum files
# COPY go.mod ./           # These paths are correct if the build context is project root
# COPY services/userService/ .    # Adjust to copy only the relevant service files

# RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o myservice ./cmd/userService/

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/myservice .

CMD ["./myservice"]
