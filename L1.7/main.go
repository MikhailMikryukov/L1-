package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, mutex *sync.Mutex, m map[int]int) {
	mutex.Lock()
	for i := 0; i < 3; i++ {
		for j := 0; j < 500; j++ {
			m[i] += 1
			m[i] += 1
			m[i] += 1
		}
	}
	mutex.Unlock()
	wg.Done()
}

// L1.7 Конкурентная запись в map
func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&wg, &mutex, m)
	}
	wg.Wait()
	fmt.Println(m)
}
