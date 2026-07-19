# Edge LLM validation report

Run `20260719T043235Z` started at 2026-07-19 04:32:35Z in `all` mode. Evidence is sanitized and intentionally excludes credentials, Kubernetes Secrets, kubeconfig content, pod logs, prompts, responses, and model weights.

## Outcome

| Passed | Failed | Skipped | Total |
|---:|---:|---:|---:|
| 89 | 4 | 7 | 100 |

## Environment

| Component | Observed value |
|---|---|
| OS | Ubuntu 24.04.3 LTS |
| Kernel | Linux 6.17.0-40-generic x86_64 GNU/Linux |
| Go | go version go1.26.5 linux/amd64 |
| k3s | k3s version v1.36.2+k3s1 (01b6f04a)<br>go version go1.26.4 |
| Kubernetes | Client Version: v1.36.2+k3s1<br>Kustomize Version: v5.8.1<br>Server Version: v1.36.2+k3s1 |
| Helm | v4.2.3+g43e8b7f |
| GPU | NVIDIA GeForce 940M, 580.95.05, 1024 MiB |
| Allocatable GPUs | 1 |

## Repositories

| Repository | Kind | Branch | Commit | Ahead | Behind | Dirty files |
|---|---|---|---|---:|---:|---:|
| edge-cli | organization | `codex/deep-audit-ubuntu-k3s-nvidia-20260719` | `59924f3798ad` | 0 | 0 | 0 |
| k3s-nvidia-edge | organization | `codex/deep-audit-ubuntu-k3s-nvidia-20260719` | `95f5e6cb9c29` | 0 | 0 | 0 |
| llm-observability-stack | organization | `codex/deep-audit-ubuntu-k3s-nvidia-20260719` | `a1b1bbfb77d9` | 0 | 0 | 0 |
| qwen-gguf-observability | organization | `codex/deep-audit-ubuntu-k3s-nvidia-20260719` | `07a3d55d4a2b` | 0 | 0 | 0 |
| grafana-faro-nextjs-example | candidate | `main` | `b9547013578e` | 0 | 0 | 0 |
| material-kit-react | candidate | `main` | `6423e1a2f092` | 0 | 0 | 0 |
| next-shadcn-admin-dashboard | candidate | `main` | `43f47b8b7022` | 0 | 18 | 0 |
| next-shadcn-dashboard-starter | candidate | `main` | `21848bd25125` | 0 | 86 | 0 |
| nextadmin-dashboard | candidate | `main` | `8913468e1d00` | 0 | 1 | 0 |
| rh-ai-observability-summarizer | candidate | `main` | `048f784a19d0` | 0 | 0 | 0 |
| tailadmin-nextjs | candidate | `main` | `d3526b35fb7e` | 0 | 0 | 0 |

## Checks

