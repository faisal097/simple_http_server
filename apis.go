package simple_http_server

import (
	"net/http"
)

type Apis interface {
	Get(w http.ResponseWriter, r *http.Request)
}
