package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/ejemplo-API-go/src/api/controllers"
)

const (
	port = ":8080"
)

var(
	router = gin.Default()
)

func main() {

	router.GET("/user/:id",controllers.GetDataSite)
	router.Run(port)

}