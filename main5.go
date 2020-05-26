package main

import (
	"fmt"
	"reflect"
)
func main() {
	copyForSlice()
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
	mNewMap := new(map[int]string)
	mMakeMap := make(map[int]string)
	t := reflect.TypeOf(mNewMap)
	t1:= reflect.TypeOf(mMakeMap)
	fmt.Println("new map :", t)
	fmt.Println("new make:", t1)
}

func appendElementForSlice() {
	mIDSlice := make([]string, 3)
	mIDSlice[0] = "id-1"
	mIDSlice = append(mIDSlice, "id-2")
	fmt.Println(mIDSlice)
	fmt.Println(len(mIDSlice))
	fmt.Println(cap(mIDSlice))
}

func copyForSlice() {
	mIDSliceDst := make([]string, 3)
	mIDSliceDst[0] = "id-1-dst"
	mIDSliceDst = append(mIDSliceDst, "id-2-dst")


	mIDSliceSrc := make([]string, 3)
	mIDSliceSrc[0] = "id-1-src"
	mIDSliceSrc = append(mIDSliceSrc, "id-2-src")

	copy(mIDSliceSrc, mIDSliceDst);
	fmt.Println("src",mIDSliceSrc, "dst", mIDSliceDst);
}