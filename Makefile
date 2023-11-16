#!/usr/bin/make -f

DOCKER := $(shell which docker)

#################
###   Build   ###
#################
test:
	@echo "--> Running tests"
	go test -v ./...

test-integration:
	@echo "--> Running integration tests"
	cd integration; go test -v ./...

.PHONY: test test-integration

##################
###  Protobuf  ###
##################

protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all:
	go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm@latest
	go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm-proto@latest
	(cd proto; buf generate --template buf.gen.proto.yaml)
	(cd proto; buf generate)
	rm -rf ./api
	mv ./sonrhq/service ./api

.PHONY: proto-all proto-gen proto-format proto-lint

#################
###  Linting  ###
#################

golangci_lint_cmd=golangci-lint
golangci_version=v1.51.2

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --timeout 15m

lint-fix:
	@echo "--> Running linter and fixing issues"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --fix --timeout 15m

.PHONY: lint lint-fix
