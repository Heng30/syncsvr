package svr

import (
	"github.com/gin-gonic/gin"
)

func markCoins(r gin.IRouter) {
	r.GET("/:token/markCoins", getComFactory("token", "markCoins"))
	r.POST("/:token/markCoins", postComFactory("token", "markCoins"))
}
