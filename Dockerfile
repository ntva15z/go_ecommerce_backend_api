# --- Build stage ---
FROM golang:alpine AS builder

WORKDIR /build

# Cài thêm git (bắt buộc cho go mod)
RUN apk add --no-cache git

# Copy go mod trước để cache dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o crm.shopdev.com ./cmd/server

# --- Final stage ---
FROM alpine:3.20

WORKDIR /app

# Cài chứng chỉ SSL để gọi HTTPS
RUN apk add --no-cache ca-certificates

# Copy config và binary
COPY ./config ./config
COPY --from=builder /build/crm.shopdev.com .

# Run binary
ENTRYPOINT ["./crm.shopdev.com", "config/local.yaml"]
