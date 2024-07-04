DEFAULT: help

help: ## Show commands of the makefile (and any included files)
	@awk 'BEGIN {FS = ":.*?## "}; /^[0-9a-zA-Z_.-]+:.*?## .*/ {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

match.build: ## Build matching tool
	CGO_ENABLED=0 go build -o build/match cmd/main.go

match.bench: ## Benchmark the tooling
	CGO_ENABLED=0 go test -bench=. ./bench -v -benchmem

match.linear: match.build ## Run linear matching on KEYWORD
	./build/match -l -k $(KEYWORD)

match.regex: match.build ## Run linear matching on KEYWORD
	./build/match -r -k $(KEYWORD)

match.rolling-hash: match.build ## Run linear matching on KEYWORD
	./build/match -rh -k $(KEYWORD)