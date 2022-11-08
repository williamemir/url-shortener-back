package routes

import (
	"github.com/gin-gonic/gin"
	"url-shortener-back/controllers"
)

func URLShortRoute(router *gin.Engine)  {
	//All routes related to users comes here
	router.POST("/short-url", controllers.CreateURLShort())
	router.GET("/urls", controllers.GetAllURLs())
	router.GET("/:urlId", controllers.GetURLFull())
}
