package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered panic:", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}
