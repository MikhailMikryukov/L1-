package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func (h Human) Speak() {
	fmt.Println("Speaking")
}
func main() {

	action := Action{}
	action.Speak()
}
