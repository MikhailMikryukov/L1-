package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// L1.9 Конвейер чисел
func main() {
	var arr [10]int
	fmt.Print("Сгенерированный массив: ")
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	var wg sync.WaitGroup

	firstCh := make(chan int)
	secondCh := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(arr); i++ {
			firstCh <- arr[i]
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(secondCh)

		for i := 0; i < len(arr); i++ {
			v := <-firstCh
			secondCh <- v * 2
		}
	}()

	fmt.Print("Значения массива, умноженные на 2: ")
	for v := range secondCh {
		fmt.Print(v, " ")
	}

	wg.Wait()
}
