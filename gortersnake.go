package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Position struct {
	X, Y int
}

// https://webcache.googleusercontent.com/search?q=cache:WjZ1-ek8jEwJ:https://docs.battlesnake.com/references/api/sample-move-request+&cd=2&hl=en&ct=clnk&gl=us
type moveData struct {
	Game struct {
		Id      string
		Ruleset struct {
			Name, Version string
		}
		Timeout int
	}
	Turn  int
	Board struct {
		Height, Width int
		Food, Hazards []Position
		Snakes        []Snake
	}
	You Snake
}

type Snake struct {
	Id, Name, Shout, Squad, Latency  string
	Health, Length int
	Body                    []Position
	Head                    Position
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	var m moveData
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// func handler(w http.ResponseWriter, r *http.Request) { }

func main() {
	mux := http.NewServeMux()
	// server := NewTaskServer()
	mux.HandleFunc("/move", moveHandler)

	log.Fatal(http.ListenAndServe(":9433", mux))
}
