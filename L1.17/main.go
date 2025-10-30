package main

import (
	"fmt"
)

func binarySearch(arr []int, elem int) int {

	halfIndex := len(arr) / 2

	if len(arr) == 0 {
		return -1
	}

	if arr[halfIndex] == elem {
		return halfIndex
	}

	// Ищем в меньшей части массива
	if elem < arr[halfIndex] {
		return binarySearch(arr[:halfIndex], elem)
	}

	// Ищем в большей части массива
	if elem > arr[halfIndex] {
		// в переменную index записываем следующий номер после элемента посередине массива
		// и добавляем номер найденного элемента переданного массива
		index := binarySearch(arr[halfIndex+1:], elem) + halfIndex + 1
		// если после поиска индекс уменьшился значит было передано -1 и элемент не найден
		if index < halfIndex+1 {
			return -1
		} else {
			return index
		}
	}

	return -1

}

// L1.17 Бинарный поиск
func main() {
	arr := []int{0, 1, 38, 67, 90, 432, 520, 544, 645, 709, 1029}

	fmt.Println(binarySearch(arr, 90))
}
