#@IgnoreInspection BashAddShebang

export APP=lookups

export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

all: install format lint

install:
	go mod download

check-formatter:
	which gofumpt || GO111MODULE=off go get -u mvdan.cc/gofumpt
	which goimports || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

format: check-formatter
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R goimports -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R gofmt -s -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R gofumpt -s -w R

check-linter:
	which golangci-lint || GO111MODULE=off curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.1

lint: check-linter
	golangci-lint run --deadline 10m $(ROOT)/...

test:
	go test -v -race -p 1 ./...

ci-test:
	go test -v -race -p 1 -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func coverage.txt
