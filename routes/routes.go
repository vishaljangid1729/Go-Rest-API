package routes

import (
	m "rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.New()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true
	router.Use(m.GenerateUUID(), m.RequestLogger(), gin.Recovery())

	// Add routes here
	addUserRoute(router.Group("/users"))

	return router

}
