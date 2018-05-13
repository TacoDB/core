package server

type Server struct {
	Connection *Connection

	onNewConnectionIncoming func()
}

func (server *Server) OnNewConnectionIncoming(callback func()) {
	server.onNewConnectionIncoming = callback
}