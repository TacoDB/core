package main

import (
	"mvmo.eu/tacoDB-server/server"
	"fmt"
	"net"
)

func server_() {
	tcpServer := server.New(1889)
	go tcpServer.Listen()

	tcpServer.OnMessageReceived(func(c *server.Connection, message string) {
		fmt.Println(message)
	})
	tcpServer.OnConnectionClosed(func(c *server.Connection, err error) {
		fmt.Println(err)
	})
	tcpServer.OnNewConnectionIncoming(func(c *server.Connection) {
		fmt.Println("hel")
	})


}

func main() {
	go server_()
	go func() {
		conn, err := net.Dial("tcp", ":1889")
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte("Hello\n"))
	}()
	select {}
}
