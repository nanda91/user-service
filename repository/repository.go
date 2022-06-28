package repository

import (
	"github.com/jinzhu/gorm"
	"user-service/config/env"
	models "user-service/model"
)

type RepositoryInterface interface {
	FindAllUser() ([]models.User, error)
	FindUser(userId string) (models.User, error)
}


type userRepository struct {
	Conn *gorm.DB
	config env.Config
}


func NewLogRepository(conn *gorm.DB, config env.Config) RepositoryInterface {
	return &userRepository{Conn: conn, config: config}
}

func (repo *userRepository) FindAllUser() ([]models.User, error) {
	var users []models.User
	err := repo.Conn.Find(&users).Error

	return users, err
}

func (repo *userRepository) FindUser(userId string) (models.User, error) {
	var user models.User
	err := repo.Conn.Where("\"Userid\" = ?", userId).First(&user).Error

	return user, err
}