package user

import (
	"errors"
	_entities "simpleApi/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) RegisterUser(user _entities.User) error {
	tx := ur.database.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ur *UserRepository) EditUser(user _entities.User) error {
	userSave := user
	tx := ur.database.Save(&userSave)
	if tx != nil {
		return tx.Error
	}
	return nil
}
func (ur *UserRepository) GetByName(name string) (_entities.User, error) {
	var user _entities.User

	tx := ur.database.Where("name = ?", name).Find(&user)
	if tx.RowsAffected == 0 {
		return _entities.User{}, errors.New("user not found")
	}
	return user, nil
}

func (ur *UserRepository) ReferalCode(code string) (string, error) {
	var user _entities.User
	var update _entities.User
	tx := ur.database.Where("referal_code = ?", code).Find(&user)
	if tx.RowsAffected == 0 {
		return "", errors.New("invalid code")
	}
	update.Reference = user.Name
	save := ur.database.Save(&update)
	if save.Error != nil {
		return "", errors.New("failed to insert referal code")
	}
	return user.Name, nil
}
