package main

import (
	"fmt"
	"github.com/Kratos40-sba/data-service/api"
	"github.com/Kratos40-sba/data-service/config"
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

func init() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Config file has been loaded ...")
}
func main() {
	/*
		todo check Gin/InfluxDB/Mqtt best practices
	*/
	router := gin.Default()
	psgrs := database.DbConnection()
	router.Use(func(c *gin.Context) {
		c.Set("db", psgrs)
	})

	mqttConnection := message.NewMqttConnection()
	mqttConnection.Subscribe(psgrs, DhtTopic)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", api.HealthStatusHandler(mqttConnection))
		v1.GET("/dht", api.GetAll)
		v1.GET("/dht/:last", api.Last)
	}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Route Not Defined"})
	})
	log.Fatalln(router.Run(fmt.Sprintf(":%s", os.Getenv(HttpServer))))

}
