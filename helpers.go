package main

import "strings"

func clearinput(input string) string {
	var cleartext = strings.Replace(input, "\n", "", -1)
	return cleartext
}

