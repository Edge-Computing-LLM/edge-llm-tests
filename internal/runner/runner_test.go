package runner

import (
	"context"
	"testing"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
)

func TestRunnerCapturesAndSanitizesOutput(t *testing.T) {
	result := (Runner{OutputLimit: 128}).Run(context.Background(), Spec{ID: "test", Scope: "unit", Name: "printf", Args: []string{"address=192.0.2.10 token=abc123"}, Required: true})
	if result.Status != model.Pass {
		t.Fatalf("status = %s, output=%s", result.Status, result.Output)
	}
	if result.Output == "" || result.Output == "address=192.0.2.10 token=abc123" {
		t.Fatalf("output was not sanitized: %q", result.Output)
	}
}
