package search

import (
	_entities "simpleApi/entities"
	_searchRepository "simpleApi/repository/search"
)

type SearchUseCase struct {
	searchRepository _searchRepository.SearchRepositoryInterface
}

func NewSearchUseCase(searchRepo _searchRepository.SearchRepositoryInterface) SearchUseCaseInterface {
	return &SearchUseCase{
		searchRepository: searchRepo,
	}
}

func (auc *SearchUseCase) Search(input string) (_entities.Data, error) {
	hero, err := auc.searchRepository.Search(input)
	return hero, err
}
