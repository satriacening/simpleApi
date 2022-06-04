package search

import (
	"fmt"
	"net/http"
	"simpleApi/delivery/helper"
	_searchUseCase "simpleApi/usecase/search"

	"github.com/labstack/echo/v4"
)

type SearchHandler struct {
	searchUseCase _searchUseCase.SearchUseCaseInterface
}

func NewSearchHandler(search _searchUseCase.SearchUseCaseInterface) *SearchHandler {

	return &SearchHandler{
		searchUseCase: search,
	}
}

func (sh *SearchHandler) SearchHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// var hero _entities.Data
		input := c.Param("name")
		fmt.Println("gagal konek redis")
		// conn, errCon := redis.Dial("tcp", "localhost:6379")
		// if errCon != nil {
		// 	log.Panic(errCon)
		// }
		fmt.Println("berhasil konek redis")

		// // reply, errGet := redis.Bytes(conn.Do("GET", input))
		// if errGet == nil {
		// 	return c.JSON(http.StatusOK, helper.ResponseSuccess("sucses to get from cache", reply))

		// }
		heroRes, err := sh.searchUseCase.Search(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		// _, err = conn.Do("SET", hero, heroRes)
		// if err != nil {
		// 	log.Panic(err)
		// }

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses get from api", heroRes))
	}
}
