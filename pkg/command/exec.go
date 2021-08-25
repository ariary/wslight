package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"wslight/pkg/utils"
)

func Exec(command string, ctx *Context) {
	//debug
	if ctx.Debug {
		fmt.Println("+++++", command)
	}

	cmd := exec.Command("cmd", "/c", command)
	cmd.Dir = ctx.Path
	cmd.Env = os.Environ()
	for key, value := range ctx.Env {
		cmd.Env = append(cmd.Env, key+"="+value)
	}

	out, err := cmd.CombinedOutput()
	cmd.Process.Kill()
	output := utils.CleanOutput(string(out))
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
