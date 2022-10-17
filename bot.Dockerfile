FROM golang:1.17-alpine

WORKDIR /bot

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY cmd/bot ./cmd/bot
COPY config ./config
COPY internal/bot ./internal/bot
COPY internal/database ./internal/database
COPY config.yaml ./

RUN go build -o main ./cmd/bot/main.go

CMD [ "./main" ]