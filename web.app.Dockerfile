FROM golang:1.17-alpine

WORKDIR /app

COPY web/go.mod web/go.sum ./

RUN go mod download && go mod verify

COPY web/*.go ./
COPY web/config.yaml ./

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]