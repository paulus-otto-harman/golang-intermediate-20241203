package repository

import (
	"gorm.io/gorm"
	"project/class/domain"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo AuthRepository) Authenticate(user domain.User) (bool, error) {
	var exists bool
	err := repo.db.Model(&domain.User{}).Select("count(*)>0").Where(user).Find(&exists).Error
	return exists, err
}
