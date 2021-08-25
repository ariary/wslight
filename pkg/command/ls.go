package command

import (
	"strings"
	"wslight/pkg/utils"
)

//Get flag of the ls command and translate them
func TranslateLsFlags(args []string) (flags string) {
	// -a ?
	if !utils.Contains(args, "-l") {
		flags += " /b "
	}
	if utils.Contains(args, "-R") {
		flags += " /s "
	}
	if utils.Contains(args, "-a") {
		flags += " /a:h " //ou alors command attrib
	}

	return flags
}

//Retrieve filename from ls command
func ParseFilename(args []string) (filename string) {
	// get args wich are not flags
	var noflags []string
	for i := 0; i < len(args); i++ {
		if !strings.HasPrefix(args[i], "-") {
			noflags = append(noflags, args[i])
		}
	}
	if len(noflags) > 1 {
		//debug purpose (no filename provided so use of stdin)
		filename = strings.Join(noflags, " ")
	}
	//parse them
	return filename
}
