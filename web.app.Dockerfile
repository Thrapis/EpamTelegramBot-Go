FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY cmd/web-app ./cmd/web-app
COPY config ./config
COPY internal/web-app ./internal/web-app
COPY internal/database ./internal/database
COPY config.yaml ./

RUN go build -o main ./cmd/web-app/main.go

EXPOSE 8080

CMD [ "./main" ]