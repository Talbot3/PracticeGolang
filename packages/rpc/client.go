package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
