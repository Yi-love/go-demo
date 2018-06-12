package main 

import "fmt"

func retrunTrue() bool {
    fmt.Println("is run true  .....")
    return true
}

func returnFalse() bool {
    fmt.Println("is run false .....")
    return false
}

func main() {
    if ( returnFalse() && retrunTrue() ){
        fmt.Println("all true")
    }else{
        fmt.Println("retrun false is stop")
    }
}