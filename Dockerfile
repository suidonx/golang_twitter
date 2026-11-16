# development
FROM golang:alpine AS development

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD [ "air", "-c", ".air.toml" ]

# Production
FROM alpine AS production

WORKDIR /app

COPY --from=development /app/main .

EXPOSE 8080

CMD ["/app/main"]