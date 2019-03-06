package prompter

import (
	"bufio"
	"io"
)

// Prompter prompts users to let them input data, and parses it.
type Prompter struct {
	writer io.Writer
	reader *bufio.Reader
}

// New instantiates a new prompter.
func New(w io.Writer, r io.Reader) *Prompter {
	return &Prompter{
		writer: w,
		reader: bufio.NewReader(r),
	}
}
