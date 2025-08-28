FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila binario optimizado
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server

# Imagen final mínima
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 5000

CMD ["./server"]
