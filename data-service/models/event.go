package models

type Event struct {
	Time        int64   `json:"time"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
