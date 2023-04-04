package main

import (
	"main/controller"
	"main/model"
)

func init() {
	model.Migrate()
}

func main() {
	router := controller.StartRouter()
	router.Run("localhost:8080")
}
