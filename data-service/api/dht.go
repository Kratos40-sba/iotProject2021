package api

import (
	"github.com/Kratos40-sba/data-service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dht []models.DhtEvent
	db.Find(&dht)
	c.JSON(http.StatusOK, gin.H{"data": dht})
}
func Last(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	k := c.Param("last")
	kk, _ := strconv.Atoi(k)
	if kk <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "check path parameter"})
		return
	}
	var dht []models.DhtEvent
	db.Limit(kk).Find(&dht)
	c.JSON(http.StatusOK, gin.H{"data": dht})
}
