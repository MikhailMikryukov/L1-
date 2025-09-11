package main

import (
	"fmt"
)

func square(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= arr[i]
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	go square(arr)
	go square(arr)

	fmt.Println(arr)
}
