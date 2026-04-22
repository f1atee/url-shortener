
FROM golang:1.22-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/shortener ./cmd/shortener


FROM alpine:3.19

COPY --from=builder /app/shortener /shortener

EXPOSE 8080

CMD ["/shortener"]
