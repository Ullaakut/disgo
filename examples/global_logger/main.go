package main

import (
	"sync"

	"github.com/ullaakut/disgo/logger"
)

func main() {
	logger.SetGlobalOptions(logger.WithDebug(true))

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			logger.Infoln(i, "> Goroutine 1 writes an info log")
			logger.Debugln(i, "> Goroutine 1 writes a debug log")
			logger.Errorln(i, "> Goroutine 1 writes an error log")
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			logger.Infoln(i, "> Goroutine 2 writes an info log")
			logger.Debugln(i, "> Goroutine 2 writes a debug log")
			logger.Errorln(i, "> Goroutine 2 writes an error log")
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			logger.Infoln(i, "> Goroutine 3 writes an info log")
			logger.Debugln(i, "> Goroutine 3 writes a debug log")
			logger.Errorln(i, "> Goroutine 3 writes an error log")
		}
	}()

	wg.Wait()
}
