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
