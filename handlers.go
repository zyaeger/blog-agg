package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("one argument expected for login: username")
	}

	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("The user has been set to:", username)

	return nil
}
