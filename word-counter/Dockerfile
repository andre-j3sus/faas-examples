FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o word-counter-app .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/word-counter-app .

EXPOSE 8080

CMD ["./word-counter-app"]