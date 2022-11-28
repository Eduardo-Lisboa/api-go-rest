package routes

import (
	"log"
	"net/http"

	"github.com/Eduardo-Lisboa/go-res-api/controllers"
	"github.com/Eduardo-Lisboa/go-res-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.OnePersonality).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":9090", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
