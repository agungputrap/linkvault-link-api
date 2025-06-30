package postgres

import (
	"github.com/agungputrap/linkvault-link-api/internal/domain/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	var u user.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) FindByID(id uint) (*user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
