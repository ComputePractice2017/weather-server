package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	)
//
func Run() {
	r:=mux.NewRouter()
	r.HandleFunc("/",helloWorldHandler).Methods("GET")
	log.Println("Running the server on port 8080...")
	http.ListenAndServe(":8080", r)
}
