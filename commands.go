package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmdToHandler map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.cmdToHandler[cmd.Name]
	if !exists {
		return fmt.Errorf("command not found: %s", cmd.Name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdToHandler[name] = f
}
