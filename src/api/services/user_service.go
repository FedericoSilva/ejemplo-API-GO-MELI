package services

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/ejemplo-API-go/src/api/domain"
	"github.com/mercadolibre/ejemplo-API-go/src/api/utils"
	"io/ioutil"
	"net/http"
)

func GetUser (userID int) (*domain.User,*utils.ApiErrors){

	var user domain.User

	res, err := http.Get("https://api.mercadolibre.com/users/"+string(userID))

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

	if er := json.Unmarshal([] byte(data), &user) ; er != nil {
		return nil, &utils.ApiErrors{
			er.Error(),
			http.StatusBadRequest,
		}
	}

	fmt.Println(&user)
	return &user, nil
}
