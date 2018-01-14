FROM golang

ADD . /go/src/github.com/GnuYtuce/ytuyemekhane-api
RUN go install github.com/GnuYtuce/ytuyemekhane-api

ENV PORT 8080

ENTRYPOINT [ "/go/bin/ytuyemekhane-api" ]

EXPOSE 8080