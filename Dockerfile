FROM golang:1.8-alpine

RUN mkdir -p /go/src/luckyzune.com/yqbit/parkomat

COPY . /go/src/luckyzune.com/yqbit/parkomat

RUN apk add --update git

RUN go get luckyzune.com/yqbit/parkomat/...

EXPOSE 53
EXPOSE 53/udp

CMD parkomat
