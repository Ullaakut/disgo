package colog

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Prompter ...
type Prompter struct {
	writer io.Writer
	reader *bufio.Reader
}

// ConfirmationParser ...
type ConfirmationParser func(string) (bool, error)

// DefaultConfirmation ...
func DefaultConfirmation(input string) (bool, error) {
	switch input {
	case "y", "Y", "yes", "YES":
		return true, nil
	case "n", "N", "no", "NO":
		return false, nil
	}

	// ParseBool handles cases such as 0/1, true/false, t/f etc.
	return strconv.ParseBool(input)
}

// NewPrompter ...
func NewPrompter(w io.Writer, r io.Reader) *Prompter {
	return &Prompter{
		writer: w,
		reader: bufio.NewReader(r),
	}
}

// Confirm ...
func (p Prompter) Confirm(label string, defaultValue bool, parser ConfirmationParser) (bool, error) {
	fmt.Fprint(p.writer, label)

	text, err := p.reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	if text == "\n" {
		return defaultValue, nil
	}

	return parser(strings.TrimSpace(text))
}
