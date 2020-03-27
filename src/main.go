package src

import (
	"bufio"
	"fmt"
	"os"
)

var (
	version = "Alpha"
	date = "2020/01"

)

func main()  {
	orchestrator()
}
func orchestrator(){
	maindisplay()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	var text, _ = reader.ReadString('\n')
	var hashes = hasher(text)
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
