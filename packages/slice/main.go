package main

import (
	"fmt"
	"unicode/utf8"
)

func str2runes(s []byte) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

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

	var b1 = [...]int{1: 1, 3: 2, 4: 3}
	for i1, v1 := range b1 {
		fmt.Println(i1, v1)
	}

	// 函数数组
	// var decoder1 = [2]func(io.Reader){image.Image, error}

	fmt.Printf("%#v\n", []rune("世界"))
	z := "hello, world"
	hello := z[:5]
	world := z[7:]
	fmt.Printf("%s\n%s\n", hello, world)
	r, size := utf8.DecodeRune([]byte(z))
	fmt.Println(r, size)
	fmt.Printf("%#v \n", []byte(z))

	strArr := str2runes([]byte(z))
	fmt.Println(strArr)
}
