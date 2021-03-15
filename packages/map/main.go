package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapCreated["key3"] = 3

	fmt.Printf("Map literal at \"one\" is %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is %d\n", mapLit["key2"])
	fmt.Printf("Map assigned at \"two\" is %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is %d\n", mapLit["ten"])

	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)

	for key, value := range mf {
		fmt.Printf("key is %d -  value is : %p", key, value)
	}

	var value1 int
	var isPresent bool

	map2 := make(map[string]int, 100)
	map2["Beijing"] = 20
	map2["Delete"] = -1
	value1, isPresent = map2["Beijing"]

	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map2 is : %d\n", value1)
	}

	delete(map2, "Delete")
	value1, isPresent = map2["Delete"]
	if isPresent {
		fmt.Printf("The value of \"Delete\" is %d\n", value1)
	} else {
		fmt.Printf("The map2 not contain key of \"Delete\" is %d\n", value1)
	}

	// map类型的切片
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}

	fmt.Printf("Version A: Value of items: %v\n", items)

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
