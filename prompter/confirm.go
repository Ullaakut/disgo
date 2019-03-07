package prompter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ullaakut/disgo/logger"
)

var (
	// DefaultConfirmationChoices is the default value for the choices that are given to
	// the users in a confirmation prompt.
	DefaultConfirmationChoices = []string{"y", "n"}
)

// ConfirmationParser is a function that parses an input and returns a
// confirmation value as well as an error, if the input can't be parsed.
type ConfirmationParser func(string) (bool, error)

// DefaultConfirmation is a confirmation parser that covers most cases
// for confirmation.
// It converts y/Y/yes/YES/t/T/true/True/1 to true.
// It converts n/N/no/NO/f/F/false/FALSE/0 to false.
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

// Confirmation represents a confirmation prompt's configuration.
type Confirmation struct {
	// The label that will be prompted to the user.
	// Example: `Are you sure?`
	Label string

	// The choices that will be presented to the user.
	// Example: `Y/n`. (A good practice is to uppercase
	// the default value, if there is one).
	Choices []string

	// EnableDefaultValue tells the prompter whether or not
	// there is a default value that will be used when the
	// user doesn't input any data.
	EnableDefaultValue bool

	// DefaultValue is the default value that will be used when
	// the user doesn't input any data, if EnableDefaultValue
	// is set to true OR that the prompter is set to not
	// interactive.
	DefaultValue bool

	// The parser that will be used to convert the user's input
	// into a true/false value.
	Parser ConfirmationParser

	// RequireValidInput makes the prompt loop until the user
	// enters an input that doesn't trigger a parser error.
	RequireValidInput bool
}

func (c Confirmation) parser() ConfirmationParser {
	if c.Parser != nil {
		return c.Parser
	}
	return DefaultConfirmation
}

func (c Confirmation) choices() string {
	if c.Choices != nil {
		return strings.Join(c.Choices, "/")
	}
	return strings.Join(DefaultConfirmationChoices, "/")
}

// Confirm prompts the user to confirm something.
func (p Prompter) Confirm(config Confirmation) (bool, error) {
	// If prompter is not set to interactive, directly return the default value.
	if !p.interactive {
		return config.DefaultValue, nil
	}

	// Print the label and choices.
	fmt.Fprintf(p.writer, "%s [%s] ", config.Label, config.choices())

	// Wait for user input.
	text, err := p.reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	// If user just pressed enter directly, return default value.
	if config.EnableDefaultValue && text == "\n" {
		return config.DefaultValue, nil
	}

	// Parse user input.
	res, err := config.parser()(strings.TrimSpace(text))
	if err != nil && config.RequireValidInput {
		// Show parse error and call confirm again until a valid input is given
		fmt.Fprintf(p.writer, "%s\n", logger.Failure(err))
		return p.Confirm(config)
	}

	return res, err
}
