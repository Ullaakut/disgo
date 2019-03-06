package main

import (
	"os"

	"github.com/Ullaakut/colog/logger"
	"github.com/Ullaakut/colog/prompter"
)

func main() {
	log, err := logger.New(os.Stdout, logger.WithDebug(), logger.WithErrorOutput(os.Stderr))
	if err != nil {
		os.Exit(1)
	}

	log.Infoln("Starting listener on localhost:4242 in PID 8484")

	log.Debug(logger.Trace("Accessing database... "))
	log.Debugln(logger.Success("ok"))

	log.Debug(logger.Trace("Checking database integrity... "))
	log.Debugln(logger.Success("ok"))

	log.Debug(logger.Trace("Synchronizing local store... "))
	log.Debugln(logger.Success("ok"))

	log.Infoln(logger.Important("Database is healthy"))

	log.Infoln(logger.Failure("Database connection lost"))
	log.Info(logger.Trace("Connecting to fallback database..."))
	log.Infoln(logger.Success("ok"))

	log.Debugln("Dashboard deployed at", logger.Link("https://172.187.10.23:37356/dashboard"))

	prompt := prompter.New(os.Stdout, os.Stdin)
	result, err := prompt.Confirm(prompter.Confirmation{
		Label: "Are you blue?",

		EnableDefaultValue: true,
		DefaultValue:       true,
	})
	if err != nil {
		log.Errorf("Unexpected level of blueness: %s\n", logger.Failure(err))
		os.Exit(1)
	}

	log.Infoln("Blue:", result)
	log.Infoln(logger.Success("\xE2\x9C\x94 Application ready"))
}
