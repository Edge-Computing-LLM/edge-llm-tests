package runner

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/sanitize"
)

const DefaultOutputLimit = 64 * 1024

type Spec struct {
	ID          string
	Scope       string
	Description string
	Dir         string
	Name        string
	Args        []string
	Required    bool
	OmitOutput  bool
	Timeout     time.Duration
}

type Runner struct {
	OutputLimit int
}

type cappedBuffer struct {
	data      []byte
	limit     int
	truncated bool
}

func (b *cappedBuffer) Write(data []byte) (int, error) {
	available := b.limit - len(b.data)
	if available > 0 {
		if available > len(data) {
			available = len(data)
		}
		b.data = append(b.data, data[:available]...)
	}
	if available < len(data) {
		b.truncated = true
	}
	return len(data), nil
}

func (r Runner) Run(ctx context.Context, spec Spec) model.Check {
	started := time.Now().UTC()
	result := model.Check{ID: spec.ID, Scope: spec.Scope, Description: spec.Description, StartedAt: started, Command: sanitize.Args(append([]string{spec.Name}, spec.Args...))}
	if _, err := exec.LookPath(spec.Name); err != nil {
		result.Duration = time.Since(started)
		result.ExitCode = -1
		result.Reason = "required command is unavailable"
		if spec.Required {
			result.Status = model.Fail
		} else {
			result.Status = model.Skip
		}
		return result
	}
	timeout := spec.Timeout
	if timeout <= 0 {
		timeout = 10 * time.Minute
	}
	commandCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	command := exec.CommandContext(commandCtx, spec.Name, spec.Args...)
	command.Dir = spec.Dir
	command.Env = append(os.Environ(), "LC_ALL=C", "LANG=C")
	limit := r.OutputLimit
	if limit <= 0 {
		limit = DefaultOutputLimit
	}
	output := &cappedBuffer{limit: limit}
	command.Stdout = output
	command.Stderr = output
	err := command.Run()
	result.Duration = time.Since(started)
	result.OutputTruncated = output.truncated
	if !spec.OmitOutput {
		result.Output = sanitize.Text(string(output.data))
	} else if len(output.data) > 0 {
		result.Output = "command output intentionally omitted by evidence policy"
	}
	if err == nil {
		result.Status = model.Pass
		return result
	}
	result.Status = model.Fail
	result.ExitCode = -1
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		result.ExitCode = exitErr.ExitCode()
	}
	if errors.Is(commandCtx.Err(), context.DeadlineExceeded) {
		result.Reason = "command timed out after " + timeout.String()
	} else {
		result.Reason = "command exited with status " + strconv.Itoa(result.ExitCode)
	}
	return result
}

func Output(ctx context.Context, dir, name string, args ...string) (string, error) {
	command := exec.CommandContext(ctx, name, args...)
	command.Dir = dir
	command.Env = append(os.Environ(), "LC_ALL=C", "LANG=C")
	data, err := command.CombinedOutput()
	return strings.TrimSpace(string(data)), err
}
