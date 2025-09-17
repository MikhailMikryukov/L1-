package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, readChan <-chan int, sendChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(time.Second)
		value := <-readChan
		sendChan <- fmt.Sprintf("Worker: %d, Value: %d", id, value)
	}

}

// L1.3 Работа нескольких воркеров
func main() {
	workersCount := 3

	var wg sync.WaitGroup

	readChan := make(chan int)
	sendChan := make(chan string)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(i+1, readChan, sendChan, &wg)
	}

	go func() {
		for {
			readChan <- rand.Intn(100)
			time.Sleep(time.Second)
		}
	}()

	for v := range sendChan {
		fmt.Println(v)
	}
	wg.Wait()
}
