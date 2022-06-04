package search

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	_entities "simpleApi/entities"
	"strings"

	"gorm.io/gorm"
)

type SearchRepository struct {
	database *gorm.DB
}

func NewSearchRepository(db *gorm.DB) *SearchRepository {
	return &SearchRepository{
		database: db,
	}
}

func (ar *SearchRepository) Search(input string) (_entities.Data, error) {
	var responseObject _entities.Response
	var hero _entities.Data
	response, err := http.Get("https://ddragon.leagueoflegends.com/cdn/6.24.1/data/en_US/champion.json")

	if err != nil {
		return _entities.Data{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return _entities.Data{}, err
	}

	json.Unmarshal(responseData, &responseObject)

	for k, v := range responseObject.Data {
		if strings.Contains(k, input) {
			return v, nil
		}

		// if k == input {
		// 	return v, nil
		// }
	}

	return hero, errors.New("404")

}
