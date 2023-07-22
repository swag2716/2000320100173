package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swapnika/train_data/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/trains", controllers.GetTrainsHanlder).Methods("GET")

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