| Status | Scope | Check | Duration |
|---|---|---|---:|
| PASS | edge-cli | Git worktree is clean | 3ms |
| PASS | edge-cli | Git object connectivity | 81ms |
| PASS | edge-cli | Git worktree has no tracked or untracked changes | 4ms |
| PASS | edge-cli | Go sources are formatted | 14ms |
| PASS | edge-cli | Verify Go module downloads | 53ms |
| PASS | edge-cli | Run Go unit tests | 1.403s |
| PASS | edge-cli | Run Go race tests | 4.722s |
| PASS | edge-cli | Run Go vet | 297ms |
| PASS | edge-cli | Check reachable Go vulnerabilities | 6.516s |
| PASS | edge-cli | Build edge CLI | 707ms |
| PASS | k3s-nvidia-edge | Git worktree is clean | 3ms |
| PASS | k3s-nvidia-edge | Git object connectivity | 120ms |
| PASS | k3s-nvidia-edge | Git worktree has no tracked or untracked changes | 3ms |
| PASS | k3s-nvidia-edge | Go sources are formatted | 6ms |
| PASS | k3s-nvidia-edge | Verify Go module downloads | 5ms |
| PASS | k3s-nvidia-edge | Run Go unit tests | 278ms |
| PASS | k3s-nvidia-edge | Run Go race tests | 1.411s |
| PASS | k3s-nvidia-edge | Run Go vet | 118ms |
| PASS | k3s-nvidia-edge | Check reachable Go vulnerabilities | 4.606s |
| PASS | k3s-nvidia-edge | Build Layer 1 CLI | 473ms |
| PASS | k3s-nvidia-edge | Lint Layer 1 Helm chart | 103ms |
| PASS | k3s-nvidia-edge | Render Layer 1 Helm chart | 154ms |
| PASS | llm-observability-stack | Git worktree is clean | 13ms |
| PASS | llm-observability-stack | Git object connectivity | 175ms |
| PASS | llm-observability-stack | Git worktree has no tracked or untracked changes | 7ms |
| PASS | llm-observability-stack | Go sources are formatted | 15ms |
| PASS | llm-observability-stack | Verify Go module downloads | 3.064s |
| PASS | llm-observability-stack | Run Go unit tests | 4.233s |
| PASS | llm-observability-stack | Run Go race tests | 8.533s |
| PASS | llm-observability-stack | Run Go vet | 396ms |
| PASS | llm-observability-stack | Check reachable Go vulnerabilities | 9.198s |
| PASS | llm-observability-stack | Build Layer 2 CLI | 637ms |
| PASS | llm-observability-stack | Build Ollama gateway | 1.27s |
| PASS | llm-observability-stack | Build edge toolbox | 1.312s |
| PASS | llm-observability-stack | Resolve vendored chart dependencies | 2.05s |
| PASS | llm-observability-stack | Lint Layer 2 Helm chart | 374ms |
| PASS | llm-observability-stack | Render default Helm profile | 203ms |
| PASS | llm-observability-stack | Render cpu Helm profile | 616ms |
| PASS | llm-observability-stack | Render local Helm profile | 204ms |
| PASS | llm-observability-stack | Render geforce Helm profile | 364ms |
| PASS | llm-observability-stack | Render full-nvidia Helm profile | 461ms |
| PASS | qwen-gguf-observability | Git worktree is clean | 2ms |
| PASS | qwen-gguf-observability | Git object connectivity | 15ms |
| PASS | qwen-gguf-observability | Git worktree has no tracked or untracked changes | 2ms |
| PASS | qwen-gguf-observability | Go sources are formatted | 6ms |
| PASS | qwen-gguf-observability | Verify Go module downloads | 6ms |
| PASS | qwen-gguf-observability | Run Go unit tests | 284ms |
| PASS | qwen-gguf-observability | Run Go race tests | 1.476s |
| PASS | qwen-gguf-observability | Run Go vet | 128ms |
| PASS | qwen-gguf-observability | Check reachable Go vulnerabilities | 4.376s |
| PASS | qwen-gguf-observability | Build read-only Qwen observer | 369ms |
| PASS | dashboard-template-candidates/grafana-faro-nextjs-example | Git worktree is clean | 2ms |
| PASS | dashboard-template-candidates/grafana-faro-nextjs-example | Git object connectivity | 15ms |
| PASS | dashboard-template-candidates/grafana-faro-nextjs-example | Git worktree has no tracked or untracked changes | 5ms |
| SKIP | dashboard-template-candidates/grafana-faro-nextjs-example | Run available frontend lint/test/build scripts | 0s |
| PASS | dashboard-template-candidates/material-kit-react | Git worktree is clean | 2ms |
| PASS | dashboard-template-candidates/material-kit-react | Git object connectivity | 12ms |
| PASS | dashboard-template-candidates/material-kit-react | Git worktree has no tracked or untracked changes | 3ms |
| SKIP | dashboard-template-candidates/material-kit-react | Run available frontend lint/test/build scripts | 0s |
| PASS | dashboard-template-candidates/next-shadcn-admin-dashboard | Git worktree is clean | 3ms |
| PASS | dashboard-template-candidates/next-shadcn-admin-dashboard | Git object connectivity | 14ms |
| PASS | dashboard-template-candidates/next-shadcn-admin-dashboard | Git worktree has no tracked or untracked changes | 6ms |
| SKIP | dashboard-template-candidates/next-shadcn-admin-dashboard | Run available frontend lint/test/build scripts | 0s |
| PASS | dashboard-template-candidates/next-shadcn-dashboard-starter | Git worktree is clean | 4ms |
| PASS | dashboard-template-candidates/next-shadcn-dashboard-starter | Git object connectivity | 14ms |
| PASS | dashboard-template-candidates/next-shadcn-dashboard-starter | Git worktree has no tracked or untracked changes | 6ms |
| SKIP | dashboard-template-candidates/next-shadcn-dashboard-starter | Run available frontend lint/test/build scripts | 0s |
| PASS | dashboard-template-candidates/nextadmin-dashboard | Git worktree is clean | 4ms |
| PASS | dashboard-template-candidates/nextadmin-dashboard | Git object connectivity | 11ms |
| PASS | dashboard-template-candidates/nextadmin-dashboard | Git worktree has no tracked or untracked changes | 7ms |
| SKIP | dashboard-template-candidates/nextadmin-dashboard | Run available frontend lint/test/build scripts | 0s |
| PASS | dashboard-template-candidates/rh-ai-observability-summarizer | Git worktree is clean | 4ms |
| PASS | dashboard-template-candidates/rh-ai-observability-summarizer | Git object connectivity | 17ms |
| PASS | dashboard-template-candidates/rh-ai-observability-summarizer | Git worktree has no tracked or untracked changes | 6ms |
| SKIP | dashboard-template-candidates/rh-ai-observability-summarizer | Inspect frontend test scripts | 0s |
| PASS | dashboard-template-candidates/tailadmin-nextjs | Git worktree is clean | 3ms |
| PASS | dashboard-template-candidates/tailadmin-nextjs | Git object connectivity | 13ms |
| PASS | dashboard-template-candidates/tailadmin-nextjs | Git worktree has no tracked or untracked changes | 4ms |
| SKIP | dashboard-template-candidates/tailadmin-nextjs | Run available frontend lint/test/build scripts | 0s |
| PASS | collection | Credential, model-weight, backup, and large-file directory contents are excluded | 0s |
| PASS | host | Inspect kernel and architecture | 1ms |
| PASS | host | Inspect operating system release | 12ms |
| PASS | host | Inspect Go toolchain | 8ms |
| PASS | host | Inspect k3s version | 13ms |
| PASS | host | Inspect Kubernetes client and server versions | 166ms |
| PASS | host | Inspect Helm version | 57ms |
| PASS | gpu | Inspect NVIDIA GPU and driver | 38ms |
| PASS | cluster | Confirm an active Kubernetes context | 140ms |
| PASS | cluster | Inspect node readiness without network addresses | 153ms |
| PASS | cluster | Inspect GPU runtime classes | 165ms |
| PASS | cluster | Inspect node GPU capacity and allocatable resources | 177ms |
| PASS | cluster | Inspect workload readiness without pod IP addresses | 271ms |
| PASS | cluster | Inspect deployed Helm releases | 383ms |
| PASS | cluster | Inspect storage classes and claims | 216ms |
| FAIL | live-validation | Run unified read-only infrastructure validation | 1.682s |
| FAIL | live-validation | Run Layer 1 read-only host and cluster diagnostics | 1.653s |
| FAIL | live-validation | Run Layer 2 read-only diagnostics | 538ms |
| FAIL | live-validation | Validate local Qwen observability state | 2.776s |
| PASS | live-validation | Run the fixed, privacy-safe Qwen smoke probe | 1.957s |
| PASS | cluster | Live evidence excludes Kubernetes Secrets, kubeconfig content, pod logs, environment dumps, prompts, and responses | 0s |

