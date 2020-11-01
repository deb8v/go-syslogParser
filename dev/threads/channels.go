package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go worker(i)
	}
	for { //вечный цикл
	}

}

func worker(id int) {
	time.Sleep(3 * time.Second) //делает работу
	fmt.Println(id, " закончил работу")
}
