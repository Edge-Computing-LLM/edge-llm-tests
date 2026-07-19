# Repository instructions

This repository is the Go-first validation and sanitized evidence plane for the
Edge-Computing-LLM organization. It may run repository gates, read Kubernetes
state, execute fixed public smoke prompts, and create short-lived validation
artifacts. It must not install or uninstall k3s, charts, models, packages, or
host configuration.

Never persist Secrets, kubeconfig content, access tokens, host IPs, environment
dumps, pod logs, private prompts, model responses, or model weights. Command
output must pass through the sanitizer and remain bounded. Credential and
large-file directories outside Git repositories are always excluded.

Before completing a change run `gofmt`, `go test ./...`, `go vet ./...`, and
`go build -o /tmp/edge-llm-tests ./cmd/edge-llm-tests`. Keep cluster checks
read-only and make skipped checks explicit.
