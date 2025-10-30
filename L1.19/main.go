package main

import "fmt"

// L1.19 Разворот строки
func reverseString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseString("главрыба"))
}
