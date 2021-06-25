GIT_TAG = $(shell git tag | grep ^v | sort -V | tail -n 1)
GIT_REVISION = $(shell git rev-parse --short HEAD)
GIT_SUMMARY = $(shell git describe --tags --dirty --always)
LDFLAGS = -X main.GitTag=$(GIT_TAG) -X main.GitRevision=$(GIT_REVISION) -X main.GitSummary=$(GIT_SUMMARY)

tools: downloader generator

downloader:
	go build -o bin/downloader github.com/bububa/opentaobao/metadata/downloader

generator:
	go build -ldflags "$(LDFLAGS)" -o bin/generator github.com/bububa/opentaobao/metadata/generator
