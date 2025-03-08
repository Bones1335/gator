package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Bones1335/gator/internal/config"
	"github.com/Bones1335/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	database *database.Queries
	config   *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	st := &state{
		database: dbQueries,
		config:   &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	commandName := args[1]
	commandArgs := args[2:]
	cmd := command{name: commandName, arguments: commandArgs}

	if err := cmds.run(st, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
