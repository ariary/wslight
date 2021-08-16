package command

import (
	"strings"
	"wslight/pkg/utils"
)

//Translate Unix command to cmd/powershell one
// If it is a special it return true, like this we know we don't have to execute it
func Translate(c Command, ctx *Context) (cmd string, special bool) {
	switch c.CmdName {
	case "-x":
		ctx.Debug = false
		special = true
	case "+x":
		ctx.Debug = true
		special = true
	case "pwd":
		//retrieve current path see https://stackoverflow.com/questions/44206940/execute-the-cd-command-for-cmd-in-go
		cmd = "cd"
	case "rm":
		var cmdName string
		var filename string

		if utils.Contains(c.Args, "-r") {
			cmdName = "rmdir"
			filename = c.Args[1]
		} else {
			cmdName = "del"
			filename = c.Args[0]
		}
		full := []string{cmdName, filename}
		cmd = strings.Join(full, " ")
	case "grep":
		cmdName := "findstr"
		flags, recursive := TranslateGrepFlags(c.Args)
		pattern, filename := ParsePatternAndFilename(c.Args)
		if recursive {
			filename = "*.*"
		}
		full := []string{cmdName, flags, pattern, filename}
		cmd = strings.Join(full, " ")
	}
	return cmd, special
}
