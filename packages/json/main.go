package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	b := []byte(`{
		"Title": "Go Programing",
		"Proce": 9.99
	}`)

	var r interface{}
	json.Unmarshal(b, &r)

	gobook, ok := r.(map[string]interface{})

	if ok {
		for k, v := range gobook {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Print(i, iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet", v2)
			}
		}
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
