package main 

import "fmt"

func changeVal( arr *[5]int ) {
    for i := 0 ; i < len(arr) ; i ++ {
        arr[i] =  i + 2
    }
}

func changeSliceArr( arr []int ) {
    for i := 0 ; i < len(arr) ; i ++ {
        arr[i] =  i - 2
    }
}

func main() {
    var arr [5]int
    sliceArr := make([]int , 5)

    changeVal( &arr )
    changeSliceArr( sliceArr )

    for i := 0 ; i < len(arr) ; i ++ {
        fmt.Printf("arr val === %d\n" , arr[i])
    }

    for i := 0 ; i < len(sliceArr) ; i ++ {
        fmt.Printf("sliceArr val === %d\n" , sliceArr[i])
    } 

    for _ , v := range sliceArr {
        fmt.Printf("range val === %d\n" , v)
    }
}