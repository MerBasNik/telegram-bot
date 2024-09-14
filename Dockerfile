FROM golang:1.22-alpine3.20 AS builder

COPY . /github.com/MerBasNik/telegram-bot/
WORKDIR /github.com/MerBasNik/telegram-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/MerBasNik/telegram-bot/bin/bot .
COPY --from=0 /github.com/MerBasNik/telegram-bot/configs configs/

EXPOSE 80

CMD [ "./bot" ]