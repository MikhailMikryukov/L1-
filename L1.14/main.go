package main

import "fmt"

func valueType(v interface{}) {

	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("chan")
	default:
		fmt.Printf("%T\n", v)
	}
}

// L1.14 Определение типа переменной в runtime
func main() {
	valueType(5)
	valueType(true)
	valueType("true")
	valueType(make(chan string))
}
