# Make this makefile self-documented with target `help`
.PHONY: help
.DEFAULT_GOAL := help
help: ## Show help
	@grep -Eh '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint: download ## Lint the repository with golang-ci lint
	golangci-lint run --max-same-issues 0 --max-issues-per-linter 0 $(if $(CI),--out-format code-climate > gl-code-quality-report.json 2>golangci-stderr-output)

.PHONY: test
test: download ## Run all tests
	go test -v

.PHONY: bench
bench: download ## Run all benchmarks
	go test -bench=.

.PHONY: download
download: ## Download dependencies
	go mod download
