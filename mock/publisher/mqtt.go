package publisher

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
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
	temp = 20.0
	hum  = 50.0
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
func GenerateRandomValues(r *rand.Rand) string {
	t := temp + r.Float64()*10
	h := hum + r.Float64()*10
	return fmt.Sprintf("%.2f||%.2f", t, h)
}
func (conn *Connection) Publish(topicName string, r *rand.Rand) {
	token := conn.mqttClient.Publish(topicName, QOS, false, GenerateRandomValues(r))
	token.Wait()
	log.Printf("Publishing on Topic : %s  payload : %s ", topicName, GenerateRandomValues(r))
}
