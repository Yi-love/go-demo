package main

import "fmt"

func main() {
    a , b , c := 1 , 2 , 3

    fmt.Println(a , b , c)

    ap := &a

    *ap = 6666
    fmt.Println(ap , *ap , a)
}