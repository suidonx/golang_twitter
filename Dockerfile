# development
FROM golang:alpine AS development

RUN go install github.com/air-verse/air@latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

CMD [ "air", "-c", ".air.toml" ]

# Production
FROM alpine AS production

WORKDIR /app

COPY --from=development /app/main .

EXPOSE 8080

CMD ["/app/main"]