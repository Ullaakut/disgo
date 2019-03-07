package main

import (
	"fmt"
	"os"

	"github.com/Ullaakut/disgo/prompter"
)

func main() {
	prompt := prompter.New(os.Stdout, os.Stdin, true)
	confirmOptions := prompter.Confirmation{
		Label:   "Install with current database?",
		Choices: []string{"that's right", "maybe not", "definitely not"},
		Parser: func(input string) (bool, error) {
			switch input {
			case "that's right":
				return true, nil
			case "maybe not", "definitely not":
				return false, nil
			default:
				return false, fmt.Errorf("%q is not a valid choice", input)
			}
		},
		EnableDefaultValue: true,
		DefaultValue:       true,
	}

	installWithCurrentDB, err := prompt.Confirm(confirmOptions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected user input: %v\n", err)
		os.Exit(1)
	}

	if installWithCurrentDB {
		// ...
	}
	// ...
}
