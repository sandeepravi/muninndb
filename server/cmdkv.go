package server

func kvset(c *Client) {
	c.DB.Set(c.Args[1], c.Args[2])
	c.RespString("OK")
}

func kvget(c *Client) {
	v, ok := c.DB.Get(c.Args[1])
	if !ok {
		c.RespNull()
		return
	}
	switch s := v.(type) {
	default:
		c.RespTypeError()
		return
	case string:
		c.RespBulk(s)
	}
}
