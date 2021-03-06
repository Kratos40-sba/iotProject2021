package message

import (
	"fmt"
	"github.com/Kratos40-sba/data-service/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
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

func onMessageReceived(conn *gorm.DB) func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		event := models.DhtEvent{Time: time.Now().Format("2006-01-02 15:04:05")}
		event.Sensor = event.GetName()
		switch msg.Topic() {
		// esp/rfid
		case "esp/dht":
			// todo 24||50 => temp = 24 , humidity = 50
			p := strings.Split(string(msg.Payload()), "||")
			event.Temperature, _ = strconv.ParseFloat(p[0], 32)
			event.Humidity, _ = strconv.ParseFloat(p[1], 32)
			log.Println(event)
			conn.Create(event)

		default:
			log.Println("Unknown Topic : ", msg.Topic())
		}

	}
}

func (conn *Connection) Subscribe(connDb *gorm.DB, topic string) {
	token := conn.mqttClient.Subscribe(topic, QOS, onMessageReceived(connDb))
	token.Wait()
	log.Println("Subscribed to topic : ", topic)
}
