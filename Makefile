.PHONY: all
all: generate

.PHONY: generate
generate:
	@abigen --sol ./contracts/CatchallResolver.sol --pkg bindings --out ./bindings/CatchallResolver.go

.PHONY: test
test: generate
	@go test ./test $(TEST_ARGS)
