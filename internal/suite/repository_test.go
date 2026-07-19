package suite

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
)

func TestGoFormattingCheck(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "valid.go"), []byte("package example\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if check := goFormattingCheck(dir, "fixture"); check.Status != model.Pass {
		t.Fatalf("formatted source rejected: %#v", check)
	}
	if err := os.WriteFile(filepath.Join(dir, "invalid.go"), []byte("package example\nfunc x( ){ }\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if check := goFormattingCheck(dir, "fixture"); check.Status != model.Fail {
		t.Fatalf("unformatted source accepted: %#v", check)
	}
}
