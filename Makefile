NAME=todo
BINDIR=build
VERSION=1.0.0
BUILDTIME=$(shell date -u)
GOBUILD=go build
GOFLAGS=-ldflags '-s -w -X "main.version=$(VERSION)" -X "main.buildTime=$(BUILDTIME)"'

all: linux-amd64

linux-amd64: 
	GOOS=linux GOARCH=amd64 cd client && $(GOBUILD) -o ../$(BINDIR)/$(NAME)-client-$@
	GOOS=linux GOARCH=amd64 cd server && $(GOBUILD) -o ../$(BINDIR)/$(NAME)-server-$@
