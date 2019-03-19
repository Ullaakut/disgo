package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ullaakut/disgo"
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
	disgo.SetupGlobalConsole(disgo.WithDebug(true))

	if err := install(); err != nil {
		disgo.Errorln(disgo.Failure(err))
		disgo.Infoln(disgo.Failure(disgo.SymbolCross), "Installation failed")
		os.Exit(1)
	}

	disgo.Infoln(disgo.Success(disgo.SymbolCheck), "Installation successful")
}

func install() error {
	disgo.Infoln("Looking for remote database on", disgo.Link("172.187.10.23"))

	disgo.StartStep("Accessing database")

	disgo.StartStep("Checking database integrity")

	disgo.StartStep("Synchronizing local store")
	disgo.Infoln("Local store up to date with remote database")
	disgo.EndStep()

	disgo.Debugln("Dashboard deployed at", disgo.Link("https://172.187.10.23:37356/dashboard"))

	result, err := disgo.Confirm(disgo.Confirmation{
		Label:              "Install with current database?",
		EnableDefaultValue: true,
		DefaultValue:       true,
		Choices:            []string{"Y", "n"},
	})
	if err != nil {
		return fmt.Errorf("Unexpected user input: %s", disgo.Failure(err))
	}

	disgo.StartStep("Installation in progress")

	if result {
		disgo.Infoln("Installing with current database")
	} else {
		return disgo.FailStep(errors.New("unable to install without a database"))
	}

	disgo.Infoln("Connection to 172.187.10.23 secure")
	disgo.Infoln("Found 36 dependency requirements")
	disgo.Infoln("Dependencies resolved")
	disgo.EndStep()

	return nil
}
