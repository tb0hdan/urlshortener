BUILD = $(shell git rev-parse HEAD)
BDATE = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p_UTC')
GO_VERSION = $(shell go version|awk '{print $$3}')
VERSION = $(shell cat ./VERSION)
TESTS = test-server
LINTS = lint-codec lint-miscellaneous lint-server lint-storage lint-storage/memory
all: lint slow-lint test build

test: $(TESTS)

slow-lint: $(LINTS)

$(LINTS):
	@golint -set_exit_status=1 $(shell echo $@|awk -F'lint-' '{print $$2}')

$(TESTS):
	@go test -bench=. -v -benchmem -race ./$(shell echo $@|awk -F'test-' '{print $$2}')

build:
	@go mod why
	@go build -a -trimpath -tags netgo -installsuffix netgo -v -x -ldflags "-s -w -X main.Build=$(BUILD) -X main.BuildDate=$(BDATE) -X main.GoVersion=$(GO_VERSION) -X main.Version=$(VERSION)" -o urlshortener *.go

lint:
	@golangci-lint run --enable-all

deps:
	@go get -u golang.org/x/lint/golint
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.21.0
