package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type incrementStruct struct {
	count int64
}

func (receiver *incrementStruct) increment() {
	atomic.AddInt64(&receiver.count, 1)
}

// L1.18 Конкурентный счетчик
func main() {
	var incrementStruct incrementStruct
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			incrementStruct.increment()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			incrementStruct.increment()
		}
	}()

	wg.Wait()
	fmt.Println(incrementStruct.count)
}
