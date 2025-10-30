package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	strLower := strings.ToLower(str)
	seen := make(map[rune]bool)

	for _, char := range strLower {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

// L1.26 Уникальные символы в строке
func main() {
	fmt.Println(IsUnique("abAdc"))
}
