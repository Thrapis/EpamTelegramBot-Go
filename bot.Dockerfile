FROM golang:1.17-alpine

WORKDIR /bot

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY configs ./configs
COPY internal/api ./internal/api
COPY internal/config ./internal/config
COPY internal/database ./internal/database

COPY cmd/bot ./cmd/bot
COPY internal/bot ./internal/bot

RUN go build -o main ./cmd/bot/main.go

CMD [ "./main" ]