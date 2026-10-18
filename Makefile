#@IgnoreInspection BashAddShebang

export APP=lookups

export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

all: install format lint

install:
	go mod download

check-formatter:
	which gofumpt || go install mvdan.cc/gofumpt@latest
	which goimports || go install golang.org/x/tools/cmd/goimports@latest
	which godot || go install github.com/tetafro/godot/cmd/godot@latest

format: check-formatter
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R goimports -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R gofmt -s -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R gofumpt -s -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R godot -w R

check-linter:
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.15.0

lint: check-linter
	golangci-lint run --timeout 10m $(ROOT)/...

test:
	go test -v -race -p 1 ./...

ci-test:
	go test -v -race -p 1 -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func coverage.txt
