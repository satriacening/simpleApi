package user

import (
	_entities "simpleApi/entities"
)

type UserUseCaseInterface interface {
	RegisterUser(user _entities.User) error
	EditUser(user _entities.User) error
	GetByName(name string) (_entities.User, error)
	ReferalCode(id int, code string) (string, error)
}
