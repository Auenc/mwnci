package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/auenc/mwnci/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("S'mae %s! Mae hi'n iath y rhaglennu Mwnci\n", user.Username)
	fmt.Printf("Teipio gorchmynion os dych chi eisiau\n")
	repl.Start(os.Stdin, os.Stdout)
}
