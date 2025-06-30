package usecases

import (
	"context"
	"errors"
	"github.com/agungputrap/linkvault-link-api/internal/application/link/dto"
	"github.com/agungputrap/linkvault-link-api/internal/domain/link"
)

var ErrNotFound = errors.New("link not found")

type UpdateLinksUseCase struct {
	repo link.Repository
}

func NewUpdateLinksUseCase(repo link.Repository) *UpdateLinksUseCase {
	return &UpdateLinksUseCase{repo}
}

func (uc *UpdateLinksUseCase) Execute(ctx context.Context, id uint, userID uint, req dto.UpdateLinkRequest) (*dto.LinkResponse, error) {
	links, err := uc.repo.FindByUser(userID)
	if err != nil {
		return nil, err
	}

	var target *link.Link
	for _, l := range links {
		if l.ID == id {
			target = &l
			break
		}
	}

	if target == nil {
		return nil, ErrNotFound
	}

	target.Title = req.Title
	target.Url = req.Url
	target.Description = req.Description
	target.Tags = req.Tags

	if err := uc.repo.Update(target); err != nil {
		return nil, err
	}

	return &dto.LinkResponse{
		ID:          target.ID,
		Title:       target.Title,
		Url:         target.Url,
		Description: target.Description,
		Tags:        target.Tags,
	}, nil
}
