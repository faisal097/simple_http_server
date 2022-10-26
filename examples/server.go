package main

import (
	"log"
	"os"

	hs "github.com/faisal097/simple_http_server"
)

var (
	ADDR = "127.0.0.1:8000"
)

func init() {
	if addr := os.Getenv("ADDR"); addr != "" {
		ADDR = addr
	}
}

func main() {
	defer func() {
		log.Println("Exiting http server!")
	}()
	r, err := hs.NewRouter("mux")
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println(r.Start(ADDR))
}
