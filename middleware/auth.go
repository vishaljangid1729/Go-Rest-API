package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isAuthenticated(c) {
			c.Redirect(302, "/login")
			c.Abort()
		}
	}
}

func isAuthenticated(c *gin.Context) bool {
	// Check if user is authenticated
	return true
}
