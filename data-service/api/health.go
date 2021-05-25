package api

import (
	"fmt"
	"github.com/Kratos40-sba/data-service/database"
	"github.com/Kratos40-sba/data-service/message"
	"github.com/Kratos40-sba/data-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HealthStatusHandler(mqttConn *message.Connection, influxConn *database.Connection) func(c *gin.Context) {
	return func(c *gin.Context) {
		b := models.Status{}
		if mqttConn.IsClientConnected() && influxConn.IsClientConnected() {
			b.StatusChek = fmt.Sprintf("MQTT/InfluxDB CLIENT IS CONNECTED | time : %s", time.Now().String())
			c.JSON(http.StatusOK, b)
		} else {
			b.StatusChek = fmt.Sprintf("MQTT/InfluxDB CLIENT IS DOWN | time : %s", time.Now().String())
			c.JSON(http.StatusInternalServerError, b)
		}
	}
}
