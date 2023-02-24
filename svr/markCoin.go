package svr

import "github.com/gin-gonic/gin"
import "net/http"
import "log"

func markCoins(r gin.IRouter) {
	r.GET("/:token/markCoins", func(c *gin.Context) {
		token := c.Param("token")
		log.Println("token:", token)

		data := map[string]interface{}{
			"data": []string{"BTC", "ETH"},
		}
		c.JSON(http.StatusOK, data)
	})
}
