FROM golang:1.16.10-alpine3.13 AS build-env
USER root
RUN apk add --no-cache git gcc musl-dev
RUN apk add --update make
RUN go get github.com/google/wire/cmd/wire
WORKDIR /go/src/github.com/devtron-labs/authenticator
ADD . /go/src/github.com/devtron-labs/authenticator/
RUN GOOS=linux make

FROM alpine:3.15.0
RUN apk add --no-cache ca-certificates
RUN apk add git --no-cache
COPY --from=build-env  /go/src/github.com/devtron-labs/authenticator/authenticator .
CMD ["./authenticator"]