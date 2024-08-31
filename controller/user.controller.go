package controller

import (
	"Tugas/model"

	"github.com/gin-gonic/gin"

	"strconv"
)

var users = []model.User{
	{
		ID:        1,
		Username:  "syauqidu",
		Name:      "Syauqi Dhiya Ulhaq",
		Followers: 120,
		Following: 200,
		Posted: []model.Posting{
			{
				IDPost:  1,
				Like:    10,
				Comment: 5,
			},
			{
				IDPost:  2,
				Like:    20,
				Comment: 8,
			},
		},
	},
}

func GetUser(c *gin.Context) {
	c.JSON(202, users)
}

func TambahAngka(c *gin.Context) {
	var newJumlah model.Jumlah
	if err := c.BindJSON(&newJumlah); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	var hasil = newJumlah.Angka1 + newJumlah.Angka2
	data := gin.H{
		"Hasil": hasil,
	}
	c.JSON(203, data)
}

func HitungLuasLingkaran(c *gin.Context) {
	jariStr := c.Param("jari")
	jariFloat, err := strconv.ParseFloat(jariStr, 64)
	if err != nil {
		c.JSON(401, gin.H{"Error": err.Error()})
		return
	}
	luas := 3.14 * jariFloat * jariFloat
	c.JSON(200, gin.H{"luas": luas})
}
