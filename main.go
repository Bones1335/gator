package main

import (
	"fmt"
	"os"

	"github.com/Bones1335/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	st := state{config: &cfg}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	commandName := args[1]
	commandArgs := args[2:]
	cmd := command{name: commandName, arguments: commandArgs}

	if err := cmds.run(&st, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
