HAS_YQ=$(shell which yq)
HAS_MIGRATE=$(shell which migrate)
NS=session

.PHONY: help
help: ## help for telmei-go
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# go work
.PHONY: go_work
go_work: ## go work
ifeq ("$(wildcard go.work)","")
	go work init internals
else
	go work edit -use internals
endif
	ls services | xargs -Iarg go work edit -use services/arg

# go mod tidy
.PHONY: go_tidy
go_tidy: ## go mod tidy
	@make go_work
	@bash scripts/mod_tidy.sh
	go work sync
