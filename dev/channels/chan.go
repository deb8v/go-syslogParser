package main

import "fmt"

func listen(messages chan string) {
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()
	go listen(messages)
}
