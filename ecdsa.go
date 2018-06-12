package main 

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "golang.org/x/crypto/ripemd160"
    "fmt"
    "log"
    "math/big"
)

type Key struct {
    PrivateKey []byte
    PublicKey  ecdsa.PublicKey
    Address    string
}

const privKeyBytesLen = 32

func paddedAppend(size uint, dst, src []byte) []byte {
    for i := 0; i < int(size)-len(src); i++ {
        dst = append(dst, 0)
    }
    return append(dst, src...)
}

// b58encode encodes a byte slice b into a base-58 encoded string.
func b58encode(b []byte) (s string) {
    /* See https://en.bitcoin.it/wiki/Base58Check_encoding */

    const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

    /* Convert big endian bytes to big int */
    x := new(big.Int).SetBytes(b)

    /* Initialize */
    r := new(big.Int)
    m := big.NewInt(58)
    zero := big.NewInt(0)
    s = ""

    /* Convert big int to string */
    for x.Cmp(zero) > 0 {
        /* x, r = (x / 58, x % 58) */
        x.QuoRem(x, m, r)
        /* Prepend ASCII character */
        s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
    }

    return s
}

func createKey() Key{
    key := Key{}
    //椭圆算法
    curve := elliptic.P256()
    //生成私钥,公钥
    private, err := ecdsa.GenerateKey(curve, rand.Reader)
    if err != nil {
        log.Panic(err)
    }
    d := private.D.Bytes()
    b := make([]byte, 0, privKeyBytesLen)
    priKet := paddedAppend(privKeyBytesLen, b, d)
    pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

    fmt.Printf("priKet: %02X\n" , priKet)
    fmt.Printf("pubKey: %02X\n" , pubKey)

    //hash公钥
    h := sha256.New()
    h.Write(pubKey)
    fmt.Printf("sha256 pubKey: %02X\n", h.Sum(nil))

    //ripemd160 hash
    hasher := ripemd160.New()
    hasher.Write(h.Sum(nil))
    sha256Hash := hasher.Sum(nil)
    fmt.Printf("公钥哈希值: ripemd160 sha256 hash: %02X\n", sha256Hash)
    
    pubKeyAddrSource := append([]byte{0x00} , sha256Hash...)
    //校验码
    verifyCode := sha256.New()
    verifyCode.Write(pubKeyAddrSource)
    verifyCodeHash := sha256.New()
    verifyCodeHash.Write(verifyCode.Sum(nil))
    verifyCode2 := verifyCodeHash.Sum(nil)
    fmt.Printf("校验码: sha256(sha256) pubKey: %02X\n", verifyCode2[0:4])

    //地址元数据
    pubKeyAddrSource = append(pubKeyAddrSource , verifyCode2[0:4]...)
    fmt.Printf("地址元数据: 0+公钥哈希值+校验码 : %02X\n", pubKeyAddrSource)

    //地址
    pubKeyAddr := b58encode(pubKeyAddrSource)

    for _, v := range pubKeyAddrSource {
        if v != 0 {
            break
        }
        pubKeyAddr = "1" + pubKeyAddr
    }

    fmt.Println("地址: b58encode : ", pubKeyAddr)

    key.PrivateKey = private.D.Bytes()
    key.PublicKey = private.PublicKey
    key.Address = pubKeyAddr

    return key
}


func createSign(pri *ecdsa.PrivateKey , data []byte)(r, s *big.Int, hash []byte, err error) {

    h := sha256.New()
    h.Write(data)
    hash = h.Sum(nil)
    r , s , err = ecdsa.Sign(rand.Reader , pri , hash)
    if err != nil {
        return nil , nil , nil , err
    }
    return r , s , hash , nil
}

func main() {
    key := createKey()
    fmt.Println(key.PrivateKey)

    priKet := ecdsa.PrivateKey{}

    //椭圆算法
    curve := elliptic.P256()
    //生成私钥,公钥
    private, err := ecdsa.GenerateKey(curve, rand.Reader)

    priKet.D = new(big.Int).SetBytes(key.PrivateKey)
    priKet.PublicKey = key.PublicKey

    fmt.Println("pubKey: " , key.PublicKey)
    
    r , s , hash , err := createSign(&priKet , []byte("787878"))

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("r  ,s :" , r , s)

    if ecdsa.Verify(&private.PublicKey ,hash , r , s) {
        fmt.Println("is ok")
    }else {
        fmt.Println("is error")
    }
    
}










