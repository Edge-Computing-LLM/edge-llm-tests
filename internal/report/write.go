package report

import (
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
)

func Write(directory string, result model.Report) error {
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return err
	}
	encoded, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	encoded = append(encoded, '\n')
	if err := os.WriteFile(filepath.Join(directory, "summary.json"), encoded, 0o644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(directory, "report.md"), []byte(markdown(result)), 0o644)
}

func markdown(result model.Report) string {
	var output strings.Builder
	fmt.Fprintf(&output, "# Edge LLM validation report\n\n")
	fmt.Fprintf(&output, "Run `%s` started at %s in `%s` mode. Evidence is sanitized and intentionally excludes credentials, Kubernetes Secrets, kubeconfig content, pod logs, prompts, responses, and model weights.\n\n", result.RunID, result.StartedAt.Format("2006-01-02 15:04:05Z"), result.Mode)
	fmt.Fprintf(&output, "## Outcome\n\n| Passed | Failed | Skipped | Total |\n|---:|---:|---:|---:|\n| %d | %d | %d | %d |\n\n", result.Summary.Passed, result.Summary.Failed, result.Summary.Skipped, result.Summary.Total)
	fmt.Fprintf(&output, "## Environment\n\n| Component | Observed value |\n|---|---|\n")
	environment := [][2]string{{"OS", result.Environment.OS}, {"Kernel", result.Environment.Kernel}, {"Go", result.Environment.Go}, {"k3s", result.Environment.K3s}, {"Kubernetes", result.Environment.Kubernetes}, {"Helm", result.Environment.Helm}, {"GPU", result.Environment.GPU}, {"Allocatable GPUs", fmt.Sprint(result.Environment.GPUAllocatable)}}
	for _, row := range environment {
		fmt.Fprintf(&output, "| %s | %s |\n", row[0], table(row[1]))
	}
	fmt.Fprintf(&output, "\n## Repositories\n\n| Repository | Kind | Branch | Commit | Ahead | Behind | Dirty files |\n|---|---|---|---|---:|---:|---:|\n")
	for _, repository := range result.Repositories {
		fmt.Fprintf(&output, "| %s | %s | `%s` | `%s` | %d | %d | %d |\n", table(repository.Name), table(repository.Kind), table(repository.Branch), table(repository.Commit), repository.Ahead, repository.Behind, repository.DirtyFiles)
	}
	fmt.Fprintf(&output, "\n## Checks\n\n| Status | Scope | Check | Duration |\n|---|---|---|---:|\n")
	for _, check := range result.Checks {
		fmt.Fprintf(&output, "| %s | %s | %s | %s |\n", check.Status, table(check.Scope), table(check.Description), check.Duration.Round(1e6))
	}
	for _, check := range result.Checks {
		if check.Status == model.Pass || (check.Output == "" && check.Reason == "") {
			continue
		}
		fmt.Fprintf(&output, "\n<details>\n<summary>%s — %s</summary>\n\n", check.Status, html.EscapeString(check.ID))
		if check.Reason != "" {
			fmt.Fprintf(&output, "Reason: %s\n\n", html.EscapeString(check.Reason))
		}
		if len(check.Command) > 0 {
			fmt.Fprintf(&output, "Command: `%s`\n\n", html.EscapeString(strings.Join(check.Command, " ")))
		}
		if check.Output != "" {
			fmt.Fprintf(&output, "```text\n%s\n```\n\n", strings.ReplaceAll(check.Output, "```", "` ` `"))
		}
		fmt.Fprintf(&output, "</details>\n")
	}
	return output.String()
}

func table(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "—"
	}
	value = strings.ReplaceAll(value, "|", "\\|")
	value = strings.ReplaceAll(value, "\r", "")
	return strings.ReplaceAll(value, "\n", "<br>")
}
