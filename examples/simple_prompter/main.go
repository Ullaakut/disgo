package main

import (
	"fmt"
	"os"

	"github.com/Ullaakut/disgo/prompter"
)

func main() {
	prompt := prompter.New(os.Stdout, os.Stdin, true)
	installWithCurrentDB, err := prompt.Confirm(prompter.Confirmation{
		Label: "Install with current database?",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected user input: %v\n", err)
		os.Exit(1)
	}

	if installWithCurrentDB {
		// ...
	}
	// ...
}
