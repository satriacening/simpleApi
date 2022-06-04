package user

import (
	"fmt"
	"net/http"
	"simpleApi/delivery/helper"
	_middlewares "simpleApi/delivery/middlewares"
	_entities "simpleApi/entities"
	_userUseCase "simpleApi/usecase/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(userUseCase _userUseCase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (uh *UserHandler) RegisterUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		tx := c.Bind(&user)
		err := uh.userUseCase.RegisterUser(user)
		if tx != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(tx.Error()))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succses to register"))

	}
}

func (uh *UserHandler) EditUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		var editUser _entities.User
		editUser.ID = uint(idToken)
		errBind := c.Bind(&editUser)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error()))
		}
		err := uh.userUseCase.EditUser(editUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to update cafe"))
	}
}

func (uh *UserHandler) GetByNameHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("name")
		user, err := uh.userUseCase.GetByName(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses get user", user))
	}
}
func (uh *UserHandler) ReferalHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		c.Bind(&user)
		id, _ := _middlewares.ExtractToken(c)
		code := user.ReferalCode
		fmt.Println("code di handler : ", user.ReferalCode)
		refer, err := uh.userUseCase.ReferalCode(id, code)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses reference from %s", refer))
	}
}
