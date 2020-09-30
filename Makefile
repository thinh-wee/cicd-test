# SHELL = /usr/bin/env bash

GOCMD=go

GO_LDFLAGS	:= -s -w
GO_FLAGS	:= -ldflags "-extldflags \"-static\" $(GO_LDFLAGS)"

build:
	$(GOCMD) build $(GO_FLAGS) -o bin/application -trimpath api-gateway/*.go

run:
	/bin/bash -c bin/application