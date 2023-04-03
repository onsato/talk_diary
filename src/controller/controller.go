package controller

import (
	"net/http"
	"src/model"
	"src/service"

	"github.com/gin-gonic/gin"
)

type DairyController interface {
	Save(c *gin.Context) model.Diary
	Show(c *gin.Context) model.Diary
}

type controller struct {
	service service.DairyService
}

func ShowHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ShowList(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{
		"id":    1,
		"title": "sample",
	})
}

func CreateDairy(c *gin.Context) {

}

func SaveDiary(c *gin.Context) {

}

// func CreateNew(service service.DairyService) DairyController {
// 	return &controller{
// 		service: service,
// 	}
// }

// func (ctr *controller) Save(c *gin.Context) model.Diary {
// 	var dairy model.Diary
// 	c.BindJSON(&dairy)
// 	ctr.service.Save(dairy)
// 	return dairy
// }

// func (ctr *controller) Show(c *gin.Context) model.Diary {
// 	var dairy model.Diary
// 	c.BindJSON(&dairy)
// 	ctr.service.Show(dairy)
// 	return dairy
// }
