package models

type DhtEvent struct {
	Sensor      string  `json:"sensor"`
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
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
