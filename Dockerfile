FROM golang:1.10
MAINTAINER Giorgio Crivellari <miticojo@gmail.com>

ARG PORT=2000
ARG PROTO=udp

WORKDIR /go/src/app

ADD bin/syslog-massive-client .
ADD bin/syslog-massive-server .

EXPOSE $PORT


ENTRYPOINT [ "/go/src/app/syslog-massive-server", "$PROTO", "$PORT" ]
