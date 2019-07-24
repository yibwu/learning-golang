package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	stop := make(chan chan bool)
	quit := make(chan bool)
	
	go send(ch, stop, quit)
	go receive(ch, stop, quit)
	
	<- quit
	fmt.Println("Done")
}

func send(ch chan int, stop chan chan bool, quit chan bool) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	time.Sleep(3 * time.Second)
	stop <- quit
}

func receive(ch chan int, stop chan chan bool, quit chan bool) {
	for {
		select {
			case i := <- ch:
				fmt.Println(i)
			case q := <- stop:
				q <- true
				break
		}
	}
}
