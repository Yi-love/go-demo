package main

import "fmt"

func xxx() func() int {
    i := 1
    return func() int {
        i ++
        return i * i
    }
}

func main() {
    numXXX := xxx()

    for i := 0; i < 5 ; i++ {
        fmt.Println("num xxx === " , numXXX())
    }
}