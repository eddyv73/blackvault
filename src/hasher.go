package main
import (
	"fmt"
    "crypto/sha1"
    "encoding/hex"
)

func hasher(tohash string ) {

    h := sha1.New()
    h.Write([]byte(tohash))
    sha1_hash := hex.EncodeToString(h.Sum(nil))

    fmt.Println(tohash, sha1_hash)
}