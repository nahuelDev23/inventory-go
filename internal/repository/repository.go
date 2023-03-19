package repository

import (
  "context"
  "github.com/nahueldev23/inventory-go/internal/entity"
)

// Repository is the interface that wraps the basic CRUD operations.
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
  SaveUser(ctx context.Context,email,name,password string) error
  GetUserByEmail(ctx context.Context,email string) (entity.User error)
}

