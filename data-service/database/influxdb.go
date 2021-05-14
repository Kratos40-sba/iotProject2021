package database

import (
	"context"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"os"
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
