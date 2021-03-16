package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func init() {
	println("init now")
}

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

func getX2AndX3(input int) (x2 int, x3 int) {
	x2 = 3 * input
	x3 = 4 * input
	return
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func g(i int) {
	if i > 3 {
		fmt.Printf("Packing!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g ", i)
	g(i + 1)
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recoverec in f")
		}
	}()
	fmt.Printf("Calling g.")
	g(0)
	fmt.Printf("Returned normally form g.")
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
	num1, num2 := getX2AndX3(33)
	Multiply(num1, num2, &result)
	fmt.Println("num1 ->", num1, num2, result)
	minV := min(1, 2, 3, 4)
	fmt.Println("min value -> ", minV)
	slice := []int{9, 99, 8, -987}
	x := min(slice...)
	fmt.Println("min -> ", x)
	f()
}
