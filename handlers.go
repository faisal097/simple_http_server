package simple_http_server

import (
	"fmt"
	"net/http"
)

// Type Server has implemented Apis interface
type Server struct {
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to %s", r.RequestURI)
}
