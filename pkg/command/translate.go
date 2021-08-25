package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"wslight/pkg/utils"
)

//https://www.lemoda.net/windows/windows2unix/windows2unix.html

//Translate Unix command to cmd/powershell one
// If it is a special it return true, like this we know we don't have to execute it
func Translate(c Command, ctx *Context) (cmd string, special bool) {
	switch c.CmdName {
	case "help":
		utils.PrintSuggestions()
		special = true
	case "exit":
		os.Exit(0)
	case "-x":
		ctx.Debug = false
		special = true
	case "+x":
		ctx.Debug = true
		special = true
	case "hostname":
		//no translation, no arg and specific flag add
		return c.CmdName, false
	case "pwd":
		//retrieve current path see https://stackoverflow.com/questions/44206940/execute-the-cd-command-for-cmd-in-go
		cmd = "cd"
	case "rm":
		var cmdName string
		var filename string

		if utils.Contains(c.Args, "-r") || utils.Contains(c.Args, "-R") {
			cmdName = "rmdir /Q /S"
			filename = c.Args[1]
		} else {
			cmdName = "del"
			filename = c.Args[0]
		}
		full := []string{cmdName, filename}
		cmd = strings.Join(full, " ")
	case "cp":
		var cmdName string
		var src, dst string

		if utils.Contains(c.Args, "-r") || utils.Contains(c.Args, "-R") {
			cmdName = "xcopy /I"
			src = c.Args[1]
			dst = c.Args[2]
		} else {
			cmdName = "copy"
			src = c.Args[0]
			dst = c.Args[1]
		}
		full := []string{cmdName, src, dst}
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
	case "cat":
		cmdName := "type"
		var filename string
		if c.Args != nil {
			filename = c.Args[0]
		}

		full := []string{cmdName, filename}
		cmd = strings.Join(full, " ")
	case "ls":
		cmdName := "dir"
		flags := TranslateLsFlags(c.Args)
		filename := ParseFilename(c.Args)
		full := []string{cmdName, flags, filename}
		cmd = strings.Join(full, " ")
	case "tree":
		cmdName := "tree /f "
		cmd = cmdName + strings.Join(c.Args, " ")
	case "env":
		cmdName := "set "
		if len(c.Args) > 0 {
			err := updateEnvContext(c.Args, ctx)
			if err != nil {
				fmt.Println(err)
			}
			special = true
		}
		cmd = cmdName + strings.Join(c.Args, " ")
	case "cd":
		special = true
		//stat file
		path, special, err := StatPath(c.Args, ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			//update path accordingly
			ctx.PreviousPath = ctx.Path
			if special {
				//absolute path
				ctx.Path = path
			} else {
				ctx.Path = filepath.Clean(ctx.Path + "\\" + path)
			}

		}
	}
	return cmd, special
}
