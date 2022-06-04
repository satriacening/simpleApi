package search

import _entities "simpleApi/entities"

type SearchUseCaseInterface interface {
	Search(input string) (_entities.Data, error)
}
