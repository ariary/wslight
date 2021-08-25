package command

import (
	"fmt"
	"os/exec"
	"wslight/pkg/utils"
)

type Context struct {
	Path         string
	PreviousPath string
	Debug        bool
}

// Retrieve current directory by executing cmd /c cd (~ pwd)
func (c *Context) RetrieveRootDir() {
	cmd := exec.Command("cmd", "/c", "cd")
	out, err := cmd.CombinedOutput()
	cmd.Process.Kill()
	//output := strings.Trim(string(out), "\n")
	output := utils.CleanOutput(string(out))
	if err != nil {
		fmt.Println("Failed retrieving current dir for WSLight:", err)
	}
	c.Path = output
}
