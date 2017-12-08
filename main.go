package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to the Monkey programming language!\n", user.Username)
	fmt.Printf("Type in a command to evaluate it\n")
	repl.Start(os.Stdin, os.Stdout)
}