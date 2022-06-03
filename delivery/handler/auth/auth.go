package auth

import (
	"fmt"
	"net/http"
	"simpleApi/delivery/helper"
	_authUseCase "simpleApi/usecase/auth"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUseCase _authUseCase.AuthUseCaseInterface
}

func NewAuthHandler(auth _authUseCase.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: auth,
	}
}

func (ah *AuthHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		type loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var login loginData
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		token, errorLogin := ah.authUseCase.Login(login.Username, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success login", responseToken))
	}
}
