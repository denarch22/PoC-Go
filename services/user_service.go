package services

import (
	"sistema-usuarios-go/config"
	"sistema-usuarios-go/models"
)

func CreateUserService(user *models.User) error {
	db := config.ConnectDatabase()
	return db.Create(user).Error
}

func GetUsersService() ([]models.User, error) {
	db := config.ConnectDatabase()
	var users []models.User
	err := db.Find(&users).Error
	return users, err
}
