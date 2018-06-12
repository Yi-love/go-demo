package main 

import (
    "fmt"
    "os"
    "io/ioutil"
    "crypto/sha256"
    "golang.org/x/crypto/ripemd160"
)

func main() {
    publicKey, err := ioutil.ReadFile("public.pem")  
    if err != nil {  
        os.Exit(-1)  
    }  
    privateKey,err := ioutil.ReadFile("private.pem")  
    if err != nil {  
        os.Exit(-1)  
    }

    sum := sha256.Sum256([]byte(publicKey))
    fmt.Printf("sha256: %x\n", sum)

    hasher := ripemd160.New()
    hasher.Write([]byte(privateKey))
    hashBytes := hasher.Sum(nil)
    fmt.Printf("ripemd160: %x\n", hashBytes)

}