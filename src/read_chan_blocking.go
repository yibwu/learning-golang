package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Done")
}
