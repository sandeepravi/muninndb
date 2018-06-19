package server

import (
	"fmt"
	"strings"
)

var commands map[string]*Command

type Command struct {
	name string
	f    func(c *Client)
}

func SetupCommands() {
	// String
	registerCmd("SET", kvset)
	registerCmd("GET", kvget)

	// integer
	registerCmd("INCR", kvincr)
}

func registerCmd(name string, f func(c *Client)) {
	var cmd Command

	cmd.name = name
	cmd.f = f

	fmt.Println(s)
	s.cmds[strings.ToUpper(name)] = &cmd
}
