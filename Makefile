DOCKER_IMAGE_NAME=smoya/graphql-go-workshop
DOCKER_IMAGE_TAG=latest

GOTOOLS = github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: tools
tools:
	go get $(GOTOOLS)

.PHONY: build
build:
	GO111MODULE=on go build -o bin/graphql-go-workshop cmd/server/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .