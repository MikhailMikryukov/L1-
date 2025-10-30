package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// опорный элемент первый
	pivot := arr[0]
	// индекс куда последний раз записали элемент больший опорного
	high := len(arr) - 1
	// итерация с конца, так как большие элементы закидываем в конец
	for i := len(arr) - 1; i >= 0; i-- {

		if arr[i] > pivot {
			arr[i], arr[high] = arr[high], arr[i]
			high--
		}
	}
	// записываем опорный элемент перед элементами которые больше
	arr[high], arr[0] = arr[0], arr[high]
	quickSort(arr[:high])
	quickSort(arr[high+1:])
	return arr
}

// L1.16 Быстрая сортировка (quicksort)
func main() {
	arr := []int{100, 43, 953, 429, 342, 85, 1, 59, 3291, 4812}
	fmt.Println(quickSort(arr))
}
