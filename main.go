package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type LocationData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var location LocationData
		if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Print latitude and longitude to console
		fmt.Printf("Latitude: %f, Longitude: %f\n", location.Latitude, location.Longitude)

		// Write latitude and longitude to file
		file, err := os.OpenFile("geo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(fmt.Sprintf("Latitude: %f, Longitude: %f\n", location.Latitude, location.Longitude)); err != nil {
			log.Println("Error writing to file:", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
