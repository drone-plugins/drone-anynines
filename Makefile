.PHONY: clean deps fmt vet test docker

EXECUTABLE ?= drone-anynines

LDFLAGS = -X "main.buildDate=$(shell date -u '+%Y-%m-%d %H:%M:%S %Z')"
PACKAGES = $(shell go list ./... | grep -v /vendor/)
ARCH = $(shell uname -m)

clean:
	go clean -i ./...

deps:
	go get -t ./...

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

docker:
ifeq ($(ARCH),$(filter $(ARCH), armv6l armv7l))
	sed -i 's/FROM alpine/FROM armhfbuild\/alpine/' Dockerfile
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)'
	docker build --rm -t armhfplugins/$(EXECUTABLE) .
else
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)'
	docker build --rm -t plugins/$(EXECUTABLE) .
endif

$(EXECUTABLE): $(wildcard *.go)
	go build -ldflags '-s -w $(LDFLAGS)'

build: $(EXECUTABLE)
