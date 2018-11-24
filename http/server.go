package http

import (
	"net/http"
	"github.com/deviantony/reqdump/http/handler"
)

// Server is a web server
type Server struct {}

// NewServer returns a pointer to a Server
func NewServer() *Server {
	return &Server{}
}

// Start start the server by listening on the specified listenAddr
func (server *Server) Start(listenAddr string) error {
	h := handler.NewHandler()
	return http.ListenAndServe(listenAddr, h)
}
