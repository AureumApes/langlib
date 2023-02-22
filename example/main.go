package main

import (
	"github.com/AureumApes/langlib"
	"strings"
)

func main() {
	lang := langlib.NewLang("Basic", "bsc")
	lang.AddCmd("print", printCmd)
	file := langlib.ReadLangFile("/mnt/Linux-HDD/Development/gopath/src/github.com/AureumApes/langlib/example/main", lang)
	file.Run()
}

func printCmd(args []string) {
	print(strings.Join(args, " "))
}
