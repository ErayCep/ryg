package storage

import (
	"context"
	"database/sql"

	"github.com/ErayCep/ryg/internal/model"
)

type StorageInterface interface {
	AddGame(ctx context.Context, game model.Game) (int, error)
	GetGame(ctx context.Context, id int) (model.Game, error)
	GetGames(ctx context.Context) (model.Games, error)
	// UpdateGames(ctx context.Context, book model.Game) (int, error)
}

type Storage struct {
	DB *sql.DB
}

func (s *Storage) Close() error {
	if err := s.DB.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetDB() *sql.DB {
	return s.DB
}
