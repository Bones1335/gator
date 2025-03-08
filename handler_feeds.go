package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bones1335/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("not enough arguments")
	}
	feed, err := s.database.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   cmd.arguments[0],
		Url:    cmd.arguments[1],
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error adding feed to database: %w", err)
	}

	feedFollow, err := s.database.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")

	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Printf("New feed record: %v\n", feed)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {

	feeds, err := s.database.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting all feeds: %w", err)
	}

	for _, feed := range feeds {
		name, err := s.database.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}
		fmt.Printf("Name: %v\n", feed.Name)
		fmt.Printf("Url: %v\n", feed.Url)
		fmt.Printf("User: %v\n", name)
	}

	return nil
}
