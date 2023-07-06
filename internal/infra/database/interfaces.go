package database

import "github.com/ThiagoLeite06/goexpert/api-9/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
