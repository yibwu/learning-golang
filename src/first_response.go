package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func worker() string {
	n := rand.Intn(10) + 1
	fmt.Println("generate random number: ", n)
	time.Sleep(time.Duration(n) * time.Second)
	return strconv.Itoa(n)
}

func firstResponse(numOfRunner int) string {
	ch := make(chan string)

	for i := 0; i < numOfRunner; i += 1 {
		go func(ch chan string) {
			res := worker()
			ch <- res
		}(ch)
	}

	return <- ch
}

func main() {
	start := time.Now()
	fmt.Println(firstResponse(10), time.Since(start))
}
