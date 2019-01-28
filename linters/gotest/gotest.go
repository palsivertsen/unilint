package gotest

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/palsivertsen/unilint"
)

// A Linter that runs the go test tool
type Linter struct {
}

// Lint the given file
func (l Linter) Lint(ctx context.Context, file string) ([]unilint.Issue, error) {
	// Check if file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist")
	}

	// Build command
	cmd := exec.CommandContext(ctx, "go", "test", file)

	// Buffer the output
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// Fork the command
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	issues, err := l.InterpitIssues(stdOut)
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return issues, err
}

// InterpitIssues from go test tool output
func (l Linter) InterpitIssues(r io.Reader) ([]unilint.Issue, error) {
	panic("Not implemented")
}
