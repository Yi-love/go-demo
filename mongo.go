package main 

import (
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Name string
    Age  int
}

func main() {
    fmt.Println("connect start....")

    session , err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }

    defer session.Close()
    
    fmt.Println("connect success")
    
    session.SetMode(mgo.Monotonic , true)

    c := session.DB("test").C("user")

    err = c.Insert(&User{"jin" , 18} , 
                    &User{"abli" , 19})

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("insert success")

    result := User{}

    err = c.Find(bson.M{"name": "jin"}).One(&result)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("user: " , result)
}