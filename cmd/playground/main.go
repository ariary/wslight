package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

// const arrow ="❯ "
const arrow = "> "

type FSContext struct {
	path string
}

var ctx *FSContext

// suggestions list
var suggestions = []prompt.Suggest{
	// General
	{"exit", "Exit adretctl"},
	{"help", "get help method"},

	// key
	{"keyconfig", "set the key used to decrypt the fs"},
	{"keyprint", "print the current key"},

	// Command on ubac
	{"connect", "Connect to the configured Ubac"}, //in fact launch get and see if there is result
	{"cd", "Change the  working directory in encrypted fs. (Do not support full path)"},

	// Read Method
	{"ls", "list directory contents on remote encrypted fs"},
	{"cat", "print file content on remote encrypted fs resource"},
	{"tree", "print tree of remote encrypte fs"},

	// Write Method
	{"rm", "remove directory or file on remote encrypted fs"},
}

//Prefix for the prompt
func livePrefix() (string, bool) {
	if ctx.path == "/" {
		return "", false
	}

	// //retrieve current path see https://stackoverflow.com/questions/44206940/execute-the-cd-command-for-cmd-in-go
	// cmd := exec.Command("cmd", "/c", "cd")
	// current, err := cmd.Output()
	// cmd.Wait()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// ctx.path = strings.Trim(string(current), "\n")

	return ctx.path + arrow, true
}

//perform at each loop
func executor(in string) {
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
		path: "",
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(ctx.path+arrow),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("wslight"),
	)
	p.Run()
}
