package main

import "strings"

func clearinput(input string) string {
	var cleartext = strings.Replace(input, "\n", "", -1)
	return cleartext
}
func completingstring(input string) string {
	var str strings.Builder
	for len(str.String()) <= 24 {
		str.WriteString("(A)")
	}
	//this char can be var
	return str.String()
}

