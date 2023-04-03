package main

import (
	"main/controller"
)

func main() {
	router := controller.StartRouter()
	router.Run("localhost:8080")
}
