package main

import (
	"fmt"
)

// L1.13 Обмен значениями без третьей переменной
func main() {
	a := 5
	b := 10

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}
