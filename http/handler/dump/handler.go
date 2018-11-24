package dump

import (
	"net/http"

	"github.com/gorilla/mux"
	httperror "github.com/portainer/libhttp/error"
)

// Handler is the http handler used to handle request dump operations
type Handler struct {
	*mux.Router
}

// NewHandler creates an instance of Handler
func NewHandler() *Handler {
	handler := &Handler{
		Router:         mux.NewRouter(),
	}

	handler.Handle("/dump", httperror.LoggerHandler(handler.dumpPostRequest)).Methods(http.MethodPost)

	return handler
}
