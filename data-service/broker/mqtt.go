package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

const (
	HostFormat     = "tcp://%s:%s"
	MqttHost       = "MQTT_HOST"
	MqttPort       = "MQTT_PORT"
	MqttClientName = "MQTT_CLINT_NAME"
	MqttTopicName  = "MQTT_TOPIC_NAME"
)

var (
	host                                             = os.Getenv(MqttHost)
	port                                             = os.Getenv(MqttPort)
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

func NewMqttConnection(clientId string) (conn *Connection) {
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf(HostFormat, host, port))
	options.SetClientID(clientId)
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
