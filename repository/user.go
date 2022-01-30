package repository

import (
	"diary-app-service/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) GetUser(username string, user *entity.User) error {

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (r UserRepository) CreateUser(user *entity.User) error {

	r.db.Create(&user)

	return nil
}
