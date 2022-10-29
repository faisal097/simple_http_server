package simple_http_server

import "net/http"

func get_handler(apis Apis, w http.ResponseWriter, r *http.Request) {
	apis.Get(w, r)
}
