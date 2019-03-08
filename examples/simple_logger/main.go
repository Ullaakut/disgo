package main

import (
	"os"

	"github.com/Ullaakut/disgo/logger"
	"github.com/Ullaakut/disgo/symbol"
)

func main() {
	log, err := logger.New(os.Stdout, logger.WithDebug(true), logger.WithErrorOutput(os.Stderr))
	if err != nil {
		os.Exit(1)
	}

	log.Infoln("Looking for remote database on", logger.Link("172.187.10.23"))

	log.Debug(logger.Trace("Accessing database... "))
	log.Debugln(logger.Success(symbol.Check))

	log.Debug(logger.Trace("Checking database integrity... "))
	log.Debugln(logger.Success(symbol.Check))

	log.Debug(logger.Trace("Synchronizing local store... "))
	log.Debugln(logger.Success(symbol.Check))

	log.Infoln(logger.Important("Local store up to date with remote database"))

	log.Infoln(logger.Failure("Database connection lost"))

	log.Info(logger.Trace("Connecting to fallback database..."))
	log.Infoln(logger.Success("ok"))

	log.Debugln("Dashboard deployed at", logger.Link("https://172.187.10.23:37356/dashboard"))
	log.Infoln(logger.Success(symbol.Check + " Installation successful"))
}
