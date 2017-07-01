package api

import (
	"fmt"
 	"net/http"
 	//"github.com/gorilla/mux"
	)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "Hello World!")
}