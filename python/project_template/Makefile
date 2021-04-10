# these will speed up builds, for docker-compose >= 1.25
export COMPOSE_DOCKER_CLI_BUILD=1
export DOCKER_BUILDKIT=1

## Execute all
.PHONY: all
all: down build up test

## Build
.PHONY: build
build:
	docker-compose build

## Up
.PHONY: up
up:
	docker-compose up -d app

## Down
.PHONY: down
down:
	docker-compose down --remove-orphans

## Run all the tests
.PHONY: test
test: up
	docker-compose run --rm --no-deps --entrypoint=pytest app /tests/unit /tests/integration /tests/e2e

## Run only unit-tests
.PHONY: unit-tests
unit-tests:
	docker-compose run --rm --no-deps --entrypoint=pytest app /tests/unit

## Run only integration-tests
.PHONY: integration-tests
integration-tests: up
	docker-compose run --rm --no-deps --entrypoint=pytest app /tests/integration

## Run only e2e-tests
.PHONY: e2e-tests
e2e-tests: up
	docker-compose run --rm --no-deps --entrypoint=pytest app /tests/e2e

## Show last logs of the containers
.PHONY: logs
logs:
	docker-compose logs app | tail -100

## Format code by conventions
.PHONY: formatter
formatter:
	black --version
	black $$(find * -name '*.py')
	isort $$(find * -name '*.py')

## typehint using mypy
.PHONY: typehint
typehint:
	mypy -V
	mypy --ignore-missing-imports src/

## Lint Python code using pylint
.PHONY: lint
lint:
	pylint --version
	pylint --rcfile=.pylintrc src/ 

## Execute formatter, type-hint, lint
.PHONY: me-pretty
me-pretty: formatter typehint lint 

## Execute all code quality tools and run the tests
.PHONY: me-happy
me-happy: me-pretty test