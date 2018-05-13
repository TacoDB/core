package server

import (
	"net"
	"bufio"
)

type Connection struct {
	Conn net.Conn
	Server *Server
}

func (c *Connection) Listen() {
	reader := bufio.NewReader(c.Conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			c.Conn.Close()
			c.Server.onConnectionClosed(c, err)
			return
		}
		c.Server.onMessageReceived(c, message)
	}
}

func (c *Connection) WriteString(str string) error {
	_, err := c.Conn.Write([]byte(str + "\n"))
	return err
}

func (c *Connection) Close() error {
	return c.Conn.Close()
}
