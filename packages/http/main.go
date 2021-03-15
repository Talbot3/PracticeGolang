package main

import (
	"encoding/json"
	"fmt"
	"http/github"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	fmt.Println("Please visit http:127.0.0.1:12345")
	//http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//	s := fmt.Sprintf("hello world ! -- Time: %s", time.Now().String())
	//	fmt.Fprintf(w, "%v\n", s)
	//	log.Printf("%v\n", s)
	//})

	http.HandleFunc("/count", func(writer http.ResponseWriter, request *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		count++
		fmt.Fprintf(writer, "URL.Path = %q %d\n", request.URL.Path, count)
	})

	http.HandleFunc("/issues", func(writer http.ResponseWriter, request *http.Request) {
		result, _ := github.SearchIssues([]string{"a", "b"})
		var data, _ = json.MarshalIndent(result, "", "    ")
		var dataStr = string(data)
		fmt.Fprintln(writer, dataStr)
	})

	http.HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	})

	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
