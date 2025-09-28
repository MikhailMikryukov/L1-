package main

import (
	"fmt"
	"strconv"
)

func setBit(n int64, i int) {
	mask := 1 << i
	newNum := n ^ int64(mask)
	fmt.Println("Число до: ", strconv.FormatInt(n, 2))
	fmt.Println("нужный бит: ", strconv.FormatInt(int64(mask), 2))
	fmt.Println("Число после: ", strconv.FormatInt(newNum, 2))
}

// L1.8 Установка бита в числе
func main() {
	setBit(5, 1)
}
