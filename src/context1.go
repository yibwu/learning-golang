package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    start := time.Now()

    newCtx, cancelFunc := context.WithTimeout(ctx, time.Second * 3)
    go func() {
        time.Sleep(time.Second * 5)
        cancelFunc()
    }()
    <-newCtx.Done()

    fmt.Println(newCtx, time.Since(start), newCtx.Err())
}