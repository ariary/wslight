package command

import (
	"errors"
	"fmt"
	"os"
)

//retrieve path asked by cd command, detect if it is a special case (- and ~) determine if it is a directory or file
//Return the retrieved path (an empty string otherwise)
func StatPath(args []string, ctx *Context) (path string, special bool, err error) {
	// Only look first argument
	if len(args) > 1 {
		err = errors.New(fmt.Sprintf("cd: too many argument"))
		return path, special, err
	} else if len(args) == 1 {
		path = args[0]
		//special case
		switch path {
		case "~": //does not handle error
			drive := os.Getenv("HOMEDRIVE")
			homepath := os.Getenv("HOMEPATH")
			path = drive + homepath
			special = true
		case "-":
			path = ctx.PreviousPath
			special = true
		default:
			//determine if it is a directory
			fi, err := os.Stat(path)
			if err != nil {
				err = errors.New(fmt.Sprintf("cd: %s: No such file or directory", path))
				return path, special, err
			}
			mode := fi.Mode()
			if !mode.IsDir() {
				err = errors.New(fmt.Sprintf("cd: %s: Not a directory", path))
				return path, special, err
			}
		}
	} else { //Home
		drive := os.Getenv("HOMEDRIVE")
		homepath := os.Getenv("HOMEPATH")
		path = drive + homepath
	}
	return path, special, err
}
