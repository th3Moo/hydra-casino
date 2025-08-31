FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o server server.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8082
CMD ["./server"]
