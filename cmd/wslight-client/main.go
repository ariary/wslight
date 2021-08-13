//
// Package main is the entry-point for the go-sockets client sub-project.
// The go-sockets project is available under the GPL-3.0 License in LICENSE.
package main

import (
	"bufio"
	"fmt"
	"os"
	"wslight/pkg/client"
)

// Application constants, defining host, port, and protocol.
// const (
// 	connHost = "localhost"
// 	connPort = "8080"
// 	connType = "tcp"
// )

func main() {
	// config remote
	os.Setenv("RHOST", "localhost")
	os.Setenv("RPORT", "8080")
	host := os.Getenv("RHOST")
	port := os.Getenv("RPORT")
	remote := "http://" + host + ":" + port
	//ping
	client.Ping(remote)

	//infinite loop
	for {
		reader := bufio.NewReader(os.Stdin)
		cmd := client.ParseInputIntoCommandSJSON(reader)
		cmdName := cmd.CommandList[0].CmdName
		res, _ := client.SendCommandRequest(cmd)
		if cmdName == "cd" {
			if res == "" {
				dir := cmd.CommandList[0].Args[0]
				client.ChangePrefix(dir)
			}
		} else {
			client.ReadResponse(res)
			fmt.Print(client.Prefix)
		}

	}
}
