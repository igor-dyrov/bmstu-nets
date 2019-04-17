package main

import (
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	ssh.Handle(func(s ssh.Session) {

		// io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))

		term := terminal.NewTerminal(s, "> ")
		term.Write([]byte("ls"))
		for {
			line, err := term.ReadLine()
			if err != nil {
				break
			}
			fmt.Println(line)
		}
		log.Println("terminal closed")
	})

	log.Println("starting ssh server on port 2207...")
	log.Fatal(ssh.ListenAndServe(":2207", nil))
}
