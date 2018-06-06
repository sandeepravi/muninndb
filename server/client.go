package server

import (
	"io"
	"strconv"

	"github.com/sandeepravi/muninndb/database"
)

type Client struct {
	s *Server
	// The active db
	DB *database.DB

	Args []string

	w io.Writer
}

func (c *Client) RespUniqueError(s string) {
	io.WriteString(c.w, "-"+s+"\r\n")
}

func (c *Client) RespString(s string) {
	io.WriteString(c.w, "+"+s+"\r\n")
}

func (c *Client) RespBulk(s string) {
	io.WriteString(c.w, "$"+strconv.FormatInt(int64(len(s)), 10)+"\r\n"+s+"\r\n")
}

func (c *Client) RespNull() {
	io.WriteString(c.w, "$-1\r\n")
}

func (c *Client) RespTypeError() {
	c.RespUniqueError("WRONGTYPE Operation against a key holding the wrong kind of value")
}
