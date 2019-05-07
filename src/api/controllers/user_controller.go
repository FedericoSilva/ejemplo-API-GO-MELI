package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/ejemplo-API-go/src/api/domain"
	"github.com/mercadolibre/ejemplo-API-go/src/api/services"
	"github.com/mercadolibre/ejemplo-API-go/src/api/utils"
	UserService "github.com/mercadolibre/myml/src/api/services/user"
	"net/http"
	"strconv"
	"sync"
)

type Cross struct {
	*domain.Site
	*domain.Category
}

const paramUserId  = "id"

func GetDataSite (context *gin.Context) {

	id := context.Param(paramUserId)

	userId, err := strconv.ParseInt(id,10,64)

	if err != nil {
		apiError := &utils.ApiErrors{
			"No se puede parsear",
			http.StatusBadRequest,
		}
		context.JSON(apiError.Status,apiError)
		return
	}

	user, error := UserService.GetUserApi(userId)

	if error != nil {
		apiError := &utils.ApiErrors{
			error.Message,
			error.Status,
		}
		context.JSON(apiError.Status,apiError)
		return
	}

	cross, e := loadDataChannel(user.SiteID)

	if e != nil {
		apiError := &utils.ApiErrors{
			e.Message,
			e.Status,
		}
		context.JSON(apiError.Status,apiError)
		return
	}

	context.JSON(200,&cross)
}

func loadDataChannel (siteId string) (*Cross, *utils.ApiErrors) {

	var wg sync.WaitGroup
	crossChan := make(chan *Cross,2)
	var cross Cross

	wg.Add(2)
	go func() {
		defer wg.Done()
		cat, err := services.GetCategories(siteId)
		var crossT Cross

		if err != nil{
			return
		}

		crossT.Category = cat
		crossChan <- &crossT
	}()

	go func() {
		defer wg.Done()
		site, err := services.GetSite(siteId)
		var crossT Cross

		if err != nil{
			return
		}

		crossT.Site = site
		crossChan <- &crossT
	}()

	for i := 0 ; i < 2 ; i++  {

		select {

		case r:= <- crossChan:

			if(r.Site != nil){
				cross.Site = r.Site
				continue
			}

			if(r.Category !=  nil){
				cross.Category = r.Category
				continue
			}
		}
	}

	wg.Wait()
	return &cross, nil
}