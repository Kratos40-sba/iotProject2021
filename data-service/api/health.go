package api

import (
	"fmt"
	"github.com/Kratos40-sba/data-service/broker"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HealthStatusHandler(mqttConn *broker.Connection) func(c *gin.Context) {
	return func(c *gin.Context) {
		if mqttConn.IsClientConnected() {
			c.JSON(http.StatusOK, fmt.Sprintf("MQTT CLIENT IS CONNECTED |> time : %s", time.Now().String()))
		} else {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("MQTT CLIENT IS DOWN |> time : %s", time.Now().String()))
		}
	}
}
