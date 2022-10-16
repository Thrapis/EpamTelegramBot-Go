FROM golang:1.17-alpine

WORKDIR /bot

COPY bot/go.mod bot/go.sum ./

RUN go mod download && go mod verify

COPY bot/*.go ./
COPY bot/config.yaml ./

RUN go build -o main .

CMD [ "./main" ]