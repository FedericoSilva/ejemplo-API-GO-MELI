package services

import (
	"encoding/json"
	"github.com/mercadolibre/ejemplo-API-go/src/api/domain"
	"github.com/mercadolibre/ejemplo-API-go/src/api/utils"
	"io/ioutil"
	"net/http"
)

func GetCategories (siteId string) (*domain.Category,*utils.ApiErrors){

	var categories domain.Category

	res, err := http.Get("https://api.mercadolibre.com/sites/" + siteId + "/categories")

	if err != nil {
		return nil, &utils.ApiErrors{
			err.Error(),
			http.StatusBadRequest,
		}
	}

	data, error := ioutil.ReadAll(res.Body)

	if error != nil {
		return nil, &utils.ApiErrors{
			error.Error(),
			http.StatusBadRequest,
		}
	}

	if er := json.Unmarshal([] byte(data), &categories) ; er != nil {
		return nil, &utils.ApiErrors{
			er.Error(),
			http.StatusBadRequest,
		}
	}

	return &categories, nil
}