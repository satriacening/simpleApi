package auth

import (
	"errors"
	_helper "simpleApi/delivery/helper"
	_middlewares "simpleApi/delivery/middlewares"
	_entities "simpleApi/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) Login(username string, password string) (string, error) {
	var user _entities.User
	tx := ar.database.Where("username = ?", username).Find(&user)
	check := _helper.CheckPassHash(password, user.Password)
	if tx.Error != nil {
		return "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return "user not found", errors.New("username not found")
	}
	if !check {
		return "", errors.New("incorect password")
	}

	token, err := _middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return "create token failed", err
	}

	return token, nil

}
