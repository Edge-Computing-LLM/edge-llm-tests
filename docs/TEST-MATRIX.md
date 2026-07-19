# Test matrix

The harness separates deterministic repository gates from point-in-time live
validation. Repository failures generally indicate source or dependency drift;
live failures describe the state of this particular host and cluster.

## Repository mode

Every organization Go repository receives:

1. `git diff --check`, clean status, and object-connectivity checks.
2. A pure-Go formatting comparison over first-party `.go` files.
3. `go mod verify`, `go test -count=1 ./...`, `go test -race -count=1 ./...`,
   `go vet ./...`, and `govulncheck`.
4. Repository-specific binary and Helm gates.

The Helm matrix renders the Layer 2 default, CPU, local k3s, GeForce 940M, and
full NVIDIA profiles. Rendered Kubernetes manifests are deliberately omitted
from persisted output because values can reference environment-specific names.

Dashboard-template candidates are third-party evaluation material, not
organization products. The harness checks only existing clones and already
installed dependencies. It never runs `npm install`, `pnpm install`, or an
equivalent mutation.

## Cluster mode

Cluster mode uses an allowlist of read-only discovery commands:

- host and tool versions;
- GPU name, driver version, and total memory;
- current context name, node readiness, RuntimeClasses, and GPU allocation;
- workload readiness, Helm release status, storage classes, and claims;
- the read-only diagnostic commands owned by the four project repositories;
- a fixed public Qwen smoke request whose prompt and response are discarded.

The legacy Layer 1 CUDA validator creates a short-lived Kubernetes pod and is
therefore not executed automatically. Its source, unit tests, chart, and doctor
remain covered. Operators can opt into it explicitly from `k3s-nvidia-edge`
when cluster mutation is appropriate.

## Status meaning

| Status | Meaning |
|---|---|
| `PASS` | The required command completed successfully. |
| `FAIL` | A required tool was unavailable, timed out, or returned non-zero. |
| `SKIP` | An optional tool, manifest, dependency set, or script was unavailable. |

Skipped checks are never silently converted to passes.
