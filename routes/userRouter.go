package routes

import (
	"github.com/Golang-Authentication/controllers"
	"github.com/Golang-Authentication/middleware"
	"github.com/Golang-Authentication/vendor/github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
