package storage

import (
	"database/sql"
	"log"

	"github.com/ErayCep/ryg/internal/model"
)

func (s *Storage) GetGame(id int) (*model.Game, error) {
	row := s.DB.QueryRow("SELECT * FROM games WHERE game_id = $1", id)

	newGame := new(model.Game)
	err := row.Scan(&newGame.Game_ID, &newGame.Title, &newGame.Genre, &newGame.ReleaseDate, &newGame.Price)
	if err == sql.ErrNoRows {
		log.Printf("[SQL] Game with given ID could not found")
		return nil, err
	} else if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return newGame, nil
}

func (s *Storage) GetGames() (model.Games, error) {
	rows, err := s.DB.Query("SELECT * FROM games")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	games := make([]*model.Game, 0)
	for rows.Next() {
		game := new(model.Game)
		err := rows.Scan(&game.Game_ID, &game.Title, &game.Genre, &game.ReleaseDate, &game.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}

func (s *Storage) AddGame(game *model.Game) (int64, error) {
	var id int64
	err := s.DB.QueryRow("INSERT INTO games (title, genre, releaseDate, price) VALUES ($1, $2, $3, $4) RETURNING id;", game.Title, game.Genre, game.ReleaseDate, game.Price).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	return id, nil
}

func (s *Storage) UpdateGame(game *model.Game, id int) error {
	err := s.DB.QueryRow("UPDATE games SET title = $1, genre = $2, releaseDate = $3, price = $4 WHERE id = $5", game.Title, game.Genre, game.ReleaseDate, game.Price, id).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (s *Storage) DeleteGame(id int) error {
	_, err := s.DB.Exec("DELETE FROM games WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
