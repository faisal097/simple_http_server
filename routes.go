package simple_http_server

import (
	"fmt"
	"sync"
)

type Router interface {
	New()
	Init()
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
		r.Init()

		return r, nil
	}
	return nil, fmt.Errorf("Router Type unsupported: %s\n", routerType)
}