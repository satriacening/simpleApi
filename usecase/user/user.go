package user

import (
	"errors"
	_helper "simpleApi/delivery/helper"
	_entities "simpleApi/entities"
	_userRepository "simpleApi/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}
func (uuc *UserUseCase) RegisterUser(user _entities.User) error {
	if user.Username == "" {
		return errors.New("username required")
	}
	if user.Name == "" {
		return errors.New("name required")
	}
	if user.Password == "" {
		return errors.New("password required")
	}
	if user.Email == "" {
		return errors.New("email required")
	}
	passHash, errHash := _helper.HashPassword(user.Password)
	if errHash != nil {
		return errors.New("failed hash")
	}

	referal := _helper.RandomStr(user.Username)

	user.ReferalCode = referal
	user.Password = passHash
	err := uuc.userRepository.RegisterUser(user)
	return err
}

func (uuc *UserUseCase) EditUser(user _entities.User) error {
	err := uuc.userRepository.EditUser(user)
	return err
}

func (uuc *UserUseCase) GetByName(name string) (_entities.User, error) {
	user, err := uuc.userRepository.GetByName(name)
	return user, err
}
func (uuc *UserUseCase) ReferalCode(id int, code string) (string, error) {
	referName, err := uuc.userRepository.ReferalCode(id, code)
	return referName, err
}
