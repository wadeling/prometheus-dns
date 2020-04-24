SHELL=/bin/bash

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)


BINPATH=bin/$(GOOS)_$(GOARCH)
BINARY=prometheus-proxy

.PHONY: build
build:
	@echo "--> Building binary ..."
	@mkdir -p $(BINPATH) 
	@go build \
     -o="$(BINPATH)/$(BINARY)" \
     ./cmd/proxy
     
clean:
	rm -rf $(BINPATH)/$(BINARY)   

