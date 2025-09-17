package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, readChan <-chan int, sendChan chan<- string, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d was closed\n", id)
			return
		case value := <-readChan:
			time.Sleep(time.Second)
			sendChan <- fmt.Sprintf("Worker: %d, Value: %d", id, value)
		}

	}

}

// L1.4 Завершение по Ctrl+C
func main() {
	workersCount := 3

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	readChan := make(chan int)
	sendChan := make(chan string)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(i+1, readChan, sendChan, &wg, ctx)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				cancel()
				close(sendChan)
				return
			default:
				readChan <- rand.Intn(100)
				time.Sleep(time.Second)
			}

		}
	}()

	for v := range sendChan {
		fmt.Println(v)
	}
	wg.Wait()
}
