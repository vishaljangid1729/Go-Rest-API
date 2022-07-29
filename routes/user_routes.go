package routes

import (
	"rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func addUserRoute(rg *gin.RouterGroup) {
	rg.GET("/", handlers.GetUser)
}
