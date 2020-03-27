package main

import (
	"bufio"
	"fmt"
	"os"
)
var (
	version = "Alpha"
	date = "2020/01"

)

func main() {
	orchestrator()
}

func orchestrator(){
	maindisplay()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	var text, _ = reader.ReadString('\n')
	var hashes = makehash(text)
	fmt.Println(hashes)
	fmt.Println(text)
}

func maindisplay(){

	fmt.Println("////////////////////////////")
	fmt.Println("////Running first",version,"////")
	fmt.Println("////////////////////////////")
	fmt.Println("////	Blackvault	////")
	fmt.Println("////	 Wister		////")
	fmt.Println("////	",date,"	////")
	fmt.Println("////////////////////////////")
	fmt.Println("////////////////////////////")
}

//func makehash(tohash string ) string {

//	h := sha1.New()
//	h.Write([]byte(tohash))
//	sha1_hash := hex.EncodeToString(h.Sum(nil))

//	fmt.Println(tohash, sha1_hash)
//	return sha1_hash
//}