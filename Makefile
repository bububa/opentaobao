GIT_TAG = $(shell git tag | grep ^v | sort -V | tail -n 1)
GIT_COMMIT = $(shell git rev-parse --short HEAD)
GIT_SUMMARY = $(shell git describe --tags --dirty --always)
GIT_DATE = $(shell git show --format='%ct' HEAD --quiet)
LDFLAGS = -X main.version=$(GIT_TAG) -X main.commit=$(GIT_COMMIT) -X main.GitDate=${GIT_DATE}

tools: downloader generator

downloader:
	go build -o bin/downloader github.com/bububa/opentaobao/metadata/downloader

generator:
	go build -ldflags "$(LDFLAGS)" -o bin/generator github.com/bububa/opentaobao/metadata/generator
