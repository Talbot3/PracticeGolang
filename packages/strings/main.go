package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 前缀判断
	var has = strings.HasPrefix("test", "t")
	// 后缀判断
	var suffix = strings.HasSuffix("FromIsGood", "Good")
	fmt.Println(has, suffix)

	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d,%02d,%4d\n", t.Day(), t.Month(), t.Year())

}
