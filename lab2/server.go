package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		term := terminal.NewTerminal(s, "> ")
		for {
			input, _ := term.ReadLine()
			if input == "exit" {
				break
			}
			arguments := strings.Split(input, " ")
			var cmd *exec.Cmd
			if len(arguments) > 1 {
				cmd = exec.Command(arguments[0], arguments[1:len(arguments)]...)
			} else {
				cmd = exec.Command(arguments[0])
			}
			cmd.Stdout = term
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Run()
			fmt.Println([]byte(input))
		}
		log.Println("terminal closed")
	})

	log.Println("starting ssh server on port 2207...")
	log.Fatal(ssh.ListenAndServe(":2207", nil))
}
