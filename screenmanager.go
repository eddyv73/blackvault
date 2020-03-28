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
/*	fmt.Println("Enter the key")*/
	var keytext = "123456789012345678901234"
	fmt.Println("Enter the secret to hash")
	var plaintexttext, _ = reader.ReadString('\n')
	if  text == "1" {
		callEncrypted(keytext,plaintexttext)
	}else{
		callDecrypt(keytext,plaintexttext)
	}
}

func callEncrypted(keytext string , plaintexttext string ){
	ciphertext, err := encryptString(keytext,plaintexttext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", ciphertext)
}
func callDecrypt(keytext string , plaintexttext string) {

	result,err := decryptString(keytext, plaintexttext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}

