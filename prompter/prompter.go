// Package prompter is a simple user prompter that asks users
// for input data or confirmations.
package prompter

import (
	"bufio"
	"io"
)

// Prompter prompts users to let them input data, and parses it.
type Prompter struct {
	// Writer on which the prompt's label and choices are
	// written when prompting them.
	writer io.Writer

	// Reader from which the user's response to the prompt is
	// read.
	reader *bufio.Reader

	// Whether or not this prompter should be interactive. If this is
	// set to false, the prompter will never prompt the user and always
	// return default values. This can be useful for running this code
	// outside of a TTY for example.
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
