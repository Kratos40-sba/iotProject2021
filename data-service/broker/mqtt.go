package broker

import (
	"encoding/json"
	"fmt"
	"github.com/Kratos40-sba/data-service/database"
	"github.com/Kratos40-sba/data-service/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
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
func (conn *Connection) IsClientConnected() bool {
	connected := conn.mqttClient.IsConnected()
	if !connected {
		log.Println("MQTT client is not connected")
	}
	return connected
}

func onMessageReceived(influxConn *database.Connection) func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		//log.Printf("Received message : %s from topic : %s ", msg.Payload(), msg.Topic())
		// insertion here
		HEvent := models.HumEvent{Time: time.Now().Unix()}
		TEvent := models.TempEvent{Time: time.Now().Unix()}
		//b := new(bytes.Buffer)
		switch msg.Topic() {
		case "esp/dht/temperature":
			err := json.Unmarshal(msg.Payload(), &TEvent.Temperature)
			if err != nil {
				log.Println("Encoding Temperature failed : ", err)
			}
			influxConn.InsertTemp(&TEvent)
		case "esp/dht/humidity":
			err := json.Unmarshal(msg.Payload(), &HEvent.Humidity)
			if err != nil {
				log.Println("Encoding Humidity failed : ", err)
			}
			influxConn.InsertHum(&HEvent)
		default:
			log.Println("Unknown Topic : ", msg.Topic())
		}

	}
}

func (conn *Connection) Subscribe(influxConn *database.Connection, topic string) {
	token := conn.mqttClient.Subscribe(topic, QOS, onMessageReceived(influxConn))
	token.Wait()
	log.Println("Subscribed to topic : ", topic)
}
