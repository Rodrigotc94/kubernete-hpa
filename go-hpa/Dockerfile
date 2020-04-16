FROM golang:1.10-alpine3.7
RUN apk add bash

ADD ./main.go /go/src
WORKDIR /go/src
RUN go build main.go


ENTRYPOINT [ "/go/src/main" ]
EXPOSE 8000