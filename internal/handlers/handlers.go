package handlers

import (
	"log"

	"github.com/ErayCep/ryg/internal/storage"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	Storage storage.Storage
	l       *log.Logger
}

func NewHandler(l *log.Logger, storage storage.Storage) *Handlers {
	return &Handlers{
		l:       l,
		Storage: storage,
	}
}

var Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
