package model

import (
	"encoding/json"
	"io"
	"log"
	"time"
)

type Review struct {
	Review_ID   int
	Game_ID     int
	Rating      int
	Description string
	Base
}

type AddReviewRequest struct {
	Username    string `json:"username" validate:"required"`
	Game        `json:"game" validate:"required"`
	Rating      int    `json:"rating" validate:"required"`
	Description string `json:"description"`
}

type GetReviewResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"releaseDate"`
	Price       float32   `json:"price"`
}

func (review *Review) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(review)
	if err != nil {
		log.Printf("Error while encoding review data")
		log.Fatal(err)
		return err
	}

	return nil
}

func (review *Review) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(review)
	if err != nil {
		log.Printf("Error while decoding review data")
		log.Fatal(err)
		return err
	}

	return nil
}

type Reviews []*Review

func (reviews *Reviews) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(reviews)
	if err != nil {
		log.Printf("Error while encoding reviews data")
		log.Fatal(err)
		return err
	}

	return nil
}

func (reviews *Reviews) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(reviews)
	if err != nil {
		log.Printf("Error while decoding reviews data")
		log.Fatal(err)
		return err
	}

	return nil
}
