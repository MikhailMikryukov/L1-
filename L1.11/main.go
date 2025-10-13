package main

import "fmt"

// L1.11 Пересечение множеств
func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	var result []int

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				result = append(result, A[i])
			}
		}
	}

	fmt.Println(result)
}
