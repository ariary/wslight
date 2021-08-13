// // server.go
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os/exec"
// 	"runtime"
// 	"strings"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{} // use default options

// func ParseCommandAndArgs(message string) (cmd string, args []string) {

// 	msgSlice := strings.Split(message, " ")
// 	cmd = msgSlice[0]
// 	if len(msgSlice) > 1 {
// 		args = msgSlice[1:]
// 	}
// 	return cmd, args
// }

// func TranslateAndExec(command string, arguments []string) string {
// 	cmd := exec.Command(command, arguments...)
// 	if runtime.GOOS == "windows" {
// 		cmd = exec.Command("tasklist")
// 	}
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// 	return string(out)
// }

// func socketHandler(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade our raw HTTP connection to a websocket based one
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Print("Error during connection upgradation:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// The event loop
// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("Error during message reading:", err)
// 			break
// 		}
// 		log.Printf("Received command: %s", message)
// 		cmd, args := ParseCommandAndArgs(string(message))
// 		output := TranslateAndExec(cmd, args)
// 		//Send well received:
// 		err = conn.WriteMessage(messageType, []byte(output))
// 		if err != nil {
// 			log.Println("Error during writing to websocket:", err)
// 			return
// 		}
// 	}
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Index Page")
// }

// func main() {
// 	http.HandleFunc("/socket", socketHandler)
// 	http.HandleFunc("/", home)
// 	log.Fatal(http.ListenAndServe("localhost:8080", nil))
// }
// Package main is the entry-point for the go-sockets server sub-project.
// The go-sockets project is available under the GPL-3.0 License in LICENSE.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

// Application constants, defining host, port, and protocol.
const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func ParseCommandAndArgs(message string) (cmd string, args []string) {

	msgSlice := strings.Split(message, " ")
	cmd = msgSlice[0]
	if len(msgSlice) > 1 {
		args = msgSlice[1:]
	}
	return cmd, args
}

func Exec(command string, arguments []string) string {
	cmd := exec.Command(command, arguments...)
	if runtime.GOOS == "windows" {
		cmd = Translate(command, arguments)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		//don't exit
		fmt.Printf("cmd.Run() failed with %s\n", err)
		return err.Error()
	}
	return string(out)
}

func Translate(command string, arguments []string) *exec.Cmd {
	cmd := exec.Command("tasklist")
	return cmd
}

func HandleCmd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	cmd, args := ParseCommandAndArgs(string(body))

	result := Exec(cmd, args)
	fmt.Fprintf(w, result)
}

func main() {
	mux := http.NewServeMux()
	//Add handlers
	mux.HandleFunc("/cmd", HandleCmd)

	err := http.ListenAndServe(":"+connPort, mux)
	log.Fatal(err)
}
