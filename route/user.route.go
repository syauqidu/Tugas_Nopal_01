package route

import (
	"Tugas/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRoute(route *gin.Engine) {
	route.GET("/user", controller.GetUser)
	route.POST("/tambah", controller.TambahAngka)
	route.GET("/lingkaran/:jari", controller.HitungLuasLingkaran)
}
