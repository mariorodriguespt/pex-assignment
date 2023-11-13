package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckIfVisitorAlreadyExists(context *gin.Context) {
	authChunks := strings.Split(context.GetHeader("Authorization"), " ")

	if len(authChunks) != 2 && authChunks[0] != "Bearer" || len(authChunks[1]) == 0 {
		return
	}

	context.Set("visitor-id", authChunks[1])
}

func IdentifyUser(context *gin.Context) {
	authChunks := strings.Split(context.GetHeader("Authorization"), " ")

	if len(authChunks) != 2 && authChunks[0] != "Bearer" || len(authChunks[1]) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization header not as per RFC 7235",
		})
		context.Abort()
		return
	}

	context.Set("visitor-id", authChunks[1])
	context.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
