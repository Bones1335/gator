package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bones1335/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	name := cmd.arguments[0]

	user, err := s.database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("not enough arguments")
	}

	username := cmd.arguments[0]

	_, err := s.database.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := s.config.SetUser(username); err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Printf("User %v has been set\n", username)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID: %v\n", user.ID)
	fmt.Printf(" * Name: %v\n", user.Name)
}
