package prompter

import (
	"bufio"
	"io"
)

// Prompter prompts users to let them input data, and parses it.
type Prompter struct {
	writer io.Writer
	reader *bufio.Reader

	interactive bool
}

// New instantiates a new prompter which will prompt users on the writer
// and read their output from the reader. The interactive boolean makes
// all prompts return a default value if set to false, and won't prompt
// them.
// This should be used if your users are not in a TTY and can't
// write to answer to the prompt.
func New(w io.Writer, r io.Reader, interactive bool) *Prompter {
	return &Prompter{
		writer:      w,
		reader:      bufio.NewReader(r),
		interactive: interactive,
	}
}
