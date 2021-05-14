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
	ServerPort = ":8080"
	DhtTopic   = "esp/dht"
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
	}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Route Not Defined"})
	})
	log.Fatalln(router.Run(ServerPort))

}
