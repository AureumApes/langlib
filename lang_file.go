package langlib

import (
	"os"
	"strings"
)

/*
langFile is a struct, which contains important information for a FIle written in your language
*/
type langFile struct {
	absolutePath string
	language     language
	name         string
	content      string
}

/*
ReadLangFile reads a File by its name to a langFile Object
*/
func ReadLangFile(path string, lang language) langFile {
	abPath := path + "." + lang.fileExtension
	name := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	content, err := os.ReadFile(abPath)
	if err != nil {
		panic(err)
	}
	return langFile{
		abPath,
		lang,
		name,
		string(content),
	}
}

/*
Run obviously, runns your file with the language you wrote
*/
func (lf langFile) Run() {
	cmds := strings.Split(lf.content, "\n")
	for _, fullCmd := range cmds {
		cmd := strings.Split(fullCmd, " ")[0]
		args := strings.Split(fullCmd, " ")[1:]
		fun := lf.language.cmd[cmd]
		if fun != nil {
			fun(args)
		}
	}
}
