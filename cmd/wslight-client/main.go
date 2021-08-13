//
// Package main is the entry-point for the go-sockets client sub-project.
// The go-sockets project is available under the GPL-3.0 License in LICENSE.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Application constants, defining host, port, and protocol.
const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func parseInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	if err != nil {
		fmt.Println(err)
	}
	return input
}

func sendComandRequest(input string) (result string, err error) {
	//perform request
	body := []byte(input)
	endpoint := "http://" + connHost + ":" + connPort + "/cmd"
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

func readResponse(response string) {
	fmt.Print(response)
}

func main() {
	// run loop forever, until exit.
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("❯ ")
		cmd := parseInput(reader)
		res, _ := sendComandRequest(cmd)
		readResponse(res)
	}
}
