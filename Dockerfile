FROM golang:alpine AS development

RUN go install github.com/air-verse/air@latest

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/server ./cmd/server

CMD ["air" "-c", ".air.toml"]
