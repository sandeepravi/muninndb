package resp

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Reader struct {
	br *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	var buf *bufio.Reader
	buf = bufio.NewReader(r)
	line, err := buf.ReadBytes('\n')
	if err != nil {
		return nil
	}
	if line[0] != 42 {
		b, err := convertToBulk(line)
		if err != nil {
			return nil
		}
		r = bytes.NewReader(b)
	}
	return &Reader{
		bufio.NewReader(r),
	}
}

func (r *Reader) ReadMessage() (*Message, error) {
	return UnpackFromReader(r.br)
}

func UnpackFromReader(r *bufio.Reader) (*Message, error) {
	line, e := r.ReadBytes('\n')
	if e != nil {
		return nil, e
	}

	line = line[:len(line)-2]
	switch line[0] {

	case MessageError:
		return &Message{
			Type:  MessageError,
			Error: errors.New(string(line[1:])),
		}, nil

	case MessageStatus:
		return &Message{
			Type:   MessageStatus,
			Status: string(line[1:]),
		}, nil

	case MessageInt:
		n, err := strconv.ParseInt(string(line[1:]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &Message{
			Type:    MessageInt,
			Integer: n,
		}, nil

	case MessageBulk:
		l, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return nil, err
		}

		if l < 0 {
			return &Message{
				Bulk: nil,
				Type: MessageBulk,
			}, nil
		}

		buf := make([]byte, l+2)
		if _, err := io.ReadFull(r, buf); err != nil {
			return nil, err
		}
		return &Message{
			Bulk: buf[:l],
			Type: MessageBulk,
		}, nil

	case MessageMutli:
		l, e := strconv.Atoi(string(line[1:]))
		if e != nil {
			return nil, e
		}

		if l < 0 {
			return &Message{Type: MessageMutli}, nil
		}
		ret := make([]*Message, l)
		for i := 0; i < l; i++ {
			m, err := UnpackFromReader(r)
			if err != nil {
				return nil, err
			}
			ret[i] = m
		}
		return &Message{
			Type:  MessageMutli,
			Multi: ret,
		}, nil
	}
	return nil, errors.New("Received illegal data from redis.")
}

func convertToBulk(b []byte) ([]byte, error) {
	m := strings.Split(string(b[:]), " ")
	line, e := PackCommand(interfaceArray(m)...)
	if e != nil {
		return nil, e
	}
	return line, nil
}

func interfaceArray(s []string) []interface{} {
	is := make([]interface{}, len(s))
	for i, s := range s {
		is[i] = s
	}
	return is
}
