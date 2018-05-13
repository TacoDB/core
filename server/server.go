package server

import (
	"log"
	"strconv"
	"net"
)

type Server struct {
	Port int

	onNewConnectionIncoming func(c *Connection)
	onConnectionClosed      func(c *Connection, err error)
	onMessageReceived       func(c *Connection, message string)
}

func New(port int) *Server {
	log.Println("creating new server on port " + strconv.Itoa(port))
	server := &Server{Port: port}

	server.OnNewConnectionIncoming(func(c *Connection) {})
	server.OnMessageReceived(func(c *Connection, message string) {})
	server.OnConnectionClosed(func(c *Connection, err error) {})

	return server
}

func (server *Server) Listen() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.Port))
	if err != nil {
		log.Fatal("Error starting TCP server")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		connection := Connection{Conn: conn, Server: server}

		server.onNewConnectionIncoming(&connection)

		go connection.Listen()

	}
}
