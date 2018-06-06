package server

import (
	"bufio"
	"net"

	"github.com/sandeepravi/muninndb/database"
	"github.com/sandeepravi/muninndb/resp"
)

// Handler handles each individual command
func Handler(conn net.Conn, s *Server) {
	w := bufio.NewWriter(conn)
	defer w.Flush()

	c := &Client{
		s: s,
		w: w,
	}

	n := 0
	db, ok := s.dbs[n]
	if !ok {
		db = database.NewDB(n)
		s.dbs[n] = db
	}

	c.DB = db
	for {
		m, err := resp.NewReader(conn).ReadMessage()

		if err != nil {
			return
		}

		c.Args, _ = m.Strings()

		if cmd, ok := s.cmds[c.Args[0]]; ok {
			cmd.f(c)
		} else {
			c.RespUniqueError("unknown command '" + c.Args[0] + "'")
		}

		w.Flush()
	}
}
