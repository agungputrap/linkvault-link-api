package usecases

import (
	"context"
	"github.com/agungputrap/linkvault-link-api/internal/application/link/dto"
	"github.com/agungputrap/linkvault-link-api/internal/domain/link"
)

type CreateLinkUseCase struct {
	repo link.Repository
}

func NewCreateLinkUseCase(repo link.Repository) *CreateLinkUseCase {
	return &CreateLinkUseCase{repo}
}

func (uc *CreateLinkUseCase) Execute(ctx context.Context, userID uint, req dto.CreateLinkRequest) (*dto.LinkResponse, error) {
	newLink := &link.Link{
		UserID:      userID,
		Title:       req.Title,
		Url:         req.Url,
		Description: req.Description,
		Tags:        req.Tags,
	}
	if err := uc.repo.Create(newLink); err != nil {
		return nil, err
	}
	return &dto.LinkResponse{
		ID:          newLink.ID,
		Title:       newLink.Title,
		Url:         newLink.Url,
		Description: newLink.Description,
		Tags:        newLink.Tags,
	}, nil
}
