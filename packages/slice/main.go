package main

import "fmt"

func main() {
	var (
		a []int
		b = []int{}
		c = []int{1, 2, 3}
		d = c[:2]
		e = c[0:2:cap(c)]
		f = c[:0]
		g = make([]int, 3)
		h = make([]int, 2, 3)
		i = make([]int, 0, 3)
	)
	println(a, b, c, d, e, f, g, h, i)

	for i, v := range c {
		fmt.Println(i, v)
	}

	a = append(a, 1)
	a = append(a, 1, 2, 3)
	a = append(a, []int{1, 2, 3}...)
	a = append([]int{-3, -2, -1, 0}, a...)
	for i, v := range a {
		fmt.Printf(" %d:%d ", i, v)
	}

}
