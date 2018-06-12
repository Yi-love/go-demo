package main 

import "fmt"

type User struct{
    age int
    name string
}

func main() {
    mp := make(map[string]int)

    mp ["age"] = 12

    fmt.Println(mp["age"])

    user := User{}

    user.age = 12

    fmt.Println(user.age , user)

    ump := make(map[string]User)

    ump ["user"] = user

    fmt.Println(ump , ump["user"])

}