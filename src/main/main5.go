package main

import "fmt"

func main() {
	makeMap()
}

// makeSlice 创建切片
func makeSlice() {
	mSlice := make([]string, 3);
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"

	fmt.Println(mSlice)
}

// makeMap 创建map
func makeMap() {
	mMap := make(map[int]string)
	mMap[10] = "dog"
	mMap[100] = "cat"
	fmt.Println(mMap)
}