GOPATH := $(shell pwd)
GOROOT := /Users/Arthur/.g/go
GO := $(GOROOT)/bin/go
SRC := $(GOPATH)/src
GOBIN := $(GOPATH)/bin

.PHONY: all
all: build

.PHONY: setup build clean


build : clean setup
	echo "building";
	mkdir bin && \
	cd bin && \
	$(GO) build $(SRC)/main/clock1.go
	echo "success"

run : clean setup
	echo "running";
	mkdir bin && \
	cd bin && \
	$(GO) build $(SRC)/main.go
	bin/main
	echo "success"
clean :
	rm -rf bin
test : setup
	$(GO) test -ttimeout=30s  ./...

.PHONY: setup
setup:
	export GOPATH=$(GOPATH)