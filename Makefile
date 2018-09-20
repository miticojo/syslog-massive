GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

BINDIR=$(PWD)/bin
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: server client docker

server:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINDIR}/syslog-massive-server syslog-massive-server.go

client:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINDIR}/syslog-massive-client syslog-massive-client.go

docker:
	docker build -t miticojo/syslog-massive-server:latest .

.PHONY: server client docker
