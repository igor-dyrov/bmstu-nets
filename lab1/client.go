package main

import (
	"fmt" 
	"bufio"
	"strings" 
	"os"
	"log" 
	"io"
	"github.com/jlaffaye/ftp"
)

func CheckErr(err error) {
	if err != nil {
		log.Println("permission denied")
	}
}

func main() {

	connection, err := ftp.Connect("localhost:2121")

	if err != nil {
		log.Fatal("wrong adress\n")
	}

	err = connection.Login("admin", "123456")

	if err != nil {

		log.Fatal("wrong login/password\n")
		return

	}

	consolereader := bufio.NewReader(os.Stdin)
	path := ""

	for {
		fmt.Print("Enter your command : ")

		input, err := consolereader.ReadString('\n')

		if err != nil {
		     fmt.Println(err)
		     os.Exit(1)
		}

		input = input[0: len(input) - 1]
		arguments := strings.Split(input, " ")

		switch arguments[0] {
		case "ls" :
			entries, err := connection.NameList(path)

			if err != nil {
				log.Println("permission denied")
			} else {
				for _, str := range entries {
					fmt.Println(str)
				}				
			}

		case "cd" :
			err = connection.ChangeDir(path + "/" + arguments[1]) 
			if err != nil {
				log.Println("permission denied")
			} else {
				path = path + "/" + arguments[1]
			}

		case "mkdir":
			err = connection.MakeDir(arguments[1])
			CheckErr(err)

		case "rm": 
			err = connection.Delete(arguments[1])
			CheckErr(err)

		case "upload":
			f, err := os.Open(arguments[1])
			if err != nil {
				fmt.Println("cannot open file")
			} else {
				err = connection.Stor(path + "/" + arguments[1], f)
				CheckErr(err)
			}

		case "get":
			file, err := os.Create("./" + arguments[1])
			if err != nil {
				fmt.Println("cannot create file")
			} else {
				resp, err := connection.Retr(path + "/" + arguments[1])
				if err != nil {
					fmt.Println("permission denied")
				} else {
					bytes := make([]byte, 1024)
					for {
						n, err := resp.Read(bytes)
						if n != 0 {
							_, err = file.Write(bytes[0: n])
						}
						if err == io.EOF {
							break
						}
					}
					resp.Close()
					file.Close()
				} 
			}

		case "exit": 
			return
		}
	}
}
