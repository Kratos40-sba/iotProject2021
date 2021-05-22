package database

import (
	"context"
	"fmt"
	"github.com/Kratos40-sba/data-service/models"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
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
func (conn *Connection) ExampleInflux() []models.DhtResponse {
	var tt []models.DhtResponse
	queryAPI := conn.influxClient.QueryAPI(os.Getenv(InfluxDBOrg))
	result, err := queryAPI.Query(context.Background(),
		`from(bucket:"iot")
              |> range(start: -1h) 
              |> drop(columns : ["_start","_stop","table","_result"])
              |> rename(columns: {_value: "Value", _field: "Field",_measurement:"Measurement"})
              |> filter(fn: (r) => r["Measurement"]=="dht")
              `)
	if err == nil {
		for result.Next() {
			var t map[string]interface{}
			var res models.DhtResponse
			t = result.Record().Values()
			err := mapstructure.Decode(t, &res)
			if err != nil {
				_ = errors.New("while decoding struct accrued an error ")
			}
			fmt.Println(res)
			tt = append(tt, res)
		}
	} else {
		fmt.Printf("Query error : %s \n", result.Err().Error())
	}
	return tt
}
