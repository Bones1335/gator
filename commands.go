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

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f, exists := c.commands[cmd.name]; exists == true {
		return f(s, cmd)
	} 
	return fmt.Errorf("command %v not found", cmd.name)

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
