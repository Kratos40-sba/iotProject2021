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
func (conn *Connection) InsertTemp(TempEvent *models.TempEvent) {
	TempPoint := influx.NewPointWithMeasurement(os.Getenv(InfluxDBMeasurement)).
		AddTag("Sensor", "DHT").
		AddField("temperature", TempEvent.Temperature).
		SetTime(time.Unix(TempEvent.Time, 0))
	wAPI := conn.influxClient.WriteAPIBlocking("", os.Getenv(InfluxDBName))
	err := wAPI.WritePoint(context.Background(), TempPoint)
	if err != nil {
		log.Println("InfluxDB fails to insert : ", err)
	}
}
func (conn *Connection) InsertHum(HumEvent *models.HumEvent) {
	HumPoint := influx.NewPointWithMeasurement(os.Getenv(InfluxDBMeasurement)).
		AddTag("Sensor", "DHT").
		AddField("temperature", HumEvent.Humidity).
		SetTime(time.Unix(HumEvent.Time, 0))
	wAPI := conn.influxClient.WriteAPIBlocking("", os.Getenv(InfluxDBName))
	err := wAPI.WritePoint(context.Background(), HumPoint)
	if err != nil {
		log.Println("InfluxDB fails to insert : ", err)
	}
}
