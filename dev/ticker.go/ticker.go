package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Queue starting", t)
			}
		}
	}()

	//time.Sleep(16000 * time.Millisecond)
	//ticker.Stop()
	//done <- true
	//fmt.Println("Ticker stopped")
	for {
	}
}
