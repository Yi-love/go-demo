package main 

import (
    "net/http"
    "log"
)

func helloHandler( w http.ResponseWriter , r *http.Request ){
    w.Write([]byte("hello jin!"))
}

func main() {
    http.HandleFunc("/" , helloHandler)
    err := http.ListenAndServe(":8888" , nil)

    if err != nil {
        log.Fatal("ListenAndServe: " , err)
    }
}