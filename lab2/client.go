package main

import (
	"bytes"
	"flag"
	"fmt"

	"golang.org/x/crypto/ssh"
)

var (
	user     = flag.String("u", "", "user")
	password = flag.String("p", "", "Password")
	host     = flag.String("host", "localhost", "Host")
	command  = flag.String("cmd", "", "Command")
	port     = flag.Int("port", 22, "Port")
)

func main() {
	flag.Parse()

	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{ssh.Password(*password)},
	}
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(err)
	}

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	session.Run(*command)
	fmt.Print(stdoutBuf.String())
	stdoutBuf.Reset()
}
