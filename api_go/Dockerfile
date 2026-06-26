# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o sys-monitor-api-go .

# Runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sys-monitor-api-go .

EXPOSE 8000

CMD ["./sys-monitor-api-go"]