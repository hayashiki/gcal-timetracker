VERSION := $(shell cat VERSION)

SRCS    := $(shell find . -type f -name '*.go')
#LDFLAGS := "-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
ARCHIVE := "TimeTracker-$(VERSION).alfredworkflow"

export GO111MODULE ?= on

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o "./.workflow/tt"; \
	cd ./.workflow; \
	envsubst >./info.plist <./info.plist.template; \
	zip -r "../bin/${ARCHIVE}" ./*; \
	zip -d "../bin/${ARCHIVE}" info.plist.template;
#	go build -ldflags=$(LDFLAGS) -o bin/$(NAME) ./cmd/gipr

.PHONY: tag
tag:
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push --tags