<details>
<summary>SKIP — dashboard-template-candidates/grafana-faro-nextjs-example:frontend</summary>

Reason: dependencies are not installed; audit does not mutate third-party candidates

</details>

<details>
<summary>SKIP — dashboard-template-candidates/material-kit-react:frontend</summary>

Reason: dependencies are not installed; audit does not mutate third-party candidates

</details>

<details>
<summary>SKIP — dashboard-template-candidates/next-shadcn-admin-dashboard:frontend</summary>

Reason: dependencies are not installed; audit does not mutate third-party candidates

</details>

<details>
<summary>SKIP — dashboard-template-candidates/next-shadcn-dashboard-starter:frontend</summary>

Reason: dependencies are not installed; audit does not mutate third-party candidates

</details>

<details>
<summary>SKIP — dashboard-template-candidates/nextadmin-dashboard:frontend</summary>

Reason: dependencies are not installed; audit does not mutate third-party candidates

</details>

<details>
<summary>SKIP — dashboard-template-candidates/rh-ai-observability-summarizer:package</summary>

Reason: package.json is unavailable

</details>

<details>
<summary>SKIP — dashboard-template-candidates/tailadmin-nextjs:frontend</summary>

Reason: dependencies are not installed; audit does not mutate third-party candidates

</details>

