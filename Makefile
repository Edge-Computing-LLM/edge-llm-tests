.PHONY: check test race vet build repository cluster all

check: test race vet build

test:
	go test ./...

race:
	go test -race ./...

vet:
	go vet ./...

build:
	go build -o bin/edge-llm-tests ./cmd/edge-llm-tests

repository:
	go run ./cmd/edge-llm-tests -mode repository -root ..

cluster:
	go run ./cmd/edge-llm-tests -mode cluster -root ..

all:
	go run ./cmd/edge-llm-tests -mode all -root ..
