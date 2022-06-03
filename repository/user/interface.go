package user

import (
	_entities "simpleApi/entities"
)

type UserRepositoryInterface interface {
	RegisterUser(user _entities.User) error
	EditUser(user _entities.User) error
	GetByName(name string) (_entities.User, error)
	ReferalCode(code string) (string, error)
}