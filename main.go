package main

import (
	"github.com/gin-gonic/gin"
	"url-shortener-back/configs"
	"url-shortener-back/routes"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	routes.URLShortRoute(router)

	configs.ConnectDB()

	/*for i := 1; i < 10; i++ {
		fmt.Println(utils.RandomString(10))
	}*/

	router.Run(":8080")
}