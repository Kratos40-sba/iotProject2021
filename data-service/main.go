package main

import (
	"fmt"
	"github.com/Kratos40-sba/data-service/api"
	"github.com/Kratos40-sba/data-service/database"
	"github.com/Kratos40-sba/data-service/message"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const (
	HttpServer = "HTTP_SERVER"
	DhtTopic   = "esp/dht"
	_RfidTopic = "esp/rfid"
)

func main() {
	/*
		todo check Gin/InfluxDB/Mqtt best practices
	*/
	router := gin.Default()
	influxDBConnection := database.NewConnection()
	mqttConnection := message.NewMqttConnection()
	mqttConnection.Subscribe(influxDBConnection, DhtTopic)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", api.HealthStatusHandler(mqttConnection, influxDBConnection))
		v1.GET("/measurement?n=", api.LastNMeasurementHandler(influxDBConnection))
		v1.GET("/measurement?t=", api.LastMeasurementSinceT(influxDBConnection))
		v1.GET("/measurement", api.ExampleHandler(influxDBConnection))
	}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Route Not Defined"})
	})
	log.Fatalln(router.Run(fmt.Sprintf(":%s", os.Getenv(HttpServer))))

}
