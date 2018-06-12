package main 

import "fmt"

type User struct{
    name string
}

func (user *User) say(world string) {
    fmt.Printf("%s is say: %s\n" , user.name , world)
}

func main() {
    var user User

    user.name = "jin"

    user.say("nnnnnnnnnn")

    user2 := new(User)

    user2.name = "abli"

    user2.say("66666666666")

    var user3 *User = new(User)

    user3.name = "ehche"

    user3.say("ffffffffffffff")

    user4 := User{name: "wety"}

    user4.say("fghfhhhhhfgggggg")
}