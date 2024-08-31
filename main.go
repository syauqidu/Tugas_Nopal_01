package main

import (
	"Tugas/route"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	route.SetUpRoute(api)
	api.Run(":5112")
}
