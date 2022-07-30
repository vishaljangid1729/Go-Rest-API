package middleware

import (
	"os"
	"time"

	log "rest-api/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogger() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		stop := time.Since(start)

		uuid := c.GetString("uuid")
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		clientReferer := c.Request.Referer()
		requestBody := c.Request.Body
		requestPath := c.Request.URL.Path

		statusCode := c.Writer.Status()
		dataLength := c.Writer.Size()

		if dataLength < 0 {
			dataLength = 0
		}
		log.Info(c.Request.Context(), "Request",
			zap.String("uuid", uuid),
			zap.String("method", c.Request.Method),
			zap.String("path", requestPath),
			zap.String("timestamp", start.Format(time.RFC1123)),
			zap.Int("statusCode", statusCode),
			zap.String("hostname", hostname),
			zap.String("latency", stop.String()),
			zap.String("clientIP", clientIP),
			zap.String("referer", clientReferer),
			zap.Int("dataLength", dataLength),
			zap.String("userAgent", clientUserAgent),
			zap.Any("request", requestBody),
		)
	}
}
