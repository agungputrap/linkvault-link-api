package usecases

import (
	"context"
	"github.com/agungputrap/linkvault-link-api/internal/domain/link"
)

type DeleteLinksUseCase struct {
	repo link.Repository
}

func NewDeleteLinksUseCase(repo link.Repository) *DeleteLinksUseCase {
	return &DeleteLinksUseCase{repo}
}

func (uc *DeleteLinksUseCase) Execute(ctx context.Context, id uint, userID uint) error {
	return uc.repo.Delete(id, userID)
}
