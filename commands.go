package main

import (
	"fmt"

	"github.com/Bones1335/gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("not enough arguments")
	}

	if s.config.SetUser(cmd.arguments[1]) == nil {
		return fmt.Errorf("please provide username")
	}
	s.config.SetUser(cmd.arguments[1])

	fmt.Printf("User %v has been set", s.config.SetUser(cmd.arguments[1]))

	return nil
}
