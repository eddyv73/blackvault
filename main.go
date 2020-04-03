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

func orchestrator(optionrunning []string) {
	help := contains(optionrunning,"-h")
	if help {
		maindisplay()
	}
	encrypt := contains(optionrunning,"-E")
	masterkey := contains(optionrunning,"-mk")
	secret := contains(optionrunning,"-secret")
	if encrypt {
		if masterkey &&  secret {
			masterkey := optionrunning[3]
			secret := optionrunning[5]
			capturerdisplay(masterkey,encryptopt,secret)
		}
	}else{
		decrypt := contains(optionrunning,"-D")
		dmasterkey := contains(optionrunning,"-mk")
		dsecret := contains(optionrunning,"-secret")
		if decrypt  {

			if dmasterkey &&  dsecret {
				masterkey := optionrunning[3]
				secret := optionrunning[5]
				capturerdisplay(masterkey,decryptopt,secret)
			}
		}
	}
	if !help && !encrypt {
		dummie()
	}
}

