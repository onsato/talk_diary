package controller

import (
	"src/model"
	"src/service"

	"github.com/gin-gonic/gin"
)

type DairyController interface {
	Save(c *gin.Context) model.Dairies
	Show(c *gin.Context) model.Dairies
}

type controller struct {
	service service.DairyService
}

func New(service service.DairyService) DairyController {
	return &controller{
		service: service,
	}
}

func (ctr *controller) Save(c *gin.Context) model.Dairies {
	var dairy model.Dairies
	c.BindJSON(&dairy)
	ctr.service.Save(dairy)
	return dairy
}

func (ctr *controller) Show(c *gin.Context) model.Dairies {
	var dairy model.Dairies
	c.BindJSON(&dairy)
	ctr.service.Show(dairy)
	return dairy
}
