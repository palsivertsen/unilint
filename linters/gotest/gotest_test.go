package gotest

import (
	"bytes"
	"testing"

	"github.com/palsivertsen/unilint"
	"github.com/stretchr/testify/assert"
)

func TestLinter_InterpitIssues(t *testing.T) {
	tests := []struct {
		name          string
		errAssertion  assert.ErrorAssertionFunc
		valAsserrtion assert.ComparisonAssertionFunc
		stdOut        string
		issues        []unilint.Issue
	}{
		{
			name:          "subtests",
			errAssertion:  assert.NoError,
			valAsserrtion: assert.Equal,
			stdOut: `?   	coap	[no test files]
--- FAIL: TestOption_SetPath (0.00s)
    --- FAIL: TestOption_SetPath/simple_path (0.00s)
    	options_test.go:58: (/a/b/c): expected /a/b/c, actual /a/a/b/c
FAIL
FAIL	coap/option2	0.004s
?   	coap/option2/example/client	[no test files]
?   	coap/option2/example/dux	[no test files]
?   	coap/option2/example/handler	[no test files]
?   	coap/option2/example/server	[no test files]`,
			issues: []unilint.Issue{
				unilint.Issue{
					File:    "options_test.go",
					Column:  58,
					Summary: "(/a/b/c): expected /a/b/c, actual /a/a/b/c",
				},
			},
		},
		// Panic
		// Normal test
		// Testify
	}
	unit := Linter{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := unit.InterpitIssues(bytes.NewBufferString(tt.stdOut))
			tt.errAssertion(t, err)
			tt.valAsserrtion(t, tt.issues, i)
		})
	}
}
