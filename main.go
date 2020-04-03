package main

import (
	"os"
)

var (
	version = "Alpha"
	date = "2020/01"
	encryptopt = "1"
	decryptopt = "2"


)

func main() {
	argsWithProg := os.Args
	orchestrator(argsWithProg)
}



