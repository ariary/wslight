package command

import (
	"fmt"
	"os/exec"
	"strings"
)

func Exec(command string, ctx *Context) {
	//debug
	if ctx.Debug {
		fmt.Println("+++++", command)
	}

	cmd := exec.Command("cmd", "/c", command)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error retrieving current directory:", err)
	}
	cmd.Process.Kill()
	output := string(out)
	fmt.Println(output)
}

func ExecPipe(commands []string, ctx *Context) {
	full := strings.Join(commands, " | ")
	// full = "'" + full + "'" ??
	Exec(full, ctx)
}
