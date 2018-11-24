package dump

import (
	"net/http"
	"fmt"
	"net/http/httputil"

	httperror "github.com/portainer/libhttp/error"
	"github.com/portainer/libhttp/response"
)

func (handler *Handler) dumpPostRequest(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
	  fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	return response.Empty(w)
}
