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
	fmt.Println(cmd.Dir)
	if err != nil {
		fmt.Println("Error launching command:", err.Error())
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
