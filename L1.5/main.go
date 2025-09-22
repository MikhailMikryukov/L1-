package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// L1.5 Таймаут на канал
func main() {
	n := 3

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	channel := make(chan string)

	exit := time.After(time.Duration(n) * time.Second)

	go func() {
		for {
			select {
			case <-exit:
				close(channel)
				cancel()
			default:
				time.Sleep(500 * time.Millisecond)
				channel <- fmt.Sprintf("Value: %d", rand.Intn(100))
			}
		}
	}()
	for v := range channel {
		fmt.Println(v)
	}
}
