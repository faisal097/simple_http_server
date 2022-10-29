package simple_http_server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/faisal097/simple_http_server/models"
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

	bytes, err := json.Marshal(models.PLACE_HOLDER_EMPLOYEES)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error cooured: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(bytes))
}
