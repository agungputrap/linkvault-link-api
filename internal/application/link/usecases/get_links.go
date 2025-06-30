package usecases

import (
	"context"
	"github.com/agungputrap/linkvault-link-api/internal/application/link/dto"
	"github.com/agungputrap/linkvault-link-api/internal/domain/link"
)

type GetLinksUseCase struct {
	repo link.Repository
}

func NewGetLinksUseCase(repo link.Repository) *GetLinksUseCase {
	return &GetLinksUseCase{repo}
}

func (uc *GetLinksUseCase) Execute(ctx context.Context, userID uint) ([]dto.LinkResponse, error) {
	links, err := uc.repo.FindByUser(userID)
	if err != nil {
		return nil, err
	}
	var result []dto.LinkResponse
	for _, l := range links {
		result = append(result, dto.LinkResponse{
			ID:          l.ID,
			Title:       l.Title,
			Url:         l.Url,
			Description: l.Description,
			Tags:        l.Tags,
		})
	}
	return result, nil
}
