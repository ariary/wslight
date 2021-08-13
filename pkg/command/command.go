package command

import "strings"

type Command struct {
	CmdName string
	Args    []string
}

type Commands struct {
	CommandList []Command
	Dir         string
}

func ParseCommandAndArgs(cmdLine string) Commands {
	var cmdList []Command
	//detect Pipe ("|") must be before and after a space
	pipeSlice := strings.Split(cmdLine, " | ") // could put a regex to be better
	for i := 0; i < len(pipeSlice); i++ {
		cmdSlice := strings.Split(pipeSlice[i], " ")
		cmd := Command{}
		cmd.CmdName = cmdSlice[0]
		if len(cmdSlice) > 1 {
			cmd.Args = cmdSlice[1:]
		}
		cmdList = append(cmdList, cmd)
	}
	commands := Commands{cmdList, ""}
	return commands
}
