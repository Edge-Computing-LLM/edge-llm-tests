package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/model"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/report"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/runner"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/sanitize"
	"github.com/Edge-Computing-LLM/edge-llm-tests/internal/suite"
)

func main() {
	os.Exit(run())
}

func run() int {
	mode := flag.String("mode", "all", "test mode: all, repository, or cluster")
	root := flag.String("root", "..", "parent directory containing the Edge Computing LLM projects")
	output := flag.String("output", "", "result directory; defaults to results/<UTC run ID>")
	includeCandidates := flag.Bool("include-candidates", true, "include non-mutating checks for dashboard template candidates")
	flag.Parse()
	if *mode != "all" && *mode != "repository" && *mode != "cluster" {
		fmt.Fprintln(os.Stderr, "-mode must be all, repository, or cluster")
		return 2
	}
	absoluteRoot, err := filepath.Abs(*root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if err := suite.ValidateRepositoryLayout(absoluteRoot); err != nil {
		fmt.Fprintln(os.Stderr, "project layout:", err)
		return 2
	}
	runID := time.Now().UTC().Format("20060102T150405Z")
	resultDir := *output
	if resultDir == "" {
		resultDir = filepath.Join("results", runID)
	}
	absoluteResultDir, err := filepath.Abs(resultDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if err := os.MkdirAll(absoluteResultDir, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	ctx := context.Background()
	commandRunner := runner.Runner{}
	result := model.Report{SchemaVersion: model.SchemaVersion, RunID: runID, StartedAt: time.Now().UTC(), Mode: *mode, Root: sanitize.Text(absoluteRoot)}
	if *mode == "all" || *mode == "repository" {
		result.Repositories, result.Checks = suite.RepositoryChecks(ctx, absoluteRoot, commandRunner, *includeCandidates)
	}
	if *mode == "all" || *mode == "cluster" {
		environment, checks := suite.ClusterChecks(ctx, absoluteRoot, absoluteResultDir, commandRunner)
		result.Environment = environment
		result.Checks = append(result.Checks, checks...)
	}
	result.Finalize()
	if err := report.Write(absoluteResultDir, result); err != nil {
		fmt.Fprintln(os.Stderr, "write report:", err)
		return 2
	}
	fmt.Printf("report: %s\npassed: %d  failed: %d  skipped: %d  total: %d\n", absoluteResultDir, result.Summary.Passed, result.Summary.Failed, result.Summary.Skipped, result.Summary.Total)
	if result.Summary.Failed > 0 {
		return 1
	}
	return 0
}
