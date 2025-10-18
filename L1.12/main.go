package main

import "fmt"

// Получение множества уникальных строк
func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var result []string

	result = append(result, arr[0])

	for i := 0; i < len(arr); i++ {
		contains := false
		for j := 0; j < len(result); j++ {
			if arr[i] == result[j] {
				contains = true
				break
			}
		}
		if !contains {
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)
}
