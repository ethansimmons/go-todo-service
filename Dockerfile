FROM golang:1.23-alpine AS builder
ARG ITEM_SERVICE_ADDRESS
# Install git and ca-certificates (needed to be able to call HTTPS)
RUN apk --update add ca-certificates git
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./item-service ./cmd/server/main.go
RUN go build -ldflags "-X main.itemServiceAddr=$ITEM_SERVICE_ADDRESS" -o ./gateway-service ./cmd/client/main.go

FROM alpine:latest AS item-service
WORKDIR /app
# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Copy app binary
COPY --from=builder /app/item-service .
EXPOSE 50051
ENTRYPOINT ["./item-service"]

FROM alpine:latest AS gateway-service
WORKDIR /app
# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Copy app binary
COPY --from=builder /app/gateway-service .
EXPOSE 8080
ENTRYPOINT ["./gateway-service"]

