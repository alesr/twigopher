.DEFAULT_GOAL := help

PROJECT_NAME := twigopher

.PHONY: help

help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run tests inside Docker container
	@docker build -t "${PROJECT_NAME}"-test -f resources/test/Dockerfile .
	@docker run --rm "${PROJECT_NAME}"-test
