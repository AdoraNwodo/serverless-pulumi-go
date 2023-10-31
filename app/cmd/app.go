package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		customHandlerPort = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/date", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{ "now": %d }`, time.Now().UnixNano()/1000000)))
	})

	mux.HandleFunc("/api/weather", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		weather := WeatherData{
			City:        "New York",
			Temperature: 72.5,
			Condition:   "Partly Cloudy",
		}

		// Convert the weather data to JSON.
		responseJSON, err := json.Marshal(weather)
		if err != nil {
			http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", customHandlerPort), mux))
}

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}
