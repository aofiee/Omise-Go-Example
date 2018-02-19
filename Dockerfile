FROM golang:1.10-alpine
MAINTAINER arnon kijlerdphon (snappy.kop@gmail.com)

RUN mkdir -p /go/src/Omise-Go-Example
#ADD . /go/src/Omise-Go-Example
WORKDIR /go/src/Omise-Go-Example

RUN apk add --no-cache git curl \
    && rm -rf /var/cache/apk/*

RUN set -x \
    && curl https://glide.sh/get | sh \
    # go get revel
    && go get -v github.com/revel/revel \
    && go get -v github.com/revel/cmd/revel

EXPOSE 9000