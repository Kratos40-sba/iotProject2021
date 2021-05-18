package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

const (
	HostFormat     = "tcp://%s:%s"
	MqttHost       = "MQTT_HOST" // 192.168.1.4
	MqttPort       = "MQTT_PORT" // 1883
	QOS            = 1
	MqttClientName = "MQTT_CLINT_NAME"
)

var (
	host                                             = os.Getenv(MqttHost)
	port                                             = os.Getenv(MqttPort)
	clientName                                       = os.Getenv(MqttClientName)
	connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		log.Println("Connection lost : ", err)
	}
	connectionHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		log.Println("Client is connected")
	}
)

type Connection struct {
	mqttClient mqtt.Client
}

func NewMqttConnection() (conn *Connection) {
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf(HostFormat, host, port))
	options.SetClientID(clientName)
	options.AutoReconnect = true
	options.OnConnectionLost = connectionLostHandler
	options.OnConnect = connectionHandler
	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("Connection Problem :", token.Error())
	}
	conn = &Connection{client}
	return conn
}
func main() {
	// connect the mqtt client and publish data into mosquitto
	fmt.Println("Mocking Sensor Data")
}
