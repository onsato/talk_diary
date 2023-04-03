package controller

import (
	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*html")
	r.GET("/", ShowHome)
	r.GET("/list", ShowList)

	return r
}
