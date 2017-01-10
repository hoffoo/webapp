FROM quay.io/uservoice/go:1.7

ADD . "$GOPATH/src/github.com/hoffoo/webapp"

RUN go build -o /usr/bin/webapp github.com/hoffoo/webapp

ENTRYPOINT /usr/bin/webapp
