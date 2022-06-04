package search

import _entities "simpleApi/entities"

type SearchRepositoryInterface interface {
	Search(input string) (_entities.Data, error)
}
