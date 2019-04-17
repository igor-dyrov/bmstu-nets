package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/ssh"
)

var (
	user     = flag.String("u", "", "user")
	password = flag.String("p", "", "Password")
	host     = flag.String("host", "185.20.227.83", "Host")
	port     = flag.Int("port", 22, "Port")
)

func main() {
	flag.Parse()

	config := &ssh.ClientConfig{
		User: "iu8_81_07",
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

	b, err := session.CombinedOutput("uname -a")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
