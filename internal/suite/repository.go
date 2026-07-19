package suite

import (
	"context"
	"encoding/json"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/runner"
)

var organizationRepos = []string{"edge-cli", "k3s-nvidia-edge", "llm-observability-stack", "qwen-gguf-observability"}

func RepositoryChecks(ctx context.Context, root string, run runner.Runner, includeCandidates bool) ([]model.Repository, []model.Check) {
	buildDir, err := os.MkdirTemp("", "edge-llm-tests-builds-")
	if err == nil {
		defer os.RemoveAll(buildDir)
	} else {
		buildDir = os.TempDir()
	}
	var repositories []model.Repository
	var checks []model.Check
	for _, name := range organizationRepos {
		path := filepath.Join(root, name)
		repositories = append(repositories, InspectRepository(ctx, root, path, "organization"))
		checks = append(checks, gitChecks(ctx, path, name, run)...)
		checks = append(checks, goFormattingCheck(path, name))
		checks = append(checks, repositoryPlan(name, path, buildDir, run, ctx)...)
	}
	if includeCandidates {
		candidateRoot := filepath.Join(root, "dashboard-template-candidates")
		entries, _ := os.ReadDir(candidateRoot)
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			path := filepath.Join(candidateRoot, entry.Name())
			if _, err := os.Stat(filepath.Join(path, ".git")); err != nil {
				continue
			}
			scope := "dashboard-template-candidates/" + entry.Name()
			repositories = append(repositories, InspectRepository(ctx, root, path, "candidate"))
			checks = append(checks, gitChecks(ctx, path, scope, run)...)
			checks = append(checks, candidateChecks(ctx, path, scope, run)...)
		}
	}
	checks = append(checks, model.Check{
		ID:          "collection:evidence-boundary",
		Scope:       "collection",
		Description: "Credential, model-weight, backup, and large-file directory contents are excluded",
		Status:      model.Pass,
		StartedAt:   time.Now().UTC(),
		Reason:      "metadata-only inventory; sensitive directory contents were not opened",
	})
	return repositories, checks
}

func InspectRepository(ctx context.Context, root, path, kind string) model.Repository {
	result := model.Repository{Name: filepath.Base(path), Path: relative(root, path), Kind: kind}
	result.Origin, _ = runner.Output(ctx, path, "git", "remote", "get-url", "origin")
	result.Branch, _ = runner.Output(ctx, path, "git", "branch", "--show-current")
	result.Commit, _ = runner.Output(ctx, path, "git", "rev-parse", "--short=12", "HEAD")
	result.Upstream, _ = runner.Output(ctx, path, "git", "rev-parse", "--abbrev-ref", "@{upstream}")
	status, _ := runner.Output(ctx, path, "git", "status", "--porcelain")
	if status != "" {
		result.DirtyFiles = len(strings.Split(status, "\n"))
	}
	if result.Upstream != "" {
		counts, err := runner.Output(ctx, path, "git", "rev-list", "--left-right", "--count", "HEAD...@{upstream}")
		if err == nil {
			fields := strings.Fields(counts)
			if len(fields) == 2 {
				result.Ahead, _ = strconv.Atoi(fields[0])
				result.Behind, _ = strconv.Atoi(fields[1])
			}
		}
	}
	return result
}

func gitChecks(ctx context.Context, path, scope string, run runner.Runner) []model.Check {
	checks := []model.Check{
		run.Run(ctx, runner.Spec{ID: scope + ":git-clean", Scope: scope, Description: "Git worktree is clean", Dir: path, Name: "git", Args: []string{"diff", "--check"}, Required: true, Timeout: time.Minute}),
		run.Run(ctx, runner.Spec{ID: scope + ":git-fsck", Scope: scope, Description: "Git object connectivity", Dir: path, Name: "git", Args: []string{"fsck", "--connectivity-only"}, Required: true, Timeout: 2 * time.Minute}),
	}
	started := time.Now().UTC()
	status, err := runner.Output(ctx, path, "git", "status", "--porcelain")
	clean := model.Check{ID: scope + ":git-status", Scope: scope, Description: "Git worktree has no tracked or untracked changes", StartedAt: started, Status: model.Pass, Duration: time.Since(started)}
	if err != nil {
		clean.Status = model.Fail
		clean.Reason = "unable to inspect Git status"
	} else if status != "" {
		clean.Status = model.Fail
		clean.Reason = "worktree contains changes"
		clean.Output = status
	}
	return append(checks, clean)
}

