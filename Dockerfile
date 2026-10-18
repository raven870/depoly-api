# Stage 1: Build stage
FROM golang:1.27.1-alpine AS builder

# Working directory set karein
WORKDIR /app

# Dependencies install karne ke liye go.mod aur go.sum copy karein
COPY go.mod go.sum ./
RUN go mod download

# Poora source code copy karein (cmd, internal, etc.)
COPY . .

# Go application ko compile karein
# Entry point: cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Stage 2: Final lightweight run stage
FROM alpine:latest

# Security ke liye CA Certificates install karein (HTTPS API calls ke liye)
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Builder stage se binary executable copy karein
COPY --from=builder /app/main .

# Standard port (apne app ke port ke hisaab se change kar sakte hain, e.g. 8080)
EXPOSE 8080

# Application run karein
CMD ["./main"]