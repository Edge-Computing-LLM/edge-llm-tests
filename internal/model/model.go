package model

import "time"

const SchemaVersion = "1.0"

type Status string

const (
	Pass Status = "PASS"
	Fail Status = "FAIL"
	Skip Status = "SKIP"
)

type Repository struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Origin     string `json:"origin,omitempty"`
	Branch     string `json:"branch,omitempty"`
	Commit     string `json:"commit,omitempty"`
	Upstream   string `json:"upstream,omitempty"`
	Ahead      int    `json:"ahead"`
	Behind     int    `json:"behind"`
	DirtyFiles int    `json:"dirty_files"`
	Kind       string `json:"kind"`
}

type Check struct {
	ID              string        `json:"id"`
	Scope           string        `json:"scope"`
	Description     string        `json:"description"`
	Status          Status        `json:"status"`
	Command         []string      `json:"command,omitempty"`
	StartedAt       time.Time     `json:"started_at"`
	Duration        time.Duration `json:"duration"`
	ExitCode        int           `json:"exit_code,omitempty"`
	Output          string        `json:"output,omitempty"`
	OutputTruncated bool          `json:"output_truncated,omitempty"`
	Reason          string        `json:"reason,omitempty"`
}

type Environment struct {
	OS             string `json:"os,omitempty"`
	Kernel         string `json:"kernel,omitempty"`
	Go             string `json:"go,omitempty"`
	K3s            string `json:"k3s,omitempty"`
	Kubernetes     string `json:"kubernetes,omitempty"`
	Helm           string `json:"helm,omitempty"`
	GPU            string `json:"gpu,omitempty"`
	GPUAllocatable int    `json:"gpu_allocatable"`
}

type Summary struct {
	Passed  int `json:"passed"`
	Failed  int `json:"failed"`
	Skipped int `json:"skipped"`
	Total   int `json:"total"`
}

type Report struct {
	SchemaVersion string       `json:"schema_version"`
	RunID         string       `json:"run_id"`
	StartedAt     time.Time    `json:"started_at"`
	FinishedAt    time.Time    `json:"finished_at"`
	Mode          string       `json:"mode"`
	Root          string       `json:"root"`
	Environment   Environment  `json:"environment"`
	Repositories  []Repository `json:"repositories"`
	Checks        []Check      `json:"checks"`
	Summary       Summary      `json:"summary"`
}

func (r *Report) Finalize() {
	r.Summary = Summary{Total: len(r.Checks)}
	for _, check := range r.Checks {
		switch check.Status {
		case Pass:
			r.Summary.Passed++
		case Fail:
			r.Summary.Failed++
		case Skip:
			r.Summary.Skipped++
		}
	}
	r.FinishedAt = time.Now().UTC()
}
