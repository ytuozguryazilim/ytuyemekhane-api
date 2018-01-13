FROM golang

ADD . /go/src/ytuyemekhane-api
RUN go install ytuyemekhane-api

ENTRYPOINT [ "/go/bin/ytuyemekhane-api" ]

EXPOSE 8080