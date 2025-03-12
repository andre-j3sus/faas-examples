FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o palindrome-checker-app .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/palindrome-checker-app .

EXPOSE 8080

CMD ["./palindrome-checker-app"]