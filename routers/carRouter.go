package routers

import (
	"github.com/gin-gonic/gin"
	"introduction-gin/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	return router
}
