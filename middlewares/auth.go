package middlewares

import (
	"github.com/gin-gonic/gin"
	"local/db"
	"net/http"
	"strings"
)

func Auth(testMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == "GET" || method == "POST" {
			paths := strings.SplitN(c.Request.RequestURI, "/", 3)
			if len(paths) >= 2 {
				token := paths[1]
				if token != "ping" {
					if !canAccess(token, testMode) {
						c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
					}
				}
			}
		}
		c.Next()
	}
}

func canAccess(token string, testMode bool) bool {
	if !testMode && token == "testToken" {
		return false
	}
	if tokens, err := db.QueryAccessTokens(); err == nil {
		for _, item := range tokens {
			if item == token {
				return true
			}
		}
	}
	return false
}
