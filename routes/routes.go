package routes

import (
	"log"
	"net/http"

	"github.com/Eduardo-Lisboa/go-res-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.OnePersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
