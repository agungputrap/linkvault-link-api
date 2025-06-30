package usecases

import (
	"context"
	"github.com/agungputrap/linkvault-link-api/internal/domain/tag"
)

type ListUserTagUseCase struct {
	repo tag.Repository
}

func NewListUserTagUseCase(repo tag.Repository) *ListUserTagUseCase {
	return &ListUserTagUseCase{repo}
}

func (uc *ListUserTagUseCase) Execute(context context.Context, userID uint) ([]string, error) {
	return uc.repo.ListTagsByUser(userID)
}
