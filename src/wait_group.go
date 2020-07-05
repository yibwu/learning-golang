package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(i)
	}
	wg.Wait()
}

func foo(i int) {
	fmt.Println(i)
	wg.Done()
}
