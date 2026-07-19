package sanitize

import (
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
	return strings.TrimSpace(value)
}

func Args(values []string) []string {
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = Text(value)
	}
	return result
}
