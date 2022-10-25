FROM golang:1.17-alpine

WORKDIR /web-app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY configs ./configs
COPY internal/api ./internal/api
COPY internal/config ./internal/config
COPY internal/database ./internal/database

COPY cmd/web ./cmd/web
COPY internal/web ./internal/web

RUN go build -o main ./cmd/web/main.go

EXPOSE 8080

CMD [ "./main" ]