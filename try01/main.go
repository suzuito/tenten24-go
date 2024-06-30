package main

import "fmt"

func Ptr[T any](v T) *T {
	return &v
}

func main() {
	inputValue := false
	PtrBool := Ptr[bool]
	ptrBool := PtrBool(inputValue)
	fmt.Println(ptrBool)
	fmt.Printf("%p\n", &inputValue)
}
