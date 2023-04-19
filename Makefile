.PHONY: all
all: hooks update format lint

.PHONY: format
format:
	go fmt ./...

# Go mod tidy
.PHONY: tidy
tidy:
	go mod tidy

# Update all dependencies
.PHONY: update
update: tidy

# Configure git hooks
.PHONY: hooks
hooks:
	git config core.hooksPath hooks

# Lint golang code
.PHONY: lint
lint:
	golangci-lint run -v --fix -c .golangci.yaml ./...

.PHONY: mock
mock: mock_clean
	go generate ./...

.PHONY: mock_clean
mock_clean:
	find . -name "*.go" -path "**/mocks/*" | while read file; do rm $$file; done;

.PHONY: run
run:
	go run cmd/main.go
