package svr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"local/middlewares"
	"log"
)

type routerCb func(gin.IRouter)

func Start(addr string, testMode bool) {
	r := gin.Default()

	mcbs := []gin.HandlerFunc{middlewares.Auth(testMode), middlewares.Cors()}
	for _, cb := range mcbs {
		r.Use(cb)
	}

	rcbs := []routerCb{ping, markCoins}
	for _, cb := range rcbs {
		cb(r)
	}

	if err := r.Run(addr); err != nil {
		log.Fatalln("Error:", err)
	}
}

func errorBody(err error) gin.H {
	return gin.H{"code": -1, "error": fmt.Sprintf("%v", err)}
}

