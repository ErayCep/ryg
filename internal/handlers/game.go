package handlers

import (
	"net/http"
	"strconv"

	"github.com/ErayCep/ryg/internal/model"
	"github.com/gorilla/mux"
)

func (h *Handlers) GetGamesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h.l.Printf("Handle GET Games")
	var games model.Games

	result, err := h.Storage.DB.Query("SELECT * FROM games")
	if err != nil {
		http.Error(w, "Problem querying games from database", http.StatusInternalServerError)
		h.l.Fatal(err)
	}

	defer result.Close()

	for result.Next() {
		var game model.Game
		err := result.Scan(&game.Game_ID, &game.Title, &game.Genre, &game.ReleaseDate, &game.Price)
		if err != nil {
			http.Error(w, "Error while writing values from database", http.StatusInternalServerError)
			h.l.Fatal(err)
		}

		games = append(games, &game)
	}

	games.ToJSON(w)
}

func (h *Handlers) GetGameHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Printf("Handle GET Game")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error extracting id from URL", http.StatusBadRequest)
		return
	}
	game, err := h.Storage.GetGame(id)

	if err != nil {
		http.Error(w, "Error getting game from database", http.StatusInternalServerError)
		return
	}

	game.ToJSON(w)
}

func (h *Handlers) PostGamesHandler(w http.ResponseWriter, r *http.Request) {
	h.l.Printf("Handle POST Game")

	game := model.Game{}

	err := game.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Failed to decode json", http.StatusBadRequest)
	}

	h.Storage.AddGame(&game)
}

func (h *Handlers) PutGamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error extracting id from URL", http.StatusBadRequest)
		return
	}

	game := model.Game{}

	err = game.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Failed to decode json", http.StatusBadRequest)
		return
	}

	h.Storage.UpdateGame(&game, id)
}

func (h *Handlers) DeleteGamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error extracting id from URL", http.StatusBadRequest)
		return
	}

	h.Storage.DeleteGame(id)
}
