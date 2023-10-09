package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["message"] = "Hello!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := make(map[string]string)
	err := decoder.Decode(&data)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		data["message"] = "error"
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// this is for deploying on Heroku
	port := os.Getenv("port")
	if port == "" {
		port = "8000"
	}
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":"+port, nil)
}
