package server

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Translate(command string, arguments []string) *exec.Cmd {
	cmd := exec.Command("tasklist")
	return cmd
}

func Exec(command string, arguments []string) string {
	if command == "cd" {
		// special case
		file, err := os.Open(arguments[0])

		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		fi, err := file.Stat()
		switch {
		case err != nil:
			fmt.Println(err)
		case fi.IsDir():
			return ""
		default:
			return fmt.Sprintf("cd: %s: Not a directory", arguments[0])
		}
	}
	cmd := exec.Command(command, arguments...)
	if runtime.GOOS == "windows" {
		cmd = Translate(command, arguments)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		//don't exit
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}

func ExecPipe(command string, arguments []string, previousresult string) string {

	cmd := exec.Command(command, arguments...)
	cmd.Stdin = strings.NewReader(previousresult)
	if runtime.GOOS == "windows" {
		cmd = Translate(command, arguments)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		//don't exit
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}
