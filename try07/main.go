package main

import "fmt"

type Tuple[X_T any, Y_T any] struct {
	X X_T
	Y Y_T
}

func New[X_T any, Y_T any](
	x X_T,
	y Y_T,
) *Tuple[X_T, Y_T] {
	return &Tuple[X_T, Y_T]{
		X: x,
		Y: y,
	}
}

func main() {
	t := New(10, "hoge")
	fmt.Println(t.X, t.Y)
}
