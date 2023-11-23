package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitub.com/EltajGafarli/go-movies-crud-nondba/models"
	"net/http"
	"strconv"
)

var movies []models.Movie = []models.Movie{}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}

}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestedId, _ := strconv.Atoi(params["id"])
	for _, item := range movies {

		if item.ID == requestedId {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			break
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = movies[len(movies)-1].ID + 1
	movies = append(movies, movie)
	err = json.NewEncoder(w).Encode("Created successfully")
	if err != nil {
		return
	}
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestedId, _ := strconv.Atoi(params["id"])
	var currentMovie models.Movie

	err := json.NewDecoder(r.Body).Decode(&currentMovie)
	if err != nil {
		return
	}

	for index, movie := range movies {
		if movie.ID == requestedId {
			if currentMovie.Title != "" {
				movie.Title = currentMovie.Title
			}

			if currentMovie.Director != nil {
				movie.Director = currentMovie.Director
			}

			if currentMovie.Isbn != "" {
				movie.Isbn = currentMovie.Isbn
			}

			movies[index] = movie
			break
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	requestedId, _ := strconv.Atoi(params["id"])
	for index, item := range movies {

		if item.ID == requestedId {
			movies = append(movies[:index], movies[index+1:]...)
			err := json.NewEncoder(w).Encode("Deleted successfully")
			if err != nil {
				return
			}
			break
		}

	}

}
