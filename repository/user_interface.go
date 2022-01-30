package repository

import "diary-app-service/entity"

type UserRepositoryInterface interface {
	GetUser(username string, user *entity.User) error
	CreateUser(user *entity.User) error
}
