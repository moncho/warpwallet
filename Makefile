SOURCE_FILES?=$$(go list ./... | grep -v /vendor/)
TEST_PATTERN?=.
TEST_OPTIONS?=

setup: ## Install all the build and lint dependencies
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
	go install github.com/pierrre/gotestcover@latest
	go install golang.org/x/tools/cmd/cover@latest

test: ## Run all the tests
	gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=60s

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint: ## Run all the linters
	golangci-lint run --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=dupl \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=60s \
		./...

ci: test lint  ## Run all the tests and code checks

build: ## Build warpwallet 
	go build -o warpwallet cmd/warpwallet/main.go

# Magic as explained here: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
