package main

import (
	"fmt"
	"sync"
)

const (
	buffSize   = 100
	concurrent = 4
)

type SliceItem struct {
	Val int
	Idx int // 保存原始数据的位置，当coroutine乱序消费后恢复顺序时用
}

func producer(nums []int) <-chan SliceItem {
	out := make(chan SliceItem, buffSize)

	go func() {
		defer close(out)
		for i, n := range nums {
			out <- SliceItem{Idx: i, Val: n}
		}
	}()

	return out
}

func square(inCh <-chan SliceItem) <-chan SliceItem {
	out := make(chan SliceItem, buffSize)

	go func() {
		defer close(out)
		var wg sync.WaitGroup
		for i := 0; i < concurrent; i += 1 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for item := range inCh {
					out <- SliceItem{Idx: item.Idx, Val: lambdaFunc(item.Val)}
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
	input := []int{1, 2, 3, 4}
	inCh := producer(input)
	outCh := square(inCh)

	output := make([]int, len(input))
	for item := range outCh {
		output[item.Idx] = item.Val
	}
	fmt.Println(output)
}
