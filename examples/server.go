package main

import(
	"log"
	hs "github.com/faisal097/simple_http_server"
)

func main() {
	defer func() {
		log.Println("Exiting http server!")
	}()
	r, err := hs.NewRouter("mux")
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println(r.Start("127.0.0.1:8000"))
}