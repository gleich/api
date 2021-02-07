##########
# Building
##########

build-docker-prod:
	docker build -f docker/prod.Dockerfile -t mattgleich/api:latest .
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
	hadolint docker/prod.Dockerfile
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


###################
# Local Development
###################

dev-start:
	docker-compose up -d postgres
	docker-compose up -d pgweb
	docker-compose up -d github_scraper
	docker-compose up api

dev-reset:
	docker-compose down
	docker system prune -a


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
# Local development
dev-reboot: dev-reset dev-start
