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
	//cmd.Dir = filepath.Clean(ctx.Path)
	out, err := cmd.CombinedOutput()
	cmd.Process.Kill()
	output := string(out)
	if err != nil {
		output = "Error launching command: " + output
	}
	fmt.Println(output)
}

func ExecPipe(commands []string, ctx *Context) {
	full := strings.Join(commands, " | ")
	// full = "'" + full + "'" ??
	Exec(full, ctx)
}
