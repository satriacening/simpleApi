package main

import (
	"fmt"
	"log"
	"simpleApi/configs"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	_authHandler "simpleApi/delivery/handler/auth"
	_authRepository "simpleApi/repository/auth"
	_authUseCase "simpleApi/usecase/auth"

	_userHandler "simpleApi/delivery/handler/user"
	_userRepository "simpleApi/repository/user"
	_userUseCase "simpleApi/usecase/user"

	_routes "simpleApi/delivery/routes"
	_utils "simpleApi/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterUserPath(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
