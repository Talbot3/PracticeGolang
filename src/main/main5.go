package main

import (
	"fmt"
	"reflect"
)
func main() {
	NewMap()
}

// makeSlice 创建切片
func makeSlice() {
	mSlice := make([]string, 3)
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

// makeChan 创建没有缓存的chan
func makeChan() {
	mChan := make(chan int)
	close(mChan)
}

func NewMap() {
	mMap := new(map[int]string)
	fmt.Println(reflect.Typeof(mMap))
}