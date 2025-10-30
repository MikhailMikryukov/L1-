package main

import (
	"fmt"
	"slices"
)

func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	return slices.Delete(slice, i, i+1)
}

func remove3(slice []int, i int) []int {
	slice = append(slice[0:i], slice[i+1:]...)
	return slice
}

// L1.23 Удаление элемента слайса
func main() {
	a := []int{0, 1, 2, 3, 4, 5}

	a = remove1(a, 2)
	fmt.Println(a)
	a = remove2(a, 2)
	fmt.Println(a)
	a = remove3(a, 2)
	fmt.Println(a)
}
