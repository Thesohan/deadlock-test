package main

import (
	"fmt"
	"os"
	"time"
)

func tickerCleanUp(ticker *time.Ticker, heartbeatDone chan bool) {
	ticker.Stop()
	heartbeatDone <- true
	close(heartbeatDone)
}

func main() {
	fmt.Println("Hello World")
	heartbeatFreq := time.Duration(10) * time.Second
	ticker := time.NewTicker(heartbeatFreq)
	heartbeatDone := make(chan bool)

	go func() {
		for {
			select {
			case <-heartbeatDone:
				return
			case <-ticker.C:
				fmt.Println("hi")
				tickerCleanUp(ticker, heartbeatDone)
				fmt.Println("before exit")
				os.Exit(1) // Causes systemctl to restart

			}
		}
	}()
	time.Sleep(12 * time.Second)
	tickerCleanUp(ticker, heartbeatDone)
	fmt.Println("exit")
}
