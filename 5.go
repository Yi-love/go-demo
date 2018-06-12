package main

import "fmt"

func main() {
    num := 0

    A : for num < 100 {
        num ++
        if num % 5 == 0 {
            goto A
        }
        fmt.Printf("num %d is %% 5 !=== 0\n" , num)
    }
}