package usecases

import (
	"context"
	"github.com/agungputrap/linkvault-link-api/internal/application/user/dto"
	"github.com/agungputrap/linkvault-link-api/internal/domain/user"
)

type RegisterUseCase struct {
	repo user.Repository
}

func NewRegisterUseCase(repo user.Repository) *RegisterUseCase {
	return &RegisterUseCase{repo}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error) {
	email, err := user.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	password, err := user.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		Name:     req.Name,
		Email:    email,
		Password: password,
	}
	if err := uc.repo.Create(newUser); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email.String(),
	}, nil
}
