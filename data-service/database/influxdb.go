package database

import (
	"context"
	"github.com/Kratos40-sba/data-service/models"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"os"
	"time"
)

const (
	InfluxDBHost        = "INFLUX_HOST"
	InfluxDBName        = "INFLUX_NAME"
	InfluxDBMeasurement = "INFLUX_MEASUREMENT"
)

type Connection struct {
	influxClient influx.Client
}

func NewConnection() (conn *Connection) {
	influx.DefaultOptions().HTTPClient()
	client := influx.NewClient(os.Getenv(InfluxDBHost), "")
	conn = &Connection{client}
	return conn
}
func (conn *Connection) IsClientConnected() bool {
	_, err := conn.influxClient.Health(context.Background())
	if err != nil {
		log.Println("Connection to influxDB client fails : ", err)
		return false
	}
	return true
}
func (conn *Connection) Insert(event *models.DhtEvent) {
	point := influx.NewPointWithMeasurement(os.Getenv(InfluxDBMeasurement)).
		AddTag("Sensor", event.GetName()).
		AddField("temperature", event.Temperature).
		AddField("humidity", event.Humidity).
		SetTime(time.Unix(event.Time, 0))
	wAPI := conn.influxClient.WriteAPIBlocking("", os.Getenv(InfluxDBName))
	err := wAPI.WritePoint(context.Background(), point)
	if err != nil {
		log.Println("InfluxDB fails to insert : ", err)
	}
}
