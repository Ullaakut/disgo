package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Ullaakut/disgo"
	"github.com/Ullaakut/disgo/style"
)

// Produces the following output on success:
//
// Looking for remote database on 172.187.10.23
// Accessing database...ok
// Checking database integrity...ok
// Synchronizing local store...ok
//   > Local store up to date with remote database
// Database connection lost
// Connecting to fallback database...ok
// Dashboard deployed at https://172.187.10.23:37356/dashboard
// Install with current database? [y/n] y
// Installation in progress...ok
//   > Installing with current database
//   > Connection to 172.187.10.23 secure
//   > Found 36 dependency requirements
//   > Dependencies resolved
// ✔ Installation successful

func main() {
	term := disgo.NewTerminal(disgo.WithDebug(true))

	if err := install(term); err != nil {
		term.Errorln(style.Failure(err))
		term.Infoln(style.Failure(style.SymbolCross), "Installation failed")
		os.Exit(1)
	}

	term.Infoln(style.Success(style.SymbolCheck), "Installation successful")
}

func install(term *disgo.Terminal) error {
	term.Infoln("Looking for remote database on", style.Link("172.187.10.23"))

	term.StartStep("Accessing database")

	term.StartStep("Checking database integrity")

	term.StartStep("Synchronizing local store")
	term.Infoln("Local store up to date with remote database")
	term.EndStep()

	term.Debugln("Dashboard deployed at", style.Link("https://172.187.10.23:37356/dashboard"))

	result, err := term.Confirm(disgo.Confirmation{
		Label:              "Install with current database?",
		EnableDefaultValue: true,
		DefaultValue:       true,
		Choices:            []string{"Y", "n"},
	})
	if err != nil {
		return fmt.Errorf("unexpected user input: %s", style.Failure(err))
	}

	term.StartStep("Installation in progress")

	if result {
		term.Infoln("Installing with current database")
	} else {
		return term.FailStep(errors.New("unable to install without a database"))
	}

	term.Infoln("Connection to 172.187.10.23 secure")
	term.Infoln("Found 36 dependency requirements")
	term.Infoln("Dependencies resolved")
	term.EndStep()

	return nil
}
