package api

import (
	"log"
	"net/http"

	"github.com/ComputePractice2017/weather-server/model"
	"github.com/gorilla/mux"
)

//Run for begin  //

func Run() {
	log.Println("Connect to rethinkDB on localhost")
	err := model.InitSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")
	r.HandleFunc("/Data", getAllWeatherDataHandler).Methods("POST")
	log.Println("Running the server on port 8000...")
	http.ListenAndServe(":8000", r)
}
