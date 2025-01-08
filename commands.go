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
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f, exists := c.handlers[cmd.name]; exists {
		return f(s, cmd)
	}
	return fmt.Errorf("command %v not found", cmd.name)

}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("not enough arguments")
	}

	username := cmd.arguments[0]

	if err := s.config.SetUser(username); err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Printf("User %v has been set\n", username)
	return nil
}
