package src

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func hasher(tohash string ) string {

	h := sha1.New()
	h.Write([]byte(tohash))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	fmt.Println(tohash, sha1_hash)
	return sha1_hash
}