package handler

import (
	"net/http"
	"strings"

	"github.com/deviantony/reqdump/http/handler/dump"
)

// Handler is the main handler of the application.
type Handler struct {
	dumpHandler *dump.Handler
}

// NewHandler returns a new instance of Handler
func NewHandler() *Handler {
	handler := &Handler{
		dumpHandler: dump.NewHandler(),
	}

	return handler
}

// ServeHTTP delegates a request to the appropriate subhandler.
func (handler *Handler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	switch {
	case strings.HasPrefix(request.URL.Path, "/dump"):
		handler.dumpHandler.ServeHTTP(rw, request)
	}
}
