package server

import "strconv"

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

func kvincr(c *Client) {
	v, ok := c.DB.Get(c.Args[1])
	var i int
	if !ok {
		i = 1
	} else {
		switch s := v.(type) {
		default:
			c.RespTypeError()
			return
		case string:
			var err error
			i, err = strconv.Atoi(s)
			if err != nil {
				c.RespTypeError()
				return
			}
			i += int(1)
		}
	}
	c.DB.Update(c.Args[1], strconv.Itoa(i))
	c.RespInt(int(i))
}
