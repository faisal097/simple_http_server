package simple_http_server

import (
	"fmt"
	"sync"
)

type Router interface {
	New()
	AddRoutes()
	Start(Addr string) error
}

func NewRouter(routerType string) (Router, error) {
	switch routerType {
	case "mux":

		r := &MuxRouter{
			once: &sync.Once{},
			apis: &Server{},
		}

		r.New()
		r.AddRoutes()

		return r, nil
	}
	return nil, fmt.Errorf("Router Type unsupported: %s", routerType)
}
