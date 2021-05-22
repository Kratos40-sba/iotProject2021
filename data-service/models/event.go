package models

import "time"

type DhtEvent struct {
	Time        int64   `json:"time"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
type DhtResponse struct {
	Sensor      string    `json:"sensor"`
	Field       string    `json:"field"`
	Measurement string    `json:"measurement"`
	Time        time.Time `json:"time"`
	Value       float64   `json:"value"`
	Result      string    `json:"result"`
	Table       int       `json:"table"`
}
type RfidEvent struct {
	Time int64 `json:"time"`
	Id   int64 `json:"id"`
}

func (r RfidEvent) GetName() string {
	return "RFID-SENSOR"
}
func (d DhtEvent) GetName() string {
	return "DHT-SENSOR"
}
