package main

import (
	"fmt"

	"github.com/Bones1335/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("First read of config:\nDatabase: %v\nUser: %v\n", cfg.DbURL, cfg.CurrentUser)

	cfg.SetUser("alexander")

	newCfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Second read of config:\nDatabase: %v\nUser: %v\n", newCfg.DbURL, newCfg.CurrentUser)
}
