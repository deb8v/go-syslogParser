package main

import (
	"fmt"
	"time"
)

func listen(messages chan string) {
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
func main() {
	messages := make(chan string)
	go listen(messages)

	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Queue starting", t)
				messages <- "heelo im BIOS"
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
