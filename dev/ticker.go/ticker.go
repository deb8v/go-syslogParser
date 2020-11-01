package main

import (
	"fmt"
	"time"
)

func listen(messages chan string) {
	for { //в вечном цикле ждём что появилось нового
		msg := <-messages
		fmt.Println(msg)
	}
}
func main() {
	messages := make(chan string)
	go listen(messages) //запускаем отдельный поток который слушает messages

	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C: //получаем из ticker дату
				fmt.Println("Queue starting", t)
				messages <- "heelo im BIOS" //шлем в messages сообщение
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
