package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"wslight/pkg/utils"

	prompt "github.com/c-bata/go-prompt"
)

// const arrow ="❯ "
const arrow = "» "

type FSContext struct {
	path  string
	debug bool
}

var ctx *FSContext

// suggestions list
var suggestions = []prompt.Suggest{
	// General
	{"exit", "Exit adretctl"},
	{"help", "get help method"},

	//debug mode
	{"+x", "see which command is launched, equivalent of bash +x"},
	{"-x", "disable debug mode"},

	//linux utility
	{"pwd", "get current directory"},
	{"rm", "remove a file or directory (-f availble for directory)"},
}

//Prefix for the prompt
func livePrefix() (string, bool) {
	// return ctx.path + arrow, true
	return arrow, true
}

//perform at each loop
func executor(in string) {

	//FOR PIPE LATER
	// cmd := commands.CommandList[i]
	// if i == 0 {
	// 	result = Exec(cmd.CmdName, cmd.Args)
	// } else {
	// 	// pipe
	// 	result = ExecPipe(cmd.CmdName, cmd.Args, result)
	// }

	// type Command struct {
	// 	CmdName string
	// 	Args    []string
	// }

	// type Commands struct {
	// 	CommandList []Command
	// 	Dir         string
	// }

	// func ParseCommandAndArgs(cmdLine string) Commands {
	// 	var cmdList []Command
	// 	//detect Pipe ("|") must be before and after a space
	// 	pipeSlice := strings.Split(cmdLine, " | ") // could put a regex to be better
	// 	for i := 0; i < len(pipeSlice); i++ {
	// 		cmdSlice := strings.Split(pipeSlice[i], " ")
	// 		cmd := Command{}
	// 		cmd.CmdName = cmdSlice[0]
	// 		if len(cmdSlice) > 1 {
	// 			cmd.Args = cmdSlice[1:]
	// 		}
	// 		cmdList = append(cmdList, cmd)
	// 	}
	// 	commands := Commands{cmdList, ""}
	// 	return commands
	// }

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

	//Exec avec blocks sinon execPipe

	in = strings.TrimSpace(in)

	// var method, body string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	// case "keyconfig":
	// 	if len(blocks) < 2 {
	// 		fmt.Println("please enter the key")
	// 	} else {
	// 		ctx.key = blocks[1]
	// 	}
	// 	return
	case "-x":
		ctx.debug = false
	case "+x":
		ctx.debug = true
	case "pwd":
		//retrieve current path see https://stackoverflow.com/questions/44206940/execute-the-cd-command-for-cmd-in-go
		cmdName := "cd"
		if ctx.debug {
			fmt.Println("+++", cmdName)
		}
		cmd := exec.Command("cmd", "/c", cmdName)
		current, err := cmd.Output()
		if err != nil {
			fmt.Println("Error retrieving current directory:", err)
		}
		path := strings.Trim(string(current), "\n")
		path = filepath.Clean(string(current))
		ctx.path = string(path)
		cmd.Process.Kill()
		fmt.Println(ctx.path)
	case "rm":
		var cmdName string
		var filename string

		if utils.Contains(blocks, "-r") {
			cmdName = "rmdir"
			filename = blocks[2]
		} else {
			cmdName = "del"
			filename = blocks[1]
		}

		if ctx.debug {
			fmt.Println("+++", cmdName, filename)
		}
		cmd := exec.Command("cmd", "/c", cmdName, filename)
		// cmd.Dir = ctx.path
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error during rm:", err)
		}
		cmd.Process.Kill()
		fmt.Println(string(output))
	case "help":
		fmt.Println("help not constructed\navailable commands: [press TAB]")
		return
	case "exit":
		fmt.Println("Bye!🕶")
		handleExit()
		os.Exit(0)
	default:
		fmt.Printf("Unknown command: %s", blocks[0])
		fmt.Println()
		return
	}

}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

//Function launch when wslight exit. Mainly use to prevent https://github.com/c-bata/go-prompt/issues/228
func handleExit() {
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()
}

func main() {
	defer handleExit()

	ctx = &FSContext{
		path:  "",
		debug: false,
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(arrow),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("wslight"),
	)
	p.Run()
}
