package routes

import (
	"github/PraveenAththanayake/Go-JWT/middleware"

	controller "github.com/PraveenAththanayake/Go-JWT/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

