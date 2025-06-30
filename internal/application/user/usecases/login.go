package usecases

import (
	"context"
	"errors"
	"github.com/agungputrap/linkvault-link-api/internal/application/user/dto"
	"github.com/agungputrap/linkvault-link-api/internal/domain/user"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type LoginUseCase struct {
	repo user.Repository
}

func NewLoginUseCase(repo user.Repository) *LoginUseCase {
	return &LoginUseCase{repo}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req dto.LoginRequest) (*dto.UserResponse, error) {
	email, err := user.NewEmail(req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	found, err := uc.repo.FindByEmail(email.String())
	if err != nil || found.Password.Compare(req.Password) {
		return nil, ErrInvalidCredentials
	}

	return &dto.UserResponse{
		ID:    found.ID,
		Name:  found.Name,
		Email: found.Email.String(),
	}, nil
}
