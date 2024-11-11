FROM golang:1.23.3-alpine3.20
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/query
WORKDIR /go/src/query
ADD . /go/src/query
