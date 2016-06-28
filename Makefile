GO ?= go
TAG := $(shell git describe --tags `git rev-list --tags --max-count=1`)

build:
	cd cmd/exago && GOOS=linux CGO_ENABLED=0 $(GO) build -tags netgo -ldflags="-w -s" -v -i -o exago

# Compile the binary for all available OSes and ARCHes.
.PHONY: buildall
buildall:
	cd cmd/exago && gox -output "dist/exago_{{.OS}}_{{.Arch}}"

# Run several automated source checks to get rid of the most simple issues.
# This helps keeping code review focused on application logic.
.PHONY: check
check:
	@echo "gometalinter"
	@! gometalinter --disable gotype,aligncheck,interfacer,structcheck --deadline 10s ./... | \
	  grep -vE 'vendor'

# "go test -i" builds dependencies and installs them into GOPATH/pkg,
# but does not run the tests.
.PHONY: test
test:
	$(GO) test $(glide novendor)

# Create Docker image
.PHONY: image
image: | build
	@echo "Building docker image"
	docker build --rm --force-rm=true -t exago/svc .

default: build
