package search

import (
	"net/http"
	"simpleApi/delivery/helper"
	_entities "simpleApi/entities"
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
		var hero _entities.Data
		input := c.Param("name")

		hero, err := sh.searchUseCase.Search(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succses get search", hero))
	}
}
