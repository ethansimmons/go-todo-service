FROM golang:1.23 AS builder
# Install git and ca-certificates (needed to be able to call HTTPS)
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    git
WORKDIR /app
COPY . ./
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o ./item-service ./cmd/server/main.go

FROM alpine:latest AS item-service
WORKDIR /app
# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Copy app binary
COPY --from=builder /app/item-service .
EXPOSE 443
ENTRYPOINT ["./item-service"]

