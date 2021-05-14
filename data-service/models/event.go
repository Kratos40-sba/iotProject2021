package models

type DhtEvent struct {
	Time        int64   `json:"time"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func (d DhtEvent) GetName() string {
	return "DHT-SENSOR"
}
