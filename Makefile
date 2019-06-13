NAME=logrus-prefixed-formatter
PACKAGES=$(shell go list ./...)

deps:
	@echo "--> Installing dependencies"
	@go mod vendor

test:
	@echo "--> Running tests"
	@go test -timeout 30s ./ -v -count 1

format:
	@echo "--> Running go fmt"
	@go fmt $(PACKAGES)
