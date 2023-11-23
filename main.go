package main

import (
	"github.com/gorilla/mux"
	"gitub.com/EltajGafarli/go-movies-crud-nondba/controllers"
	"log"
	"net/http"
)

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	route.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	log.Println("Server run successfully on port :8000")
	if err := http.ListenAndServe(":8000", route); err != nil {
		log.Fatal(err.Error())
	}

}
