FROM golang:1.9-alpine

RUN apk add --no-cache git mercurial

WORKDIR /go/src/github/efrenfuentes/go_hello
COPY . .

RUN go get
RUN go build

EXPOSE 8080

ENTRYPOINT ["./go_hello"]
