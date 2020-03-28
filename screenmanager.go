package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func maindisplay(){

	fmt.Println("////////////////////////////")
	fmt.Println("////Running first",version,"////")
	fmt.Println("////////////////////////////")
	fmt.Println("////	Blackvault	////")
	fmt.Println("////	 Wister		////")
	fmt.Println("////	",date,"	////")
	fmt.Println("////////////////////////////")
	fmt.Println("////	 1 - to encrypt	////")
	fmt.Println("////	 2 - to decrypt ////")
	fmt.Println("////////////////////////////")
}
func capturerdisplay() {


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter option: ")
	var text, _ = reader.ReadString('\n')
	fmt.Println("Enter the key")
	var keytext, _ = reader.ReadString('\n')
	fmt.Println("Enter the secret to hash")
	var plaintexttext, _ = reader.ReadString('\n')
	key := []byte(keytext) // 32 bytes
	plaintext := []byte(plaintexttext)
	if  text == "1" {
		callEncrypted(key,plaintext)
	}else{
		callDecrypt(key,plaintext)
	}
}

func callEncrypted(keytext []byte , plaintexttext []byte ){
	ciphertext, err := encrypt(keytext, plaintexttext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", ciphertext)
}
func callDecrypt(keytext []byte , plaintexttext []byte) {

	result, err := decrypt(keytext, plaintexttext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}

