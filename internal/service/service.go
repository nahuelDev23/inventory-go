package service

import (
	"context"

	"github.com/nahueldev23/inventory-go/internal/models"
	"github.com/nahueldev23/inventory-go/internal/repository"
)

// service is the business logic of the aplication
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
  repo repository.Repository
}

func New(repo repository.Repository) Service {
  return &serv{
    repo: repo,
  }
}
