package main

import (
	"github.com/Kratos40-sba/data-service/api"
	"github.com/Kratos40-sba/data-service/broker"
	"github.com/Kratos40-sba/data-service/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	_HttpServer = "HTTP_SERVER"
	DhtTopic    = "esp/dht"
	_RfidTopic  = "rfid/id"
)

func main() {
	/*
		todo check Gin/InfluxDB/Mqtt best practices
	*/
	router := gin.Default()
	influxDBConnection := database.NewConnection()
	mqttConnection := broker.NewMqttConnection()
	mqttConnection.Subscribe(influxDBConnection, DhtTopic)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", api.HealthStatusHandler(mqttConnection, influxDBConnection))
		v1.GET("/measurement?n=", api.LastNMeasurementHandler(influxDBConnection))
		v1.GET("/measurement?t=", api.LastMeasurementSinceT(influxDBConnection))
		//v1.GET("/measurement/")
	}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Route Not Defined"})
	})
	log.Fatalln(router.Run(":8080"))

}
