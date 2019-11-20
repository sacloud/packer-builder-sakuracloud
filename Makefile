TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
CURRENT_VERSION = $(gobump show -r version/)
export GO111MODULE=on

default: test vet

.PHONY: tools
tools:
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get -u github.com/motemen/gobump/cmd/gobump
	GO111MODULE=off go get -u golang.org/x/lint/golint

.PHONY: clean
clean:
	rm -Rf $(CURDIR)/bin/*

.PHONY: install build build-x
install: build
	cp -f $(CURDIR)/bin/packer-builder-sakuracloud $(GOPATH)/bin/packer-builder-sakuracloud

build: clean 
	go build -mod vendor -ldflags "-s -w" -o $(CURDIR)/bin/packer-builder-sakuracloud $(CURDIR)/main.go

build-x: clean vet
	sh -c "'$(CURDIR)/scripts/build.sh'"


.PHONY: test testacc
test: vet
	go test $(TEST) $(TESTARGS) -v -timeout=30m -parallel=4 ;

# testacc runs acceptance tests
testacc:
	@echo "WARN: Acceptance tests will take a long time to run and may cost money. Ctrl-C if you want to cancel."
	PACKER_ACC=1 go test -v $(TEST) $(TESTARGS) -timeout=45m

.PHONY: lint vet fmt golint goimports
lint: vet fmt golint goimports

vet: fmt
	go vet ./...

golint: fmt
	test -z "$$(go list ./... | xargs -L1 golint | fgrep -v 'should be BuilderID'  | fgrep -v 'should be ID' | tee /dev/stderr )"

fmt:
	gofmt -s -l -w $(GOFMT_FILES)

goimports:
	goimports -w $(GOFMT_FILES)

.PHONY: docker-shell docker-test docker-testacc docker-build
docker-shell:
	docker-compose run --rm packer

docker-test:
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'test'"

docker-testacc:
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'testacc'"

docker-build: clean 
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'build-x'"

.PHONY: prepare-homebrew
prepare-homebrew:
	rm -rf homebrew-packer-builder-sakuracloud/; \
	sh -c "'$(CURDIR)/scripts/update_homebrew_formula.sh' '$(CURRENT_VERSION)'"

.PHONY: version bump-patch bump-minor bump-major
version:
	gobump show -r version/

bump-patch:
	gobump patch -w version/

bump-minor:
	gobump minor -w version/

bump-major:
	gobump major -w version/

.PHONY: default
