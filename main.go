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

	fmt.Println(cfg)

}
