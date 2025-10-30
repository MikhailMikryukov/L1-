package main

import (
	"fmt"
	"math"
	"math/big"
)

func subtraction(a int, b int) {
	fmt.Println(a - b)
}

func sum(a int, b int) {
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	result := new(big.Int)
	fmt.Println(result.Add(bigA, bigB))
}

func multiplication(a int, b int) {
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	result := new(big.Int)
	fmt.Println(result.Mul(bigA, bigB))
}

func division(a int, b int) {
	fmt.Println(a / b)
}

//  L1.22 Большие числа и операции
func main() {
	a := math.MaxInt
	b := math.MaxInt
	subtraction(a, b)
	sum(a, b)
	multiplication(a, b)
	division(a, b)

}
