FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o uuid-generator-app .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/uuid-generator-app .

EXPOSE 8080

CMD ["./uuid-generator-app"]