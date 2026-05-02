FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o marketplace .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/marketplace .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./marketplace"]