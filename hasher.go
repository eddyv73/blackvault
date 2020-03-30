package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func makehash(text string, keytext string ,plaintexttext string ) bool {
	fmt.Println("hasher running")
	var greenlight = false
	if  clearinput(text) == "1" {
		e := Encrypt([]byte(keytext),clearinput(plaintexttext))
		fmt.Printf("%s\n", e)

		greenlight = true
	}else{
		Decrypt([]byte(keytext),clearinput(plaintexttext))
		greenlight = true
	}
	return greenlight
}
func Encrypt(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}
func Decrypt(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	plain := make([]byte, len(ciphertext))
	c.Decrypt(plain, ciphertext)
	s := string(plain[:])
	fmt.Printf("AES Decrypyed Text:  %s\n", s)
}

