package main

import (
	"github.com/Kratos40-sba/data-service/api"
	"github.com/Kratos40-sba/data-service/broker"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	ServerPort = ":8080"
	TempTopic  = "esp/dht/temperature"
	_HumTopic  = "esp/dht/humidity"
)

func main() {
	/*
	 Subscribe at multiple topics
	 Insert into database
	*/

	router := gin.Default()
	mqttConnection := broker.NewMqttConnection("Go-Client")
	mqttConnection.Subscribe(TempTopic)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", api.HealthStatusHandler(mqttConnection))
	}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Route Not Defined"})
	})
	log.Fatalln(router.Run(ServerPort))
}
