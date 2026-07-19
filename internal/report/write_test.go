package report

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
)

func TestWriteCreatesMachineAndHumanReports(t *testing.T) {
	directory := t.TempDir()
	result := model.Report{SchemaVersion: model.SchemaVersion, RunID: "test", StartedAt: time.Now().UTC(), Mode: "repository", Checks: []model.Check{{ID: "one", Scope: "test", Description: "passes", Status: model.Pass}}}
	result.Finalize()
	if err := Write(directory, result); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"summary.json", "report.md"} {
		data, err := os.ReadFile(filepath.Join(directory, name))
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), "test") {
			t.Fatalf("%s does not contain run ID", name)
		}
	}
}
