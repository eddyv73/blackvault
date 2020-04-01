package main

import (
	"fmt"
	"os"
)

var (
	version = "Alpha"
	date = "2020/01"


)

func main() {
	argsWithProg := os.Args
	orchestrator(argsWithProg)
}

func orchestrator(optionrunning []string) {
	optiona := contains(optionrunning,"-a")
	fmt.Println(optiona)
	maindisplay()
	capturerdisplay()
	GoodBye()
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
