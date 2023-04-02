package main

import (
	"src/controller"
	"src/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	dairyService    service.DairyService       = service.New()
	dairyController controller.DairyController = controller.New(dairyService)
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("template/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/list", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "list.html", gin.H{
			"id":    1,
			"title": "sample",
		})
	})

	router.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, dairyController.Show(ctx))
	})

	router.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, dairyController.Save(ctx))
	})

	router.Run("localhost:8080")

}
