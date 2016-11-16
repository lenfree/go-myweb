FROM golang:1.7.1
MAINTAINER lenfree.yeung@gmail.com

ENV GOPATH=/go

WORKDIR /go/src/github.com/lenfree/go-myweb

RUN mkdir -p /go/src/github.com/lenfree/go-myweb
COPY . /go/src/github.com/lenfree/go-myweb

CMD ["go run *.go"]
