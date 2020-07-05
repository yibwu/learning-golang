package main

import (
	"fmt"
	"sync"
)

const (
	buffSize   = 100
	concurrent = 4
)

func producer(nums []int) <-chan int {
	out := make(chan int, buffSize)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int, buffSize)

	go func() {
		defer close(out)
		var wg sync.WaitGroup
		for i := 0; i < concurrent; i += 1 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for n := range inCh {
					out <- lambdaFunc(n)
				}
			}()
		}
		wg.Wait()
	}()

	return out
}

func lambdaFunc(n int) int {
	return n * n
}

func main() {
	in := producer([]int{1, 2, 3, 4})
	out := square(in)

	for n := range out {
		fmt.Printf("%d ", n)
	}
	fmt.Println("Done")
}
