package suite

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestClusterSuiteMaintainsEvidenceBoundary(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("cluster.go"))
	if err != nil {
		t.Fatal(err)
	}
	source := strings.ToLower(string(data))
	for _, forbidden := range []string{"kubectl logs", "get secrets", "config view --raw", "printenv"} {
		if strings.Contains(source, forbidden) {
			t.Fatalf("cluster suite contains forbidden evidence command %q", forbidden)
		}
	}
}
