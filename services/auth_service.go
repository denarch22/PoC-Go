package services

import (
	"errors"
	"sistema-usuarios-go/config"
	"sistema-usuarios-go/models"
	"sistema-usuarios-go/utils"
)

func AuthenticateUser(email, password string) (string, error) {
	db := config.ConnectDatabase()
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("usuario no encontrado")
	}
	if user.Password != password {
		return "", errors.New("contraseña inválida")
	}
	return utils.GenerateToken(user.ID, user.Role)
}
