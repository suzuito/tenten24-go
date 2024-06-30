package main

import "fmt"

func Filter[T comparable](
	arr []T,
	f func(i int, v T) bool,
) []T {
	ret := []T{}
	for i, e := range arr {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func main() {
	ns := []int{1, 2, 3, 4}
	ms := Filter[int](ns, func(i, v int) bool {
		return v%2 == 0
	})
	fmt.Println(ms)
}
