package routes

import (
	"github.com/Golang-Authentication/controllers"
	"github.com/Golang-Authentication/vendor/github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
