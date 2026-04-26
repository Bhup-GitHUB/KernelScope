GO ?= go
BIN_DIR := bin
GOCACHE ?= $(CURDIR)/.gocache
GOMODCACHE ?= $(CURDIR)/.gomodcache
GOLANGCI_LINT_CACHE ?= $(CURDIR)/.golangci-lint-cache

.PHONY: help build test lint dev clean

help:
	@printf "Available targets:\n"
	@printf "  build  Build agent and collector binaries\n"
	@printf "  test   Run Go tests\n"
	@printf "  lint   Run golangci-lint\n"
	@printf "  dev    Start local development infrastructure\n"
	@printf "  clean  Remove build artifacts\n"

build:
	@mkdir -p $(BIN_DIR)
	env GOCACHE="$(GOCACHE)" GOMODCACHE="$(GOMODCACHE)" $(GO) build -o $(BIN_DIR)/agent ./cmd/agent
	env GOCACHE="$(GOCACHE)" GOMODCACHE="$(GOMODCACHE)" $(GO) build -o $(BIN_DIR)/collector ./cmd/collector

test:
	env GOCACHE="$(GOCACHE)" GOMODCACHE="$(GOMODCACHE)" $(GO) test ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 || { printf "golangci-lint is required. Install it from https://golangci-lint.run/welcome/install/\n" >&2; exit 1; }
	env GOCACHE="$(GOCACHE)" GOMODCACHE="$(GOMODCACHE)" GOLANGCI_LINT_CACHE="$(GOLANGCI_LINT_CACHE)" golangci-lint run

dev:
	@if docker compose version >/dev/null 2>&1; then \
		docker compose -f deploy/dev/compose.yaml up -d; \
	elif command -v docker-compose >/dev/null 2>&1; then \
		docker-compose -f deploy/dev/compose.yaml up -d; \
	else \
		printf "Docker Compose is required. Install docker compose or docker-compose.\n" >&2; \
		exit 1; \
	fi

clean:
	rm -rf $(BIN_DIR)
