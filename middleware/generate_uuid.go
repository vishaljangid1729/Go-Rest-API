package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func GenerateUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		UUID, _ := uuid.NewV4()
		uuid := UUID.String()
		c.Set("uuid", uuid)
	}
}