func repositoryPlan(name, path, buildDir string, run runner.Runner, ctx context.Context) []model.Check {
	goChecks := []runner.Spec{
		{ID: name + ":go-mod-verify", Scope: name, Description: "Verify Go module downloads", Dir: path, Name: "go", Args: []string{"mod", "verify"}, Required: true},
		{ID: name + ":go-test", Scope: name, Description: "Run Go unit tests", Dir: path, Name: "go", Args: []string{"test", "-count=1", "./..."}, Required: true},
		{ID: name + ":go-race", Scope: name, Description: "Run Go race tests", Dir: path, Name: "go", Args: []string{"test", "-race", "-count=1", "./..."}, Required: true, Timeout: 15 * time.Minute},
		{ID: name + ":go-vet", Scope: name, Description: "Run Go vet", Dir: path, Name: "go", Args: []string{"vet", "./..."}, Required: true},
		{ID: name + ":govulncheck", Scope: name, Description: "Check reachable Go vulnerabilities", Dir: path, Name: "go", Args: []string{"run", "golang.org/x/vuln/cmd/govulncheck@latest", "./..."}, Required: true, Timeout: 15 * time.Minute},
	}
	var specs []runner.Spec
	specs = append(specs, goChecks...)
	switch name {
	case "edge-cli":
		specs = append(specs, runner.Spec{ID: name + ":build", Scope: name, Description: "Build edge CLI", Dir: path, Name: "go", Args: []string{"build", "-o", filepath.Join(buildDir, "edge"), "./cmd/edge"}, Required: true})
	case "k3s-nvidia-edge":
		specs = append(specs,
			runner.Spec{ID: name + ":build", Scope: name, Description: "Build Layer 1 CLI", Dir: path, Name: "go", Args: []string{"build", "-o", filepath.Join(buildDir, "k3s-nvidia-edge"), "./cmd/k3s-nvidia-edge"}, Required: true},
			runner.Spec{ID: name + ":helm-lint", Scope: name, Description: "Lint Layer 1 Helm chart", Dir: path, Name: "helm", Args: []string{"lint", "charts/k3s-nvidia-edge"}, Required: true},
			runner.Spec{ID: name + ":helm-render", Scope: name, Description: "Render Layer 1 Helm chart", Dir: path, Name: "helm", Args: []string{"template", "k3s-nvidia-edge", "charts/k3s-nvidia-edge"}, Required: true, OmitOutput: true},
		)
	case "llm-observability-stack":
		specs = append(specs,
			runner.Spec{ID: name + ":build-cli", Scope: name, Description: "Build Layer 2 CLI", Dir: path, Name: "go", Args: []string{"build", "-o", filepath.Join(buildDir, "llm-observability"), "./cmd/llm-observability"}, Required: true},
			runner.Spec{ID: name + ":build-gateway", Scope: name, Description: "Build Ollama gateway", Dir: path, Name: "go", Args: []string{"build", "-o", filepath.Join(buildDir, "ollama-gateway"), "./cmd/ollama-gateway"}, Required: true},
			runner.Spec{ID: name + ":build-toolbox", Scope: name, Description: "Build edge toolbox", Dir: path, Name: "go", Args: []string{"build", "-o", filepath.Join(buildDir, "edge-toolbox"), "./cmd/edge-toolbox"}, Required: true},
			runner.Spec{ID: name + ":helm-dependencies", Scope: name, Description: "Resolve vendored chart dependencies", Dir: path, Name: "helm", Args: []string{"dependency", "build", "."}, Required: true, Timeout: 15 * time.Minute},
			runner.Spec{ID: name + ":helm-lint", Scope: name, Description: "Lint Layer 2 Helm chart", Dir: path, Name: "helm", Args: []string{"lint", "."}, Required: true},
		)
		profiles := []struct {
			id   string
			args []string
		}{
			{"default", nil},
			{"cpu", []string{"-f", "values.cpu-k3s.yaml"}},
			{"local", []string{"-f", "values.local-k3s.example.yaml"}},
			{"geforce", []string{"-f", "values.geforce-940m-k3s.yaml"}},
			{"full-nvidia", []string{"-f", "values.full-stack-nvidia.example.yaml", "--set", "langsmith.existingSecret=", "--set", "openWebUI.existingSecret=", "--set", "open-webui.webuiSecret.existingSecretName="}},
		}
		for _, profile := range profiles {
			args := append([]string{"template", "llm-observability-stack", "."}, profile.args...)
			specs = append(specs, runner.Spec{ID: name + ":helm-render-" + profile.id, Scope: name, Description: "Render " + profile.id + " Helm profile", Dir: path, Name: "helm", Args: args, Required: true, OmitOutput: true})
		}
	case "qwen-gguf-observability":
		specs = append(specs, runner.Spec{ID: name + ":build", Scope: name, Description: "Build read-only Qwen observer", Dir: path, Name: "go", Args: []string{"build", "-o", filepath.Join(buildDir, "qwen-observe"), "./cmd/qwen-observe"}, Required: true})
	}
	checks := make([]model.Check, 0, len(specs))
	for _, spec := range specs {
		checks = append(checks, run.Run(ctx, spec))
	}
	return checks
}

