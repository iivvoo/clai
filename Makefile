default: clai

BUILD_TIME=$(shell date +%FT%T%z)
GIT_REVISION=$(shell git rev-parse --short HEAD)
GIT_BRANCH=$(shell git rev-parse --symbolic-full-name --abbrev-ref HEAD)
GIT_DIRTY=$(shell git diff-index --quiet HEAD -- || echo "x-")

# Optionally include RELEASE_TAG, set by github when building
LDFLAGS=-ldflags "-s -X main.BuildStamp=$(BUILD_TIME) -X main.GitHash=$(GIT_DIRTY)$(GIT_REVISION) -X main.gitBranch=$(GIT_BRANCH)"


srcfiles = main.go */*.go

testpackages = ./...

default: bin/clai

bin/clai: $(srcfiles)
	go build -o bin/clai $(LDFLAGS) main.go

build-release: $(srcfiles)
	go build -o clai-release $(LDFLAGS) main.go

lint:
	golangci-lint run

test-dev: lint test

test: 
	go test -count=1 $(testpackages)

test-verbose:
	go test -v $(testpackages)

mod-tidy:
	go mod tidy

show-deps:
	go list -m all

clean:
	@rm -f bin/*

