package storage

import (
	"database/sql"
	"log"

	"github.com/ErayCep/ryg/internal/model"
)

func (s *Storage) GetReviews() (model.Reviews, error) {
	rows, err := s.DB.Query("SELECT * FROM reviews")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	reviews := make([]*model.Review, 0)

	for rows.Next() {
		review := new(model.Review)
		err := rows.Scan(&review.Review_ID, &review.Game_ID, &review.Rating, &review.Description, &review.Base.CreatedAt, &review.Base.UpdatedAt)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (s *Storage) GetReview(id int) (*model.Review, error) {
	row := s.DB.QueryRow("SELECT * FROM reviews WHERE review_id = $1", id)

	newReview := new(model.Review)
	err := row.Scan(&newReview.Review_ID, &newReview.Game_ID, &newReview.Rating, &newReview.Description, &newReview.Base.CreatedAt, &newReview.CreatedAt)
	if err == sql.ErrNoRows {
		log.Printf("[SQL] Game with given ID could not found")
		return nil, err
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	return newReview, nil
}

func (s *Storage) AddReview(review *model.Review) (sql.Result, error) {
	res, err := s.DB.Exec("INSERT INTO reviews (review_id, game_id, rating, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", &review.Review_ID, &review.Game_ID, review.Rating, review.Description, review.CreatedAt, review.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return res, err
	}

	return res, nil
}
