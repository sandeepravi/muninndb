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
	incdecby(c, "1", "+")
}

func kvincrby(c *Client) {
	incdecby(c, c.Args[2], "+")
}

func kvdecr(c *Client) {
	incdecby(c, "1", "-")
}

func kvdecrby(c *Client) {
	incdecby(c, c.Args[2], "-")
}

func incdecby(c *Client, b string, op string) {
	v, ok := c.DB.Get(c.Args[1])
	var i int
	n, err := strconv.Atoi(b)
	if err != nil {
		c.RespTypeError()
		return
	}
	if !ok {
		i = n
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
			if op == "+" {
				i += int(n)
			} else {
				i -= int(n)
			}
		}
	}
	c.DB.Update(c.Args[1], strconv.Itoa(i))
	c.RespInt(int(i))
}
