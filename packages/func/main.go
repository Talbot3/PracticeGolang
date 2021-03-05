package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func SortFloat64FastV1(a []float64) {
	// 强制转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	sort.Ints(c)
}

func main() {
	for i := 0; i < 3; i++ {
		i := i
		defer func(i int) { println(i) }(i)
	}

	var Add = func(a, b int) int {
		return a + b
	}
	result := Add(3, 2)
	println(result)

	var a = []interface{}{123, "abc"}

	fmt.Println(a...)
	fmt.Println(a)

	var a1 = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	SortFloat64FastV2(a1)
	fmt.Println(a1)
}
