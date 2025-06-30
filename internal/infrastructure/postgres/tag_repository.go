package postgres

import (
	"github.com/agungputrap/linkvault-link-api/internal/domain/tag"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) tag.Repository {
	return &tagRepository{db}
}

func (r *tagRepository) ListTagsByUser(userID uint) ([]string, error) {
	var tags []string
	query := `
			select distinct unnest(tags)
			from links
			where user_id = ?
	`
	err := r.db.Raw(query, userID).Scan(&tags).Error
	return tags, err
}

func (r *tagRepository) FindLinksByTag(userID uint, tag string) ([]uint, error) {
	var linkIDs []uint
	query := `
		select id from links
		where user_id = ? and ? = any(tags)
	`
	err := r.db.Raw(query, userID, tag).Scan(&linkIDs).Error
	return linkIDs, err
}
