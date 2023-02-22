package langlib

/*
language contains information about your language, create a new one with the NewLang Function
*/
type language struct {
	name          string
	fileExtension string
	cmd           map[string]func(args []string)
}

/*
NewLang creates a Langauge Object
*/
func NewLang(name string, extension string) language {
	return language{
		name:          name,
		fileExtension: extension,
		cmd:           map[string]func(args []string){},
	}
}

/*
AddCmd adds a Command to your language Object
*/
func (l language) AddCmd(cmdName string, fun func(args []string)) {
	l.cmd[cmdName] = fun
}
