package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.database.Reset(ctx)
	if err != nil {
		return fmt.Errorf("error resetting database: %w", err)
	}
	fmt.Println("Database reset successfully")
	return nil
}
