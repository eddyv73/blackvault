package main

import (
	"bufio"
	"fmt"
	"os"
)

func maindisplay(){

	fmt.Println(Gree("////////////////////////////"))
	fmt.Println(Gree("////Running first ",version,"////"))
	fmt.Println(Gree("////////////////////////////"))
	fmt.Println(Gree("////	Blackvault	////"))
	fmt.Println(Gree("////	 Wister		////"))
	fmt.Println(Gree("////	",date,"		////"))
	fmt.Println(Gree("////////////////////////////"))
	fmt.Println(Gree("////	 1 - to encrypt	////"))
	fmt.Println(Gree("////	 2 - to decrypt ////"))
	fmt.Println(Gree("////////////////////////////"))
}
func capturerdisplay() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(Info("Enter option: "))
	var text, _ = reader.ReadString('\n')

	fmt.Print(Info("Enter your master password: "))
	var key, _ = reader.ReadString('\n')

	fmt.Println(Info("Enter the secret to hash"))
	var plaintexttext, _ = reader.ReadString('\n')
	var result = makehash(text, checksum(key), plaintexttext)
	if result {
		fmt.Println(Info("Exec Correct"))
	} else {
		fmt.Println(Info("Exec Failed"))
	}
}
func GoodBye(){
	fmt.Println(Gree("////////////////////////////"))
	fmt.Println(Gree("////	 GoodBye	////"))
	fmt.Println(Gree("////	  V		////"))
	fmt.Println(Gree("////////////////////////////"))
}

