package unilint

// An Issue is a linter issue and contains information on how lo locate the
// in a file.
type Issue struct {
	File    string
	Line    int
	Column  int
	Summary string
}
