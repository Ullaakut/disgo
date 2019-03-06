package main

import (
	"os"

	"github.com/Ullaakut/disgo/logger"
)

func main() {
	log, err := logger.New(os.Stdout, logger.WithDebug(), logger.WithErrorOutput(os.Stderr))
	if err != nil {
		os.Exit(1)
	}

	log.Infoln("Looking for remote database on", logger.Link("172.187.10.23"))

	log.Debug(logger.Trace("Accessing database... "))
	log.Debugln(logger.Success("ok"))

	log.Debug(logger.Trace("Checking database integrity... "))
	log.Debugln(logger.Success("ok"))

	log.Debug(logger.Trace("Synchronizing local store... "))
	log.Debugln(logger.Success("ok"))

	log.Infoln(logger.Important("Local store up to date with remote database"))

	log.Infoln(logger.Failure("Database connection lost"))

	log.Info(logger.Trace("Connecting to fallback database..."))
	log.Infoln(logger.Success("ok"))

	log.Debugln("Dashboard deployed at", logger.Link("https://172.187.10.23:37356/dashboard"))
	log.Infoln(logger.Success("\xE2\x9C\x94 Installation successful"))
}
