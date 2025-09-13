package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/zyaeger/blog-agg/internal/config"
	"github.com/zyaeger/blog-agg/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		cmdToHandler: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)

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
