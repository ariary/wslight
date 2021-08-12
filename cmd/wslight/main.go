// server.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func ParseCommandAndArgs(message string) (cmd string, args []string) {

	msgSlice := strings.Split(message, " ")
	cmd = msgSlice[0]
	if len(msgSlice) > 1 {
		args = msgSlice[1:]
	}
	return cmd, args
}

func TranslateAndExec(command string, arguments []string) string {
	cmd := exec.Command(command, arguments...)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received command: %s", message)
		cmd, args := ParseCommandAndArgs(string(message))
		output := TranslateAndExec(cmd, args)
		//Send well received:
		err = conn.WriteMessage(messageType, []byte(output))
		if err != nil {
			log.Println("Error during writing to websocket:", err)
			return
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
