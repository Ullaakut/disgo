package main

import (
	"sync"

	"github.com/ullaakut/disgo/logger"
)

func main() {
	logger.SetGlobalOptions(logger.WithDebug(true))

	wg := &sync.WaitGroup{}

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(number int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				logger.Infoln(j, "> Goroutine", number, "writes an info log")
				logger.Debugln(j, "> Goroutine", number, "writes a debug log")
				logger.Errorln(j, "> Goroutine", number, "writes an error log")
			}
		}(i)
	}

	wg.Wait()
}
