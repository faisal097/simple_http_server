package simple_http_server

import (
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
	return
}

func (m *MuxRouter) Init() {
	m.mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		m.apis.Get(w, r)
	})
}

func (m *MuxRouter) Start(Addr string) error {
	srv := &http.Server{
		Handler: m.mux,
		Addr:    Addr,
	}
	return srv.ListenAndServe()
}
