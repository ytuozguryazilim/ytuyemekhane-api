FROM golang:1.12.1-alpine as builder

WORKDIR /go/src/github.com/GnuYtuce/ytuyemekhane-api
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/ytuyemekhane-api .

FROM alpine:latest
COPY --from=builder /go/bin/ytuyemekhane-api .

LABEL maintainer="Fatih Sarhan <f9n@protonmail.com>"

ENV PORT 8080

EXPOSE 8080
CMD ["./ytuyemekhane-api"]
