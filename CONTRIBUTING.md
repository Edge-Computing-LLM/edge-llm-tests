# Contributing

Use focused Go changes and preserve the read-only live-validation boundary.

Before opening a pull request:

```bash
gofmt -w ./cmd ./internal
go test ./...
go test -race ./...
go vet ./...
go build -o /tmp/edge-llm-tests ./cmd/edge-llm-tests
```

New checks need a stable ID, scope, plain-language description, bounded timeout,
and a clear required/optional decision. Never add a sensitive command with the
expectation that sanitization will make it safe later.
