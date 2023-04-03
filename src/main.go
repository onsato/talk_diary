package main

import (
	"src/controller"
)

func main() {
	router := controller.StartRouter()
	router.Run("localhost:8080")
}
