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
	fmt.Println(Gree("////	 -E  to encrypt	////"))
	fmt.Println(Gree("////	 -D  to decrypt ////"))
	fmt.Println(Gree("////	 -mk  masterkey	////"))
	fmt.Println(Gree("////	 -secret secret ////"))
	fmt.Println(Gree("////	 -h  this menu	////"))
	fmt.Println(Gree("////////////////////////////"))
	dummie()
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
func dummie(){
	fmt.Println(Gree("////////////////////////////////////////////////////////////////////"))
	fmt.Println(Gree("////	 -h for more information				////"))
	fmt.Println(Gree("////	 blackvault -h						////"))
	fmt.Println(Gree("////	 blackvault -E -mk YourMasterKey -secret yoursecret	////"))
	fmt.Println(Gree("////////////////////////////////////////////////////////////////////"))
}

