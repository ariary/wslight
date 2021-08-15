package command

import (
	"strings"
)

type Command struct {
	CmdName string
	Args    []string
}

type Commands struct {
	CommandList []Command
	Dir         string
}

//Parse command to get the name and the argument
// it could detect pipe to parse multiple cmd line [Pipe detection ("|") must be before and after a space]
func ParseCommandAndArgs(cmdLineSlice []string) Commands {
	cmdLine := strings.Join(cmdLineSlice, " ")
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

// func ExecPipe(command string, arguments []string, previousresult string) string {

// 	cmd := exec.Command(command, arguments...)
// 	cmd.Stdin = strings.NewReader(previousresult)
// 	if runtime.GOOS == "windows" {
// 		cmd = Translate(command, arguments)
// 	}
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		//don't exit
// 		fmt.Printf("cmd.Run() failed with %s\n", err)
// 	}
// 	return string(out)
// }
