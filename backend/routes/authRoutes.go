package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/tech-rounak/true-beacon/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.GET("/users/profile", controller.Profile())
}
