package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/goofynugtz/gin-jwt-auth-boilerplate/controllers"
	middlewares "github.com/goofynugtz/gin-jwt-auth-boilerplate/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
