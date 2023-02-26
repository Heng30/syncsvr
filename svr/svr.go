package svr

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"local/db"
	"local/middlewares"
	"log"
	"net/http"
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

func getComFactory(token, name string) func(*gin.Context) {
	return func(c *gin.Context) {
		table := c.Param(token)
		if tokens, err := db.Query(table, name); err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "data": tokens})
		} else {
			c.JSON(http.StatusNotFound, errorBody(err))
		}
	}
}

func postComFactory(token, name string) func(*gin.Context) {
	return func(c *gin.Context) {
		table := c.Param(token)
		tokens := make([]byte, 0, 4096)
		data := make([]byte, 1024)
		for {
			n, err := c.Request.Body.Read(data)
			if n > 0 {
				tokens = append(tokens, data[:n]...)
			}

			if err != nil {
				if err != io.EOF {
					c.JSON(http.StatusBadRequest, errorBody(err))
					return
				} else {
					break
				}
			}
			if n <= 0 {
				break
			}
		}

		if len(tokens) <= 0 {
			c.JSON(http.StatusBadRequest, errorBody(errors.New("Not support empty body")))
			return
		}

		if err := db.Update(table, name, string(tokens)); err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0})
		} else {
			c.JSON(http.StatusNotFound, errorBody(err))
		}
	}
}
