package main

import (
	"bufio"
	"fmt"
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

	fmt.Print("Enter your master password: ")
	var key, _ = reader.ReadString('\n')

	fmt.Println("Enter the secret to hash")
	var plaintexttext, _ = reader.ReadString('\n')
	var result = makehash(text, checksum(key), plaintexttext)
	if result {
		fmt.Println("Exec Correct")
	} else {
		fmt.Println("Exec Failed")
	}

}

