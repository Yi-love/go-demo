package main 

import (
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "crypto/rand"
    "log"
    "os"
)

func GenRsaKey() error {
    //生产私钥
    privateKey , err := rsa.GenerateKey(rand.Reader , 32)
    if err != nil {
        return err
    }

    derStream := x509.MarshalPKCS1PrivateKey(privateKey)
    block := &pem.Block {
        Type: "private key",
        Bytes: derStream,
    }
    file , err := os.Create("private.pem")
    if err != nil {
        return err
    }

    err = pem.Encode(file , block)
    if err != nil {
        return err
    }

    publicKey := &privateKey.PublicKey
    derPkix , err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        return err
    }

    block = &pem.Block {
        Type: "public key",
        Bytes: derPkix,
    }
    file , err = os.Create("public.pem")
    if err != nil {
        return err
    }

    err = pem.Encode(file , block)

    if err != nil {
        return err
    }

    return nil
}

func main() {
    if err := GenRsaKey(); err != nil {  
        log.Fatal("密钥文件生成失败！")  
    }  
    log.Println("密钥文件生成成功！")   
}





