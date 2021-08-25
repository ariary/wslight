package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"wslight/pkg/command"
	"wslight/pkg/utils"

	prompt "github.com/c-bata/go-prompt"
)

// const arrow ="❯ "
const arrow = "» "

var ctx *command.Context

//Prefix for the prompt
func livePrefix() (string, bool) {
	return ctx.Path + arrow, true
	// return arrow, true
}

//perform at each loop
func executor(in string) {

	//Parse input
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")

	//Parse cmd
	commands := command.ParseCommandAndArgs(blocks)
	var commandsTranslated []string
	for i := 0; i < len(commands.CommandList); i++ {
		cmdTranslate, specialCmd := command.Translate(commands.CommandList[i], ctx)
		if specialCmd {
			//no need to exec
			return
		} else if cmdTranslate == "" {
			// Failed to translate/unknown command
			fmt.Println("Unknown command: failed to translate", commands.CommandList[i].CmdName)
			return
		} else {
			commandsTranslated = append(commandsTranslated, cmdTranslate)
		}
	}

	//execute cmd
	if len(commandsTranslated) > 1 {
		command.ExecPipe(commandsTranslated, ctx)
	} else if len(commandsTranslated) == 1 {
		command.Exec(commandsTranslated[0], ctx)
	} else {
		fmt.Println("Error while parsing command. We couldn't find windows any command corresponding")
	}

}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(utils.Suggestions, w, true)
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

	ctx = &command.Context{
		Path:         "",
		PreviousPath: "",
		Debug:        false,
		Env:          make(map[string]string),
	}

	ctx.RetrieveRootDir()

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(arrow),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("wslight"),
	)
	p.Run()
}
