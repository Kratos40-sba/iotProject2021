package database

import (
	"context"
	"fmt"
	"github.com/Kratos40-sba/data-service/models"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"os"
	"time"
)

const (
	InfluxDBHost        = "INFLUX_HOST"        //  localhost
	InfluxDBPORT        = "INFLUX_PORT"        // 8086
	InfluxDBName        = "INFLUX_NAME"        // iot (bucket)
	InfluxDBMeasurement = "INFLUX_MEASUREMENT" // dht
	InfluxDBToken       = "INFLUX_TOKEN"
	InfluxDBOrg         = "INFLUX_ORG" // esi
)

type Connection struct {
	influxClient influx.Client
}

func NewConnection() (conn *Connection) {
	influx.DefaultOptions().HTTPClient()
	url := fmt.Sprintf("http://%s:%s", os.Getenv(InfluxDBHost), os.Getenv(InfluxDBPORT))
	client := influx.NewClient(url, os.Getenv(InfluxDBToken))
	defer client.Close()
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
	wAPI := conn.influxClient.WriteAPIBlocking(os.Getenv(InfluxDBOrg), os.Getenv(InfluxDBName))
	err := wAPI.WritePoint(context.Background(), point)
	if err != nil {
		log.Println("InfluxDB fails to insert : ", err)
	}
}

func (conn *Connection) GetLastNMeasurement(n uint) []models.DhtEvent {
	events := make([]models.DhtEvent, 0, n)
	//var e models.DhtEvent
	queryAPI := conn.influxClient.QueryAPI(os.Getenv(InfluxDBOrg))
	result, err := queryAPI.Query(context.Background(), ``)
	if err == nil {
		for result.Next() {

		}
	} else {
		fmt.Printf("Query error : %s \n", result.Err().Error())
	}
	return events
	// api/v1/measurement?n=10
}
func (conn *Connection) GetLastMeasurementSinceT(t int64) []models.DhtEvent {
	events := make([]models.DhtEvent, 0)
	queryAPI := conn.influxClient.QueryAPI(os.Getenv(InfluxDBOrg))
	result, err := queryAPI.Query(context.Background(), ``)
	if err == nil {
		for result.Next() {

		}
	} else {
		fmt.Printf("Query error : %s \n", result.Err().Error())
	}
	return events
	// api/v1/measurement?t=10
}
func (conn *Connection) ExampleInflux() []interface{} {
	var tt []interface{}
	queryAPI := conn.influxClient.QueryAPI(os.Getenv(InfluxDBOrg))
	result, err := queryAPI.Query(context.Background(), `from(bucket:"iot")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "dht" )`)
	if err == nil {
		for result.Next() {
			var t interface{}
			t = result.Record()
			tt = append(tt, t)
			/*
					var t models.DhtEvent
					  t.Temperature, _ =strconv.ParseFloat(fmt.Sprintf("%v",result.Record().ValueByKey("temperature")),32)
				         t.Humidity , _  = strconv.ParseFloat(fmt.Sprintf("%v",result.Record().ValueByKey("humidity")),32)
				         tt = append(tt, t)
			*/
		}
	} else {
		fmt.Printf("Query error : %s \n", result.Err().Error())
	}
	return tt
}
