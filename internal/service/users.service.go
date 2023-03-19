package service

import (
	"context"
	"errors"

	"github.com/nahueldev23/inventory-go/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user alrady exist")
	ErrInvalidCredentials = errors.New("credenciales invalidas")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	//todo hash pasword
	return s.repo.SaveUser(ctx, email, name, password)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//todo decrypt password

	if u.Password != password {
		return nil, ErrInvalidCredentials
	}
	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
