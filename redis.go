package main 

import (
    "fmt"
    "log"
    "flag"
    "time"
    "github.com/gomodule/redigo/redis"
)

var (
    pool *redis.Pool
    redisServer = flag.String("redisServer" , ":6379" , "")
)

func newPool(addr string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 3,
        IdleTimeout: 240 * time.Second,
        Dial: func ()(redis.Conn , error){ 
            c , err := redis.Dial("tcp" , addr)
            if err != nil {
                return nil , err
            }
            return c , err
        },
    }
}

func main() {
    flag.Parse()
    pool = newPool(*redisServer)

    conn := pool.Get()

    defer conn.Close()

    result , error := conn.Do("SET" , "test" , "888888")
    if error != nil {
        log.Fatal(error)
    }
    fmt.Println("result: " , result)

    data , err := redis.String(conn.Do("GET" , "test"))

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("redis : " , data)
}