package command

import (
	"fmt"
	"strings"
	"wslight/pkg/utils"
)

//Search specific grep flag and translate it
//return a string of flag (if recursive flag is detected it returns a boolean at true)
func TranslateGrepFlags(args []string) (flags string, recursive bool) {

	if utils.Contains(args, "-R") {
		flags += " /s "
		recursive = true
	}
	if utils.Contains(args, "-i") {
		flags += " /i "
	}
	return flags, recursive
}

// Retrieve pattern for grep command (first argument with is not a flag). and also the filename
func ParsePatternAndFilename(args []string) (pattern string, filename string) {
	// get args wich are not flags
	var noflags []string
	for i := 0; i < len(args); i++ {
		if !strings.HasPrefix(args[i], "-") {
			noflags = append(noflags, args[i])
		}
	}
	if len(noflags) == 0 {
		fmt.Println("grep: Missing pattern for grep command")
	} else if len(noflags) == 1 {
		//debug purpose (no filename provided so use of stdin)
		//fmt.Println("grep: No filename provided")
		pattern = noflags[0]
	} else {
		pattern = noflags[0]
		filename = noflags[1]
	}
	//parse them
	return pattern, filename
}
