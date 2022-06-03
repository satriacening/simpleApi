package routes

import (
	_authHandler "simpleApi/delivery/handler/auth"
	_userHandler "simpleApi/delivery/handler/user"
	_middlewares "simpleApi/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/user", uh.RegisterUserHandler())
	e.PUT("/user/me", uh.EditUserHandler(), _middlewares.JWTMiddleware())
	e.GET("/user/:name", uh.GetByNameHandler())
}
