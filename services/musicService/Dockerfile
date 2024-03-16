# Use the official Golang image to create a build artifact.
FROM golang:1.18 as builder

WORKDIR /app

# Copy go mod and sum files
# COPY go.mod ./           # These paths are correct if the build context is project root
# COPY services/userService/ .    # Adjust to copy only the relevant service files

# RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myservice ./cmd/musicService/

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/myservice .

# Command to run the executable
CMD ["./myservice"]
