package sanitize

import (
	"os"
	"strings"
	"testing"
)

func TestTextRedactsSensitiveEvidence(t *testing.T) {
	input := "server 192.0.2.10 token=abc123 password: hunter2 /home/operator/file /media/operator/disk Bearer abc.def"
	output := Text(input)
	for _, forbidden := range []string{"192.0.2.10", "abc123", "hunter2", "/home/operator", "/media/operator", "abc.def"} {
		if strings.Contains(output, forbidden) {
			t.Fatalf("sanitized output contains %q: %s", forbidden, output)
		}
	}
}

func TestTextRedactsLocalIdentity(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		t.Skip("hostname unavailable")
	}
	if got := Text("node=" + hostname); strings.Contains(got, hostname) {
		t.Fatalf("hostname was not redacted: %s", got)
	}
}
