package main

import (
	"fmt"
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
func capturerdisplay(masterkey string , secret string, mode string) {

	var result = makehash(secret, checksum(masterkey), mode)
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

