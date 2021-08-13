package server

import (
	"fmt"
	"net/http"
	"wslight/pkg/command"
)

func HandleCmd(w http.ResponseWriter, r *http.Request) {
	var commands command.Commands
	err := DecodeBody(w, r, &commands)
	if err != nil {
		fmt.Println("HandleCmd:", err)
	}
	fmt.Println("(debug) ctx:", commands.Dir)

	var result string
	//pipe
	for i := 0; i < len(commands.CommandList); i++ {
		cmd := commands.CommandList[i]
		if i == 0 {
			result = Exec(cmd.CmdName, cmd.Args)
		} else {
			// pipe
			result = ExecPipe(cmd.CmdName, cmd.Args, result)
		}

	}
	fmt.Fprintf(w, result)
}
