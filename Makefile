.PHONY: default
default: test

include common.mk

.PHONY: test
test: go-test-all

.PHONY: lint
lint: go-lint-all git-clean-check

.PHONY: generate
generate: buf-generate-all

.PHONY: build-server
build-server:
	go build -o ./bin/server ./server/cmd/

.PHONY: build-docker-server
build-docker-server:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/rbac-server:latest -f build/server/Dockerfile .

.PHONY: build-docker-envsubst
build-docker-envsubst:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/envsubst:latest -f build/envsubst/Dockerfile .
