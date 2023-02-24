package svr

import (
	"github.com/gin-gonic/gin"
	"local/middlewares"
	"log"
)

func Start(addr string) {
	r := gin.Default()
	r.Use(middlewares.Cors())
	ping(r)
	markCoins(r)
	if err := r.Run(addr); err != nil {
		log.Fatalln("Error:", err)
	}
}
