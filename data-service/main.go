package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

const (
	TEMP_TOPIC = "esp/dht/temperature"
	HUM_TOPIC  = "esp/dht/humidity"
)

type Specification struct {
	MQTTHost          string `envconfig:"MQTT_HOST" required:"true"`
	MQTTPort          int    `envconfig:"MQTT_PORT" default:"1883"`
	MQTTUser          string `envconfig:"MQTT_USER" default:""`
	MQTTPwd           string `envconfig:"MQTT_PWD" default:""`
	InfluxHost        string `envconfig:"INFLUX_HOST" required:"true"`
	InfluxPort        string `envconfig:"INFLUX_PORT" required:"true"`
	InfluxDatabase    string `envconfig:"INFLUX_DATABASE" default:"telemetry"`
	InfluxUser        string `envconfig:"INFLUX_USER" default:""`
	InfluxPwd         string `envconfig:"INFLUX_PWD" default:""`
	InfluxMeasurement string `envconfig:"INFLUX_MEASUREMENT" default:"iotdata"`
}
type Dht struct {
	Temperature float64
	Humidity    float64
	Time        time.Time
}

var s Specification

func init() {
	// initialisation des variables des env
	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	// creation de client http -> InfluxDB

}
func main() {
	// creating mqtt_client
	// subscribe to multiple topics
	// save into influxdb
	fmt.Println("hello from : DATA-SERVICE")
}
