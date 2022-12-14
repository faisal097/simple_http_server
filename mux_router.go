package simple_http_server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	once *sync.Once
	mux  *mux.Router
	apis Apis
}

func (m *MuxRouter) New() {
	m.once.Do(func() {
		m.mux = mux.NewRouter()
	})
}

func (m *MuxRouter) Start(addr string) error {
	srv := &http.Server{
		Handler: m.mux,
		Addr:    addr,
	}
	log.Println("Http server started listening at ", addr)
	return srv.ListenAndServe()
}

func (m *MuxRouter) AddRoutes() {
	m.mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		get_handler(m.apis, w, r)
	})
}
