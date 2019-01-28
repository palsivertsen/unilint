package unilint

import "context"

// Linter executes a liter and returns a list of issues. A Linter must abort
// execution if the context is cancled and return ctx.Err()
type Linter interface {
	Lint(ctx context.Context, file string) ([]Issue, error)
}
