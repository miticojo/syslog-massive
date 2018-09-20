GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: server client

server:
	GOOS=linux GOARCH=${GOARCH} GOBIN=bin go build ${LDFLAGS} syslog-massive-server.go

client:
	GOOS=linux GOARCH=${GOARCH} GOBIN=bin go build ${LDFLAGS} syslog-massive-client.go

.PHONY: server client
