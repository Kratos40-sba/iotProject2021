package main

import (
	"github.com/kratos40-sba/mock/publisher"
	"math/rand"
	"time"
)

const (
	DhtTopic = "esp/dht"
)

func main() {
	// connect the mqtt client and publish data into mosquitto
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	mqttClient := publisher.NewMqttConnection()
	for {
		mqttClient.Publish(DhtTopic, r)
		time.Sleep(5 * time.Second)
	}

}
