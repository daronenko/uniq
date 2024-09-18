CMD = cmd/main.go
NAME = uniq

.PHONY: run
run:
	@go run $(CMD)

.PHONY: build
build:
	@go build -o bin/$(NAME) $(CMD)

.PHONY: format
format:
	@gofmt -s -w .

.PHONY: test
test:
	@go clean -testcache
	@go test ./... -coverprofile=coverage.out

.PHONY: --check-coverage
--check-coverage:
	@if [ ! -f coverage.out ]; then \
		echo "coverage does not exist, running tests..."; \
		$(MAKE) test; \
	fi

.PHONY: view-coverage
view-coverage: --check-coverage
	@go tool cover -func=coverage.out

.PHONY: view-coverage-html
view-coverage-html: --check-coverage
	@go tool cover -html=coverage.out -o coverage.html
	@open coverage.html
