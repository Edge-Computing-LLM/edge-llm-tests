# Evidence policy

This repository is designed to publish useful engineering evidence without
turning a source repository into a diagnostic-data archive.

## Never persisted

- passwords, API keys, access tokens, bearer credentials, or GitHub tokens;
- Kubernetes Secret data or kubeconfig content;
- environment-variable dumps and private host network addresses;
- pod logs, user prompts, or model responses;
- GGUF files, model weights, backups, downloaded archives, or credential-store
  contents.

The inventory only enters explicitly named Git repositories and the dashboard
candidate directory. It does not recursively inspect adjacent credential,
large-file, model, backup, or archival directories.

## Safe evidence controls

Command output is capped at 64 KiB per check, then sanitized before being added
to a report. Common token formats, sensitive key/value fields, local usernames,
home paths, media paths, and IPv4 addresses are redacted. Commands whose normal
output is an entire rendered manifest are recorded without raw output.

Sanitization is defense in depth, not permission to query sensitive data. The
cluster suite maintains a read-only allowlist and never runs `kubectl get
secrets`, raw kubeconfig inspection, `kubectl logs`, or an environment dump.

Before publishing a new result, reviewers should still inspect it and run the
repository's secret-scanning workflow.