<details>
<summary>FAIL — edge-cli:validate-infra</summary>

Reason: command exited with status 1

Command: `go run ./cmd/edge --timeout 2m validate infra --skip-cuda`

```text
[ok] kubectl cluster-info
  Kubernetes control plane is running at https://[redacted-ip]:6443
  CoreDNS is running at https://[redacted-ip]:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
  Metrics-server is running at https://[redacted-ip]:6443/api/v1/namespaces/kube-system/services/https:metrics-server:https/proxy
  
  To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
[ok] cluster nodes
  NAME                      STATUS   ROLES           AGE     VERSION        INTERNAL-IP     EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION              CONTAINER-RUNTIME
  [redacted-user]-thinkpad-t450s   Ready    control-plane   2d15h   v1.36.2+k3s1   [redacted-ip]   <none>        Ubuntu 24.04.3 LTS   6.17.0-40-generic (amd64)   containerd://2.3.2-k3s2
[fail] k3s node address
[ok] NVIDIA RuntimeClass
  NAME     HANDLER   AGE
  nvidia   nvidia    21h
[fail] gpu-operator pod health
  8 active, 1 completed
[ok] GPU allocatable
  1
Error: 2 check(s) failed: k3s node address, gpu-operator pod health
Usage:
  edge validate infra [flags]

Flags:
  -h, --help        help for infra
      --skip-cuda   skip the CUDA pod when the GPU is occupied by a workload

Global Flags:
      --config string    config file path
      --dry-run          print mutating operations without executing them
      --timeout string   overall command timeout (default "10m")
      --verbose          print command details

2 check(s) failed: k3s node address, gpu-operator pod health
exit status 1
```

</details>

<details>
<summary>FAIL — k3s-nvidia-edge:doctor</summary>

Reason: command exited with status 1

Command: `go run ./cmd/k3s-nvidia-edge doctor --sudo=false`

```text
command output intentionally omitted by evidence policy
```

</details>

<details>
<summary>FAIL — llm-observability-stack:doctor</summary>

Reason: command exited with status 1

Command: `go run ./cmd/llm-observability doctor -timeout 1m`

```text
[run] Kubernetes connectivity
Kubernetes control plane is running at https://[redacted-ip]:6443
CoreDNS is running at https://[redacted-ip]:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://[redacted-ip]:6443/api/v1/namespaces/kube-system/services/https:metrics-server:https/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
[run] k3s node address
k3s node InternalIP is not assigned to a current host interface
check node-ip and flannel-iface in /etc/rancher/k3s/config.yaml, then restart k3s
error: k3s node address failed: exit status 1
exit status 1
```

</details>

<details>
<summary>FAIL — qwen-gguf-observability:validate</summary>

Reason: command exited with status 1

Command: `go run ./cmd/qwen-observe validate`

```text
[PASS] kubernetes-node-ready: all observed nodes are Ready
[PASS] nvidia-gpu-allocatable: at least one nvidia.com/gpu is allocatable
[PASS] nvidia-runtimeclass: RuntimeClass/nvidia exists
[PASS] helm-release-deployed: LLM stack Helm release is deployed
[FAIL] workloads-ready: all observed application containers are Ready
[PASS] qwen-registered: expected local Qwen alias is registered
[PASS] qwen-resident: expected local Qwen alias is loaded
[PASS] qwen-gpu-active: Ollama reports GPU participation
[PASS] qwen-keep-alive: Ollama reports Until=Forever
[PASS] vram-ceiling: GPU memory stays at or below 850 MiB
[PASS] num-gpu-layers: num_gpu is 23
[PASS] context-window: num_ctx is 256
[PASS] batch-size: num_batch is 1

12 passed, 1 failed
error: runtime contract failed
exit status 1
```

</details>
