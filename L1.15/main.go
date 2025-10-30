package main

import (
	"fmt"
	"strings"
)

func someFunc() string {
	v := createHugeString(1 << 10)
	buf := make([]byte, 100)
	copy(buf, v[:100])
	return string(buf)
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

// L1.15
func main() {
	fmt.Println(someFunc())
}
