##########
# Building
##########

build-docker-prod:
	docker build -f docker/Dockerfile -t mattgleich/api:latest .
build-docker-dev:
	docker build -f docker/dev.Dockerfile -t mattgleich/api:dev .
build-docker-dev-test:
	docker build -f docker/dev.test.Dockerfile -t mattgleich/api:test .
build-docker-dev-lint:
	docker build -f docker/dev.lint.Dockerfile -t mattgleich/api:lint .
build-go:
	go get -v -t -d ./...
	go build -v ./cmd/api
	rm api

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-hadolint:
	hadolint docker/Dockerfile
	hadolint docker/dev.Dockerfile
	hadolint docker/dev.test.Dockerfile
	hadolint docker/dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/api:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...
test-in-docker: build-docker-dev-test
	docker run mattgleich/api:test

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev-test build-docker-dev-lint build-docker-dev
