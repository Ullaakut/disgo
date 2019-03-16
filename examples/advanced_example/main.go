package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Ullaakut/disgo/prompter"
	"github.com/Ullaakut/disgo/symbol"
	"github.com/ullaakut/disgo/console"
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
	console.SetGlobalOptions(console.WithDebug(true))

	if err := install(); err != nil {
		console.Errorln(console.Failure(err))
		console.Infoln(console.Failure(symbol.Cross), "Installation failed")
		os.Exit(1)
	}

	console.Infoln(console.Success(symbol.Check), "Installation successful")
}

func install() error {
	console.Infoln("Looking for remote database on", console.Link("172.187.10.23"))

	console.StartStep("Accessing database")

	console.StartStep("Checking database integrity")

	console.StartStep("Synchronizing local store")
	console.Infoln("Local store up to date with remote database")
	console.EndStep()

	console.Errorln(console.Failure("Database connection lost"))

	console.StartStep("Connecting to fallback database")
	console.EndStep()

	console.Debugln("Dashboard deployed at", console.Link("https://172.187.10.23:37356/dashboard"))

	result, err := prompter.Confirm(prompter.Confirmation{
		Label:              "Install with current database?",
		EnableDefaultValue: true,
		DefaultValue:       true,
		Choices:            []string{"Y", "n"},
	})
	if err != nil {
		return fmt.Errorf("Unexpected user input: %s", console.Failure(err))
	}

	console.StartStep("Installation in progress")

	if result {
		console.Infoln("Installing with current database")
	} else {
		return console.FailStep(errors.New("unable to install without a database"))
	}

	console.Infoln("Connection to 172.187.10.23 secure")
	console.Infoln("Found 36 dependency requirements")
	console.Infoln("Dependencies resolved")
	console.EndStep()

	return nil
}
