# ---------- Build Stage ----------
    FROM golang:1.24.3-alpine AS builder

    WORKDIR /app
    
    # Optional: install git if using private modules
    RUN apk add --no-cache git
    
    # Copy go.mod and go.sum separately to leverage caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the source code
    COPY . .
    
    # Build a statically-linked binary
    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bot .
    
    # ---------- Runtime Stage ----------
    FROM alpine:3.19
    
    WORKDIR /app
    
    # Add TLS certs for HTTPS
    RUN apk add --no-cache ca-certificates
    
    # Copy the built binary and env file
    COPY --from=builder /app/bot .
    COPY .env .
    
    CMD ["./bot"]