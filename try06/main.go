package main

import "fmt"

func Map[X any, Y any](s []X, f func(X) Y) []Y {
	ret := make([]Y, len(s))
	for i, v := range s {
		ret[i] = f(v)
	}
	return ret
}

func main() {
	var ss []string = Map([]int{10, 20}, func(n int) string {
		return fmt.Sprintf("0x%x", n)
	})
	fmt.Println(ss)
}