func goFormattingCheck(path, scope string) model.Check {
	started := time.Now().UTC()
	result := model.Check{ID: scope + ":gofmt", Scope: scope, Description: "Go sources are formatted", StartedAt: started, Status: model.Pass}
	var unformatted []string
	err := filepath.WalkDir(path, func(file string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			switch entry.Name() {
			case ".git", "charts", "vendor", "node_modules", "results", "artifacts":
				if file != path {
					return filepath.SkipDir
				}
			}
			return nil
		}
		if filepath.Ext(file) != ".go" {
			return nil
		}
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		formatted, err := format.Source(data)
		if err != nil {
			return err
		}
		if string(formatted) != string(data) {
			unformatted = append(unformatted, relative(path, file))
		}
		return nil
	})
	result.Duration = time.Since(started)
	if err != nil {
		result.Status = model.Fail
		result.Reason = err.Error()
	} else if len(unformatted) > 0 {
		result.Status = model.Fail
		result.Output = strings.Join(unformatted, "\n")
		result.Reason = "unformatted Go files"
	}
	return result
}

type packageManifest struct {
	Scripts map[string]string `json:"scripts"`
}

func candidateChecks(ctx context.Context, path, scope string, run runner.Runner) []model.Check {
	started := time.Now().UTC()
	data, err := os.ReadFile(filepath.Join(path, "package.json"))
	if err != nil {
		return []model.Check{{ID: scope + ":package", Scope: scope, Description: "Inspect frontend test scripts", Status: model.Skip, StartedAt: started, Reason: "package.json is unavailable"}}
	}
	var manifest packageManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return []model.Check{{ID: scope + ":package", Scope: scope, Description: "Inspect frontend test scripts", Status: model.Fail, StartedAt: started, Reason: "invalid package.json"}}
	}
	if _, err := os.Stat(filepath.Join(path, "node_modules")); err != nil {
		return []model.Check{{ID: scope + ":frontend", Scope: scope, Description: "Run available frontend lint/test/build scripts", Status: model.Skip, StartedAt: started, Reason: "dependencies are not installed; audit does not mutate third-party candidates"}}
	}
	manager := "npm"
	if _, err := os.Stat(filepath.Join(path, "pnpm-lock.yaml")); err == nil {
		manager = "pnpm"
	} else if _, err := os.Stat(filepath.Join(path, "yarn.lock")); err == nil {
		manager = "yarn"
	}
	var checks []model.Check
	for _, script := range []string{"lint", "test", "build"} {
		if _, ok := manifest.Scripts[script]; !ok {
			continue
		}
		checks = append(checks, run.Run(ctx, runner.Spec{ID: scope + ":" + script, Scope: scope, Description: "Run frontend " + script + " script", Dir: path, Name: manager, Args: []string{"run", script}, Required: false, Timeout: 15 * time.Minute}))
	}
	if len(checks) == 0 {
		checks = append(checks, model.Check{ID: scope + ":frontend", Scope: scope, Description: "Run available frontend scripts", Status: model.Skip, StartedAt: started, Reason: "package.json defines no lint, test, or build script"})
	}
	return checks
}

func relative(root, value string) string {
	rel, err := filepath.Rel(root, value)
	if err != nil {
		return filepath.Base(value)
	}
	return filepath.ToSlash(rel)
}

func ValidateRepositoryLayout(root string) error {
	for _, name := range organizationRepos {
		for _, required := range []string{".git", "go.mod", "AGENTS.md"} {
			if _, err := os.Stat(filepath.Join(root, name, required)); err != nil {
				return fmt.Errorf("%s is missing %s", name, required)
			}
		}
	}
	return nil
}
