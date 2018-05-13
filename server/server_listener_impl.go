package server

func (server *Server) OnNewConnectionIncoming(callback func(c *Connection)) {
	server.onNewConnectionIncoming = callback
}

func (server *Server) OnConnectionClosed(callback func(c *Connection, err error)) {
	server.onConnectionClosed = callback
}

func (server *Server) OnMessageReceived(callback func(c *Connection, message string)) {
	server.onMessageReceived = callback
}