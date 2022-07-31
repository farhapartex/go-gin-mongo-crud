package services

import (
	"github.com/farhapartex/go-gin-mongo-crud/models"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
}
