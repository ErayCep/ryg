package model

import (
	"encoding/json"
	"io"
	"log"
	"time"
)

type Game struct {
	Game_ID     int       `json:"id"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"releaseDate"`
	Price       float32   `json:"price"`
}

type AddGameRequest struct {
	Title       string    `json:"title" validate:"required"`
	Genre       string    `json:"genre" validate:"required"`
	ReleaseDate time.Time `json:"releaseDate"`
	Price       float32   `json:"price" validate:"gte=0"`
}

type GetGameResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"releaseDate"`
	Price       float32   `json:"price"`
}

type Games []*Game

type KeyGame struct{}

func (g *Games) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(g)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (g *Game) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(g)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (g *Games) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(g)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (g *Game) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(g)
	if err != nil {
		log.Fatal(err)
		log.Printf("Error while decoding game from JSON format")
		return err
	}

	return nil
}
