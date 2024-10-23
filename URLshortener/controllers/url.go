package controllers

import (
	"URLshortener/models"
	"URLshortener/storage"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var URL models.URL
var urlStorage = storage.NewStorage()

// HandlePost
func HandlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&URL)
	if err != nil {
		http.Error(w, "Something when wrong when decoding", http.StatusBadGateway)
		return
	}
	// validate urls
	_, err1 := url.Parse(URL.ShortURL)
	if URL.OriginalURL == "" || URL.ShortURL == "" || err1 != nil {
		http.Error(w, "Error: Check url and make sure is not empty or has invalid format", http.StatusBadRequest)
		return
	}

	randString := generateRandomString(6)
	//store original url using the rand string
	err2 := urlStorage.AddURL(randString, URL.OriginalURL)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusConflict)
	}
	// Create a response with the original and shortened URL
	shortenedURL := "http://short.ly" + randString
	response := map[string]string{
		"original": URL.OriginalURL,
		"short":    shortenedURL,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// helper func to generate random string
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano()) //send the rand num generator
	b := make([]byte, length)        //wi;; store rand  chracts

	for i := range b {
		randIndex := rand.Intn(len(charset)) //rand ind within len of charset
		b[i] = charset[randIndex]            // add rand charct to byte sl
	}
	return string(b) //convert sl to a string
}

// HandleGet
func HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	short := r.URL.Query().Get("short") //extract short url
	if short == "" {
		http.Error(w, "URL is empty", http.StatusNotFound)
		return
	} //retrueve original url corresponding to the shor url
	url, ok := urlStorage.GetURL(short) //GetURL returns a boolean
	if !ok {
		http.Error(w, "short URL does not exist", http.StatusNotFound)
		return
	}
	response := map[string]string{
		"original": url,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleDelete
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	short := r.URL.Query().Get("short")
	if short == "" {
		http.Error(w, "valid URL required!", http.StatusBadRequest)
		return
	}
	err := urlStorage.RemoveURL(short)
	if err != nil {
		http.Error(w, "Sorry, URL could not be found!", http.StatusNotFound)
		return
	}
	response := map[string]string{
		"message": "URL successfully removed!",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Get all the urls
func HandleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := urlStorage.ListAllURLs()
	if len(list) == 0 {
		http.Error(w, "no URLs to display", http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}
