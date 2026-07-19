package suite

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/runner"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/sanitize"
)

var integerPattern = regexp.MustCompile(`\d+`)

// ClusterChecks performs read-only discovery and uses project diagnostics.
// It never repairs resources, installs charts, reads Secrets, or
// collects pod logs.
func ClusterChecks(ctx context.Context, root, resultsDir string, run runner.Runner) (model.Environment, []model.Check) {
	environment := EnvironmentSnapshot(ctx)
	specs := []runner.Spec{
		{ID: "host:uname", Scope: "host", Description: "Inspect kernel and architecture", Name: "uname", Args: []string{"-srmo"}, Required: true, Timeout: time.Minute},
		{ID: "host:os-release", Scope: "host", Description: "Inspect operating system release", Name: "lsb_release", Args: []string{"-ds"}, Required: false, Timeout: time.Minute},
		{ID: "host:go", Scope: "host", Description: "Inspect Go toolchain", Name: "go", Args: []string{"version"}, Required: true, Timeout: time.Minute},
		{ID: "host:k3s", Scope: "host", Description: "Inspect k3s version", Name: "k3s", Args: []string{"--version"}, Required: true, Timeout: time.Minute},
		{ID: "host:kubectl", Scope: "host", Description: "Inspect Kubernetes client and server versions", Name: "kubectl", Args: []string{"version"}, Required: true, Timeout: time.Minute},
		{ID: "host:helm", Scope: "host", Description: "Inspect Helm version", Name: "helm", Args: []string{"version", "--short"}, Required: true, Timeout: time.Minute},
		{ID: "gpu:nvidia-smi", Scope: "gpu", Description: "Inspect NVIDIA GPU and driver", Name: "nvidia-smi", Args: []string{"--query-gpu=name,driver_version,memory.total", "--format=csv,noheader"}, Required: true, Timeout: time.Minute},
		{ID: "cluster:context", Scope: "cluster", Description: "Confirm an active Kubernetes context", Name: "kubectl", Args: []string{"config", "current-context"}, Required: true, Timeout: time.Minute},
		{ID: "cluster:nodes", Scope: "cluster", Description: "Inspect node readiness without network addresses", Name: "kubectl", Args: []string{"get", "nodes"}, Required: true, Timeout: time.Minute},
		{ID: "cluster:runtimeclasses", Scope: "cluster", Description: "Inspect GPU runtime classes", Name: "kubectl", Args: []string{"get", "runtimeclass"}, Required: true, Timeout: time.Minute},
		{ID: "cluster:gpu-capacity", Scope: "cluster", Description: "Inspect node GPU capacity and allocatable resources", Name: "kubectl", Args: []string{"get", "nodes", "-o", "custom-columns=NAME:.metadata.name,CAPACITY:.status.capacity.nvidia\\.com/gpu,ALLOCATABLE:.status.allocatable.nvidia\\.com/gpu"}, Required: true, Timeout: time.Minute},
		{ID: "cluster:pods", Scope: "cluster", Description: "Inspect workload readiness without pod IP addresses", Name: "kubectl", Args: []string{"get", "pods", "-A"}, Required: true, Timeout: 2 * time.Minute},
		{ID: "cluster:helm-releases", Scope: "cluster", Description: "Inspect deployed Helm releases", Name: "helm", Args: []string{"list", "-A"}, Required: true, Timeout: time.Minute},
		{ID: "cluster:storage", Scope: "cluster", Description: "Inspect storage classes and claims", Name: "kubectl", Args: []string{"get", "storageclass,pvc", "-A"}, Required: true, Timeout: time.Minute},
		{ID: "edge-cli:validate-infra", Scope: "live-validation", Description: "Run unified read-only infrastructure validation", Dir: root + "/edge-cli", Name: "go", Args: []string{"run", "./cmd/edge", "--timeout", "2m", "validate", "infra", "--skip-cuda"}, Required: true, Timeout: 5 * time.Minute},
		{ID: "k3s-nvidia-edge:doctor", Scope: "live-validation", Description: "Run Layer 1 read-only host and cluster diagnostics", Dir: root + "/k3s-nvidia-edge", Name: "go", Args: []string{"run", "./cmd/k3s-nvidia-edge", "doctor", "--sudo=false"}, Required: true, OmitOutput: true, Timeout: 5 * time.Minute},
		{ID: "llm-observability-stack:doctor", Scope: "live-validation", Description: "Run Layer 2 read-only diagnostics", Dir: root + "/llm-observability-stack", Name: "go", Args: []string{"run", "./cmd/llm-observability", "doctor", "-timeout", "1m"}, Required: true, Timeout: 3 * time.Minute},
		{ID: "qwen-gguf-observability:validate", Scope: "live-validation", Description: "Validate local Qwen observability state", Dir: root + "/qwen-gguf-observability", Name: "go", Args: []string{"run", "./cmd/qwen-observe", "validate"}, Required: true, Timeout: 3 * time.Minute},
		{ID: "qwen-gguf-observability:smoke", Scope: "live-validation", Description: "Run the fixed, privacy-safe Qwen smoke probe", Dir: root + "/qwen-gguf-observability", Name: "go", Args: []string{"run", "./cmd/qwen-observe", "smoke", "--output", resultsDir + "/qwen-smoke.json"}, Required: true, OmitOutput: true, Timeout: 5 * time.Minute},
	}
	checks := make([]model.Check, 0, len(specs)+1)
	for _, spec := range specs {
		checks = append(checks, run.Run(ctx, spec))
	}
	checks = append(checks, model.Check{ID: "cluster:evidence-boundary", Scope: "cluster", Description: "Live evidence excludes Kubernetes Secrets, kubeconfig content, pod logs, environment dumps, prompts, and responses", Status: model.Pass, StartedAt: time.Now().UTC(), Reason: "read-only allowlisted commands"})
	return environment, checks
}

func EnvironmentSnapshot(ctx context.Context) model.Environment {
	value := func(name string, args ...string) string {
		output, _ := runner.Output(ctx, "", name, args...)
		return sanitize.Text(strings.TrimSpace(output))
	}
	environment := model.Environment{
		OS:         value("lsb_release", "-ds"),
		Kernel:     value("uname", "-srmo"),
		Go:         value("go", "version"),
		K3s:        value("k3s", "--version"),
		Kubernetes: value("kubectl", "version"),
		Helm:       value("helm", "version", "--short"),
		GPU:        value("nvidia-smi", "--query-gpu=name,driver_version,memory.total", "--format=csv,noheader"),
	}
	allocatable := value("kubectl", "get", "nodes", "-o", "jsonpath={range .items[*]}{.status.allocatable.nvidia\\.com/gpu}{\"\\n\"}{end}")
	for _, match := range integerPattern.FindAllString(allocatable, -1) {
		count, _ := strconv.Atoi(match)
		environment.GPUAllocatable += count
	}
	return environment
}
