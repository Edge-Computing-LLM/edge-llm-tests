# Local Ubuntu + k3s + NVIDIA validation

This test plane targets a single-node local edge environment. It distinguishes
three truths that should not be conflated:

1. Source and chart quality: deterministic repository gates can pass even while
   a live cluster is unhealthy.
2. Host capability: a working `nvidia-smi` proves the host driver can see the
   GPU, but not that Kubernetes can schedule it.
3. Cluster availability: node readiness, GPU allocatable resources, workloads,
   and the Qwen contract are point-in-time observations.

Changing a host network adapter can leave a single-node k3s installation bound
to a stale node address or Flannel interface. Typical symptoms include an
unhealthy node, unavailable Node Feature Discovery or NVIDIA components,
missing GPU allocation, and downstream observability or smoke failures. The
harness reports those symptoms; repair belongs in the Layer 1 runbook and must
be performed intentionally by the operator.

After any repair, rerun both modes rather than treating an earlier report as
current:

```bash
go run ./cmd/edge-llm-tests -mode repository -root ..
go run ./cmd/edge-llm-tests -mode cluster -root ..
```

Reports are immutable point-in-time evidence. Do not overwrite a prior run;
create a new UTC-stamped directory so the change is auditable.
