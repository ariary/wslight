package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"wslight/pkg/command"
)

var Prefix string
var Current string

//Function to check if remote host is alive. It also retrieve the current directory
func Ping(remote string) {
	//check if host is alive
	fmt.Printf("Checking if host (%s) is alive...", remote)
	fmt.Println()
	_, err := http.Get(remote) //could also use net.DialTimeout
	if err != nil {
		fmt.Println("Ping: connect: connection refused:", err)
		return
	}
	fmt.Println("Connection established! retrieve current directory...")
	Prefix = "❯ "
	pwd := command.Command{"pwd", nil}
	var pwdIntoCmds []command.Command
	pwdIntoCmds = append(pwdIntoCmds, pwd)
	Current, err = SendCommandRequest(command.Commands{pwdIntoCmds, Current})
	if err != nil {
		fmt.Println("Failed retrieving current directory:", err)
	} else {
		Prefix = strings.TrimSuffix(Current, "\n") + Prefix
	}
	fmt.Print(Prefix)
}

func ChangePrefix(path string) {
	Current, err := filepath.Abs(Current + path)

	if err != nil {
		fmt.Println("ChangePrefix:", err)
	}

	Prefix = strings.TrimSuffix(Current, "\n") + "❯ "
	fmt.Println(Prefix)
}

func ParseInputIntoCommandSJSON(reader *bufio.Reader) command.Commands {
	input, err := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	if err != nil {
		fmt.Println(err)
	}
	commands := command.ParseCommandAndArgs(input)
	commands.Dir = Current
	return commands
}

func SendCommandRequest(input command.Commands) (result string, err error) {
	//perform request
	cmdsJSON, err := json.Marshal(input)
	if err != nil {
		fmt.Println("SendCommandRequest:", err)
	}
	body := []byte(cmdsJSON)
	host := os.Getenv("RHOST")
	port := os.Getenv("RPORT")
	endpoint := "http://" + host + ":" + port + "/cmd"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//Read response
	bodyRes, _ := ioutil.ReadAll(resp.Body)
	// bad status code (!=200, logically 404)
	if !strings.Contains(resp.Status, "200") {
		//Error on exec
		return "", err
	}

	return string(bodyRes), nil
}

func ReadResponse(response string) {
	fmt.Print(response)
}
