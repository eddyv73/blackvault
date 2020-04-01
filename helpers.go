package main

import (
	"fmt"
	"strings"
)
var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)
var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
func clearinput(input string) string {
	var cleartext = strings.Replace(input, "\n", "", -1)
	return cleartext
}
func completingstring(input string) string {
	var str strings.Builder
	str.WriteString(input)
	for len(str.String()) <= 24 {
		str.WriteString("(A)")
	}
	//this char can be var
	return str.String()
}
func cleanstring(s string) string {
	result := strings.ReplaceAll(s, "(A)", "")
	resulta := strings.ReplaceAll(result, "(", "")
	return resulta
}
