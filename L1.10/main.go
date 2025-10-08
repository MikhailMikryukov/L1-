package main

import "fmt"

// L1.10 Группировка температур
func main() {
	arr := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for i := 0; i < len(arr); i++ {
		j := int(arr[i] / 10)
		m[j*10] = append(m[j*10], arr[i])
	}
	fmt.Println(m)
}
