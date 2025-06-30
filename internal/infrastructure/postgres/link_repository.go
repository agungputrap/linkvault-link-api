package postgres

import (
	"github.com/agungputrap/linkvault-link-api/internal/domain/link"
	"gorm.io/gorm"
)

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) link.Repository {
	return &linkRepository{db}
}

func (r *linkRepository) Create(link *link.Link) error {
	return r.db.Create(link).Error
}

func (r *linkRepository) FindByUser(userID uint) ([]link.Link, error) {
	var links []link.Link

	err := r.db.Where("user_id = ?", userID).Find(&links).Error
	return links, err
}

func (r *linkRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&link.Link{}).Error
}

func (r *linkRepository) Update(link *link.Link) error {
	return r.db.Save(link).Error
}
