FROM golang:1.10
MAINTAINER Giorgio Crivellari <miticojo@gmail.com>

ENV PORT=2000
ENV PROTO=udp

WORKDIR /go/src/app

ADD bin/syslog-massive-client .
ADD bin/syslog-massive-server .
ADD entrypoint.sh .

EXPOSE $PORT

CMD /go/src/app/entrypoint.sh
