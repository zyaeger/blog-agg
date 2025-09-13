package main

import (
	"log"
	"os"

	"github.com/zyaeger/blog-agg/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state{
		cfg: &cfg,
	}

	cmds := commands{
		cmdToHandler: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		log.Fatalf("At least 2 arguments must be given")
	}

	cmd := command{
		Name: cliArgs[1],
		Args: cliArgs[2:],
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
