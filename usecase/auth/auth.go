package auth

import (
	_authRepository "simpleApi/repository/auth"
)

type AuthUseCase struct {
	authRepository _authRepository.AuthRepositoryInterface
}

func NewAuthUseCase(authRepo _authRepository.AuthRepositoryInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		authRepository: authRepo,
	}
}

func (auc *AuthUseCase) Login(username string, password string) (string, error) {
	token, err := auc.authRepository.Login(username, password)
	return token, err
}
