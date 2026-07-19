package sanitize

import (
	"os"
	"os/user"
	"regexp"
	"strings"
)

var (
	ipv4        = regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
	homePath    = regexp.MustCompile(`/home/[^/\s]+`)
	mediaPath   = regexp.MustCompile(`/media/[^/\s]+`)
	sensitive   = regexp.MustCompile(`(?i)(token|password|passwd|secret|authorization|api[_-]?key)(\s*[:=]\s*|\s+)[^\s,;]+`)
	bearer      = regexp.MustCompile(`(?i)bearer\s+[A-Za-z0-9._~+/=-]+`)
	githubToken = regexp.MustCompile(`\b(?:gh[pousr]_[A-Za-z0-9_]{20,}|github_pat_[A-Za-z0-9_]{20,})\b`)
)

func Text(value string) string {
	value = strings.ReplaceAll(value, "\x00", "")
	value = githubToken.ReplaceAllString(value, "[redacted-github-token]")
	value = bearer.ReplaceAllString(value, "Bearer [redacted]")
	value = sensitive.ReplaceAllString(value, "$1$2[redacted]")
	value = ipv4.ReplaceAllString(value, "[redacted-ip]")
	value = homePath.ReplaceAllString(value, "/home/[redacted-user]")
	value = mediaPath.ReplaceAllString(value, "/media/[redacted-user]")
	if current, err := user.Current(); err == nil && current.Username != "" {
		value = strings.ReplaceAll(value, current.Username, "[redacted-user]")
	}
	if hostname, err := os.Hostname(); err == nil && hostname != "" {
		value = strings.ReplaceAll(value, hostname, "[redacted-host]")
	}
	return strings.TrimSpace(value)
}

func Args(values []string) []string {
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = Text(value)
	}
	return result
}
