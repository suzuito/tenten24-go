package main

import "fmt"

func Apply[T comparable](
	arr []T,
	f func(int, T),
) {
	for i, e := range arr {
		f(i, e)
	}
}

func main() {
	var sum int
	Apply[int]([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println(sum) // 30
}
