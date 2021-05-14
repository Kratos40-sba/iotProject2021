package models

type DhtEvent interface {
}
type TempEvent struct {
	Time        int64   `json:"time"`
	Temperature float64 `json:"temperature"`
}
type HumEvent struct {
	Time     int64   `json:"time"`
	Humidity float64 `json:"humidty"`
}
