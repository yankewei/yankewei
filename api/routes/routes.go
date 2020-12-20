package routes

import (
	"api/app/http/controllers"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetRoutes() *gin.Engine {
	r = gin.Default()
	authController := controllers.NewAuthController()
	r.POST("/login",authController.Login )
	return r
}