package main

import (
	"os"

	"github.com/Ullaakut/colog"
)

func main() {
	logger, err := colog.NewLogger(os.Stdout, colog.WithDebug())
	if err != nil {
		os.Exit(1)
	}

	logger.Infoln("Starting listener", colog.Trace("on localhost:4242"), colog.Trace("in PID 8484"))

	logger.Debug(colog.Trace("Accessing database... "))
	logger.Debugln(colog.Success("ok"))

	logger.Infoln(colog.Important("Database is healthy"))

	logger.Infoln(colog.Failure("Database connection lost"))
	logger.Info(colog.Trace("Connecting to fallback database..."))
	logger.Infoln(colog.Success("ok"))

	logger.Debugln("Dashboard deployed at ", colog.Link("https://172.187.10.23:37356/dashboard"))

	prompt := colog.NewPrompter(os.Stdout, os.Stdin)
	result, err := prompt.Confirm(colog.Important("Are you stupid? [y/N] "), true, colog.DefaultConfirmation)
	logger.Infof("Result: %v Err: %v\n", result, err)

	logger.Infoln(colog.Success("\xE2\x9C\x94 Application ready"))
}
