package controller

import (
	"main/model"
	"main/service"
	"net/http"

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
