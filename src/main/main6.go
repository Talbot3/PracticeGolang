package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    runtime.GOMAXPROCS(1)

    fmt.Println("The program starts ...")

    go func() {
        for {
			fmt.Println("A.B.B..B")
        }
    }()

    time.Sleep(time.Second)
    fmt.Println("I got scheduled!")
}