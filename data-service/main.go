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
)

func main() {
	/*
	 Subscribe at multiple topics
	 Insert into influxDB
	*/
	router := gin.Default()
	mqttConnection := broker.NewMqttConnection("Client-1")
	mqttConnection.Subscribe("")
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", api.HealthStatusHandler(mqttConnection))
	}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Route Not Defined"})
	})
	log.Fatalln(router.Run(ServerPort))
}
