package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var SnakeInfo = map[string]string{
	"apiversion": "1",
	"author":     "n8henrie",
	"color":      "#00FF00",
	"head":       "000000",
	"tail":       "FF0000",
	"version":    "0.0.1-alpha",
}

type Position struct {
	X, Y int
}

type Data struct {
	Game  Game
	Turn  int
	Board Board
	You   Snake
}

type Board struct {
	Height, Width int
	Food, Hazards []Position
	Snakes        []Snake
}

type Snake struct {
	Id, Name, Shout, Squad, Latency string
	Health, Length                  int
	Body                            []Position
	Head                            Position
}

func writeAsJson(w http.ResponseWriter, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var d Data
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	move := map[string]string{
		"move":  "down",
		"shout": "Get down!",
	}
	writeAsJson(w, move)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	writeAsJson(w, SnakeInfo)
}

type Game struct {
	id      string
	ruleset Ruleset
	timeout int
}

type Ruleset struct {
	name, version string
}

func startEndHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var d Data
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// No response necessary
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/move", moveHandler)
	mux.HandleFunc("/start", startEndHandler)
	mux.HandleFunc("/end", startEndHandler)

	log.Fatal(http.ListenAndServe(":9433", mux))
}
