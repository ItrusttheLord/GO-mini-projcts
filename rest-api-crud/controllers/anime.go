package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Anime struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Genre    string  `json:"genre"`
	Episodes string  `json:"episodes"`
	Rating   string  `json:"rating"`
	Studio   *Studio `json:"studio"`
	Author   *Author `json:"author"`
}

type Studio struct {
	Studio string `json:"studio"`
}

type Author struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthdate"`
}

var animes = []Anime{
	{ID: "1", Title: "Naruto", Genre: "Shonen", Episodes: "720", Rating: "8.7/10", Studio: &Studio{Studio: "Pierrot"}, Author: &Author{Name: "Misashi Kishimoto", BirthDate: "November 8,1974"}},
	{ID: "2", Title: "Bleach", Genre: "Shonen", Episodes: "366", Rating: "8.7/10", Studio: &Studio{Studio: "Pierrot"}, Author: &Author{Name: "Tite Kubo", BirthDate: "June 26,1977"}},
	{ID: "3", Title: "One Piece", Genre: "Shonen", Episodes: "1120", Rating: "8.7/10", Studio: &Studio{Studio: "Pierrot"}, Author: &Author{Name: "Eiichiro Oda", BirthDate: "January 1,1975"}},
}

func CreateAnime(w http.ResponseWriter, r *http.Request) {
	var anime Anime
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&anime); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	animeID := strconv.Itoa(rand.Intn(999999999))
	anime.ID = animeID
	animes = append(animes, anime)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(anime) // Return only the created anime
}

func GetAnimes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(animes) //encode the animes
}

func GetAnimeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //extract params from the request
	// loop through animes slice
	for _, ID := range animes {
		if ID.ID == params["id"] {
			json.NewEncoder(w).Encode(ID)
			return
		}
	}
	w.WriteHeader(404)
}

func UpdateAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedAnime Anime

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&updatedAnime); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the anime by ID
	for i, ID := range animes {
		if ID.ID == params["id"] {
			updatedAnime.ID = ID.ID  // preserve the original ID
			animes[i] = updatedAnime // update the anime slice
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(updatedAnime) // send response
			return
		}
	}
	w.WriteHeader(404) // Not found
}

func DeleteAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, ID := range animes {
		if ID.ID == params["id"] {
			animes = append(animes[:i], animes[i+1:]...) //deletes the elm
			w.WriteHeader(204)
			return
		}
	}
	w.WriteHeader(404)
}
